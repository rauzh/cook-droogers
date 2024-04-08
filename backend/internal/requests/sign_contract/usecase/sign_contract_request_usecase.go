package usecase

import (
	"context"
	repo "cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/broker"
	"cookdroogers/internal/requests/sign_contract"
	"cookdroogers/internal/requests/sign_contract/errors"
	signContractRepo "cookdroogers/internal/requests/sign_contract/repo"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"fmt"
)

type SignContractRequestUseCase struct {
	mngRepo    repo.ManagerRepo
	userRepo   repo.UserRepo
	artistRepo repo.ArtistRepo
	transactor transactor.Transactor
	scBroker   broker.IBroker

	repo signContractRepo.SignContractRequestRepo
}

func NewSignContractRequestUseCase(
	mngRepo repo.ManagerRepo,
	usrRepo repo.UserRepo,
	artRepo repo.ArtistRepo,
	transactor transactor.Transactor,
	scBroker broker.IBroker,
	repo signContractRepo.SignContractRequestRepo,
) (base.IRequestUseCase, error) {

	err := scBroker.SignConsumerToTopic(SignRequestProceedToManager)
	if err != nil {
		return nil, err
	}

	sctUseCase := &SignContractRequestUseCase{
		mngRepo:    mngRepo,
		userRepo:   usrRepo,
		artistRepo: artRepo,
		repo:       repo,
		transactor: transactor,
		scBroker:   scBroker,
	}

	err = sctUseCase.runProceedToManagerConsumer()
	if err != nil {
		return nil, err
	}

	return sctUseCase, nil
}

func (sctUseCase *SignContractRequestUseCase) Apply(request base.IRequest) error {

	if err := request.Validate(sign_contract.SignRequest); err != nil {
		return err
	}
	signReq := request.(*sign_contract.SignContractRequest)

	base.InitDateStatus(&signReq.Request)

	if err := sctUseCase.repo.Create(context.Background(), signReq); err != nil {
		return fmt.Errorf("can't apply sign contract request with err %w", err)
	}

	if err := sctUseCase.sendProceedToManagerMSG(signReq); err != nil {
		return err
	}

	return nil
}

func (sctUseCase *SignContractRequestUseCase) proceedToManager(signReq *sign_contract.SignContractRequest) error {
	signReq.Status = base.OnApprovalRequest

	managerID, err := sctUseCase.mngRepo.GetRandManagerID(context.Background())
	if err != nil {
		return errors.ErrCantFindManager
	}

	signReq.ManagerID = managerID

	return sctUseCase.repo.Update(context.Background(), signReq)
}

func (sctUseCase *SignContractRequestUseCase) Accept(request base.IRequest) error {

	if err := request.Validate(sign_contract.SignRequest); err != nil {
		return err
	}
	signReq := request.(*sign_contract.SignContractRequest)

	artist := models.Artist{
		UserID:       signReq.ApplierID,
		Nickname:     signReq.Nickname,
		ContractTerm: cdtime.GetEndOfContract(),
		Activity:     true,
		ManagerID:    signReq.ManagerID,
	}

	ctx := context.Background()
	return sctUseCase.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := sctUseCase.artistRepo.Create(ctx, &artist); err != nil {
			return fmt.Errorf("can't create artist %s with err %w", artist.Nickname, err)
		}

		if err := sctUseCase.userRepo.UpdateType(ctx, artist.UserID, models.ArtistUser); err != nil {
			return fmt.Errorf("can't update user with err %w", err)
		}

		signReq.Status = base.ClosedRequest
		if err := sctUseCase.repo.Update(ctx, signReq); err != nil {
			return fmt.Errorf("can't update reqiest with err %w", err)
		}

		return nil
	})
}

func (sctUseCase *SignContractRequestUseCase) Decline(request base.IRequest) error {

	if err := request.Validate(sign_contract.SignRequest); err != nil {
		return err
	}
	signReq := request.(*sign_contract.SignContractRequest)

	signReq.Status = base.ClosedRequest
	signReq.Description = base.DescrDeclinedRequest

	return sctUseCase.repo.Update(context.Background(), signReq)
}

func (sctUseCase *SignContractRequestUseCase) Get(id uint64) (*sign_contract.SignContractRequest, error) {

	req, err := sctUseCase.repo.Get(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("can't get sign contract request with err %w", err)
	}

	return req, nil
}
