package usecase

import (
	"context"
	repo "cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/sign_contract"
	sctErrors "cookdroogers/internal/requests/sign_contract/errors"
	signContractRepo "cookdroogers/internal/requests/sign_contract/repo"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"fmt"
)

type SignContractRequestUseCase struct {
	req *sign_contract.SignContractRequest

	mngRepo    repo.ManagerRepo
	userRepo   repo.UserRepo
	artistRepo repo.ArtistRepo
	transactor transactor.Transactor

	repo signContractRepo.SignContractRequestRepo
}

func NewSignContractRequestUseCase(
	mngRepo repo.ManagerRepo,
	usrRepo repo.UserRepo,
	artRepo repo.ArtistRepo,
	transactor transactor.Transactor,
	repo signContractRepo.SignContractRequestRepo,
) base.IRequestUseCase {
	return &SignContractRequestUseCase{
		mngRepo:    mngRepo,
		userRepo:   usrRepo,
		artistRepo: artRepo,
		repo:       repo,
		transactor: transactor,
	}
}

func (sctUseCase *SignContractRequestUseCase) validate() error {

	if sctUseCase.req == nil {
		return sctErrors.ErrNoReq
	}

	if sctUseCase.req.Nickname == "" || len(sctUseCase.req.Nickname) > sign_contract.MaxNicknameLen {
		return sctErrors.ErrNickname
	}

	if sctUseCase.req.ApplierID == sign_contract.EmptyID {
		return sctErrors.ErrNoApplierID
	}

	if sctUseCase.req.Type != base.SignRequest {
		return sctErrors.ErrInvalidType
	}

	return nil
}

func (sctUseCase *SignContractRequestUseCase) Apply() error {

	if err := sctUseCase.validate(); err != nil {
		return err
	}

	base.InitDateStatus(&sctUseCase.req.Request)

	if err := sctUseCase.repo.Create(context.Background(), sctUseCase.req); err != nil {
		return fmt.Errorf("can't apply sign contract request with err %w", err)
	}

	//go sctUseCase.proceedToManager(request)  ПРИКРУТИТЬ КАФКУ

	return nil
}

//func (sctUseCase *SignContractRequestUseCase) proceedToManager(signReq *sign_contract.SignContractRequest) {
//	signReq.Status = base.OnApprovalRequest
//
//	managerID, err := sctUseCase.mngRepo.GetRandManagerID()
//	if err == nil {
//		signReq.ManagerID = managerID
//	} else {
//		signReq.Status = base.ClosedRequest
//		signReq.Description= "Can't find manager"
//	}
//
//	sctUseCase.repo.Update(context.Background(), signReq)
//}

func (sctUseCase *SignContractRequestUseCase) Accept() error {

	if err := sctUseCase.validate(); err != nil {
		return err
	}

	artist := models.Artist{
		UserID:       sctUseCase.req.ApplierID,
		Nickname:     sctUseCase.req.Nickname,
		ContractTerm: cdtime.GetEndOfContract(),
		Activity:     true,
		ManagerID:    sctUseCase.req.ManagerID,
	}

	ctx := context.Background()
	return sctUseCase.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := sctUseCase.artistRepo.Create(ctx, &artist); err != nil {
			return fmt.Errorf("can't create artist %s with err %w", artist.Nickname, err)
		}

		if err := sctUseCase.userRepo.UpdateType(ctx, artist.UserID, models.ArtistUser); err != nil {
			return fmt.Errorf("can't update user with err %w", err)
		}

		sctUseCase.req.Status = base.ClosedRequest
		if err := sctUseCase.repo.Update(ctx, sctUseCase.req); err != nil {
			return fmt.Errorf("can't update reqiest with err %w", err)
		}

		return nil
	})
}

func (sctUseCase *SignContractRequestUseCase) Decline() error {

	if err := sctUseCase.validate(); err != nil {
		return err
	}

	sctUseCase.req.Status = base.ClosedRequest
	sctUseCase.req.Description = base.DescrDeclinedRequest

	return sctUseCase.repo.Update(context.Background(), sctUseCase.req)
}

func (sctUseCase *SignContractRequestUseCase) GetType() base.RequestType {

	if sctUseCase.req == nil {
		return ""
	}

	return sctUseCase.req.Type
}

func (sctUseCase *SignContractRequestUseCase) SetReq(signReq *sign_contract.SignContractRequest) {
	sctUseCase.req = signReq
}

func (sctUseCase *SignContractRequestUseCase) Get(id uint64) (*sign_contract.SignContractRequest, error) {
	req, err := sctUseCase.repo.Get(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("can't get sign contract request with err %w", err)
	}
	return req, nil
}
