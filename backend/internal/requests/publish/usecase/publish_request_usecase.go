package usecase

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	pubReqErrors "cookdroogers/internal/requests/publish/errors"
	publishReqRepo "cookdroogers/internal/requests/publish/repo"
	statService "cookdroogers/internal/statistics/service"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"fmt"
)

type PublishRequestUseCase struct {
	req *publish.PublishRequest

	statService     statService.IStatisticsService
	publicationRepo repo.PublicationRepo
	releaseRepo     repo.ReleaseRepo
	artistRepo      repo.ArtistRepo
	transactor      transactor.Transactor

	repo publishReqRepo.PublishRequestRepo
}

func NewPublishRequestUseCase(
	statService statService.IStatisticsService,
	publicationRepo repo.PublicationRepo,
	releaseRepo repo.ReleaseRepo,
	artistRepo repo.ArtistRepo,
	transactor transactor.Transactor,
	repo publishReqRepo.PublishRequestRepo,
) base.IRequestUseCase {
	return &PublishRequestUseCase{
		statService:     statService,
		publicationRepo: publicationRepo,
		releaseRepo:     releaseRepo,
		artistRepo:      artistRepo,
		repo:            repo,
		transactor:      transactor,
	}
}

func (publishUseCase *PublishRequestUseCase) validate() error {

	if publishUseCase.req == nil {
		return pubReqErrors.ErrNoReq
	}

	if publishUseCase.req.ExpectedDate.IsZero() || cdtime.CheckDateWeekLater(publishUseCase.req.ExpectedDate) {
		return pubReqErrors.ErrInvalidDate
	}

	if publishUseCase.req.ReleaseID == publish.EmptyID {
		return pubReqErrors.ErrNoReleaseID
	}

	if publishUseCase.req.ApplierID == publish.EmptyID {
		return pubReqErrors.ErrNoApplierID
	}

	if publishUseCase.req.Type != base.PublishRequest {
		return pubReqErrors.ErrInvalidType
	}

	return nil
}

func (publishUseCase *PublishRequestUseCase) Apply() error {

	if err := publishUseCase.validate(); err != nil {
		return err
	}

	base.InitDateStatus(&publishUseCase.req.Request)

	if err := publishUseCase.repo.Create(context.Background(), publishUseCase.req); err != nil {
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

func (publishUseCase *PublishRequestUseCase) Accept() error {

	if err := publishUseCase.validate(); err != nil {
		return err
	}

	publication := models.Publication{
		ReleaseID: publishUseCase.req.ReleaseID,
		Date:      publishUseCase.req.ExpectedDate,
		ManagerID: publishUseCase.req.ManagerID,
	}

	ctx := context.Background()
	return publishUseCase.transactor.WithinTransaction(ctx, func(ctx context.Context) error {

		if err := publishUseCase.publicationRepo.Create(ctx, &publication); err != nil {
			return fmt.Errorf("can't create publication with err %w", err)
		}

		if err := publishUseCase.releaseRepo.UpdateStatus(ctx, publication.ReleaseID, models.PublishedRelease); err != nil {
			return fmt.Errorf("can't update publication with err %w", err)
		}

		publishUseCase.req.Status = base.ClosedRequest
		if err := publishUseCase.repo.Update(ctx, publishUseCase.req); err != nil {
			return fmt.Errorf("can't update request.go with err %w", err)
		}

		return nil
	})
}

func (publishUseCase *PublishRequestUseCase) Decline() error {

	if err := publishUseCase.validate(); err != nil {
		return err
	}

	publishUseCase.req.Status = base.ClosedRequest
	publishUseCase.req.Description = base.DescrDeclinedRequest

	return publishUseCase.repo.Update(context.Background(), publishUseCase.req)
}

func (publishUseCase *PublishRequestUseCase) GetType() base.RequestType {

	if publishUseCase.req == nil {
		return ""
	}

	return publishUseCase.req.Type
}

func (publishUseCase *PublishRequestUseCase) SetReq(pubReq *publish.PublishRequest) {
	publishUseCase.req = pubReq
}

func (publishUseCase *PublishRequestUseCase) Get(id uint64) (*publish.PublishRequest, error) {
	req, err := publishUseCase.repo.Get(context.Background(), id)

	if err != nil {
		return nil, fmt.Errorf("can't get publish request with err %w", err)
	}
	return req, nil
}
