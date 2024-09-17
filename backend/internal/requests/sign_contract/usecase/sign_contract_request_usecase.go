package usecase

import (
	"context"
	repo "cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/broker"
	"cookdroogers/internal/requests/broker/broker_dto"
	signContractBroker "cookdroogers/internal/requests/broker/sign_contract"
	"cookdroogers/internal/requests/sign_contract"
	signContractRepo "cookdroogers/internal/requests/sign_contract/repo"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"fmt"
	"log/slog"
)

type SignContractRequestUseCase struct {
	userRepo   repo.UserRepo
	artistRepo repo.ArtistRepo
	transactor transactor.Transactor
	scBroker   broker.IBroker

	repo signContractRepo.SignContractRequestRepo

	logger *slog.Logger
}

func NewSignContractRequestUseCase(
	usrRepo repo.UserRepo,
	artRepo repo.ArtistRepo,
	transactor transactor.Transactor,
	scBroker broker.IBroker,
	repo signContractRepo.SignContractRequestRepo,
	logger *slog.Logger,
) (base.IRequestUseCase, error) {

	sctUseCase := &SignContractRequestUseCase{
		userRepo:   usrRepo,
		artistRepo: artRepo,
		repo:       repo,
		transactor: transactor,
		scBroker:   scBroker,
		logger:     logger,
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
		if err := sctUseCase.userRepo.UpdateType(ctx, artist.UserID, models.ArtistUser); err != nil {
			return fmt.Errorf("can't update user with err %w", err)
		}

		if err := sctUseCase.artistRepo.Create(ctx, &artist); err != nil {
			return fmt.Errorf("can't create artist %s with err %w", artist.Nickname, err)
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

func (sctUseCase *SignContractRequestUseCase) sendProceedToManagerMSG(signReq *sign_contract.SignContractRequest) error {

	msg, err := broker_dto.NewSignRequestProducerMsg(signContractBroker.SignRequestProceedToManager, signReq)
	if err != nil {
		return fmt.Errorf("can't apply sign contract request: can't proceed to manager with err %w", err)
	}

	_, _, err = sctUseCase.scBroker.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("can't apply sign contract request: can't proceed to manager with err %w", err)
	}

	return nil
}
