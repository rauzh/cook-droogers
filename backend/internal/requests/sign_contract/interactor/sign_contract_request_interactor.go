package interactor

import (
	"context"
	repo "cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/sign_contract"
	sctErrors "cookdroogers/internal/requests/sign_contract/errors"
	signContractRepo "cookdroogers/internal/requests/sign_contract/repo"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	"fmt"
	"time"
)

type SignContractRequestInteractor struct {
	req *sign_contract.SignContractRequest

	mngRepo    repo.ManagerRepo
	userRepo   repo.UserRepo
	artistRepo repo.ArtistRepo
	transactor transactor.Transactor

	repo signContractRepo.SignContractRequestRepo
}

func NewSignContractRequestInteractor(
	mngRepo repo.ManagerRepo,
	usrRepo repo.UserRepo,
	artRepo repo.ArtistRepo,
	transactor transactor.Transactor,
	repo signContractRepo.SignContractRequestRepo,
) base.IRequest {
	return &SignContractRequestInteractor{
		mngRepo:    mngRepo,
		userRepo:   usrRepo,
		artistRepo: artRepo,
		repo:       repo,
		transactor: transactor,
	}
}

func (sctInteractor *SignContractRequestInteractor) validate() error {

	if sctInteractor.req == nil {
		return sctErrors.ErrNoReq
	}

	if sctInteractor.req.Nickname == "" || len(sctInteractor.req.Nickname) > sign_contract.MaxNicknameLen {
		return sctErrors.ErrNickname
	}

	if sctInteractor.req.ApplierID == sign_contract.EmptyID {
		return sctErrors.ErrNoApplierID
	}

	if sctInteractor.req.Type != base.SignRequest {
		return sctErrors.ErrInvalidType
	}

	return nil
}

func (sctInteractor *SignContractRequestInteractor) SetReq(signReq *sign_contract.SignContractRequest) {
	sctInteractor.req = signReq
}

func (sctInteractor *SignContractRequestInteractor) Apply() error {

	if err := sctInteractor.validate(); err != nil {
		return err
	}

	base.InitDateStatus(&sctInteractor.req.Request)

	if err := sctInteractor.repo.Create(context.Background(), sctInteractor.req); err != nil {
		return fmt.Errorf("can't apply sign contract request with err %w", err)
	}

	//go sctInteractor.proceedToManager(request)  ПРИКРУТИТЬ КАФКУ

	return nil
}

//func (sctInteractor *SignContractRequestInteractor) proceedToManager(signReq *sign_contract.SignContractRequest) {
//	signReq.Status = base.OnApprovalRequest
//
//	managerID, err := sctInteractor.mngRepo.GetRandManagerID()
//	if err == nil {
//		signReq.ManagerID = managerID
//	} else {
//		signReq.Status = base.ClosedRequest
//		signReq.Description= "Can't find manager"
//	}
//
//	sctInteractor.repo.Update(context.Background(), signReq)
//}

func (sctInteractor *SignContractRequestInteractor) Accept() error {

	if err := sctInteractor.validate(); err != nil {
		return err
	}

	artist := models.Artist{
		UserID:       sctInteractor.req.ApplierID,
		Nickname:     sctInteractor.req.Nickname,
		ContractTerm: time.Now().AddDate(sign_contract.YearsContract, sign_contract.MonthsContract, sign_contract.DaysContract),
		Activity:     true,
		ManagerID:    sctInteractor.req.ManagerID,
	}

	ctx := context.Background()
	return sctInteractor.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := sctInteractor.artistRepo.Create(ctx, &artist); err != nil {
			return fmt.Errorf("can't create artist %s with err %w", artist.Nickname, err)
		}

		if err := sctInteractor.userRepo.UpdateType(ctx, artist.UserID, models.ArtistUser); err != nil {
			return fmt.Errorf("can't update user with err %w", err)
		}

		sctInteractor.req.Status = base.ClosedRequest
		if err := sctInteractor.repo.Update(ctx, sctInteractor.req); err != nil {
			return fmt.Errorf("can't update reqiest with err %w", err)
		}

		return nil
	})
}

func (sctInteractor *SignContractRequestInteractor) Decline() error {

	if err := sctInteractor.validate(); err != nil {
		return err
	}

	sctInteractor.req.Status = base.ClosedRequest
	sctInteractor.req.Description = "The request is declined."

	return sctInteractor.repo.Update(context.Background(), sctInteractor.req)
}

func (sctInteractor *SignContractRequestInteractor) GetType() base.RequestType {

	if sctInteractor.req == nil {
		return ""
	}

	return sctInteractor.req.Type
}
