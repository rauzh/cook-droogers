package usecase

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	publishReqRepo "cookdroogers/internal/requests/publish/repo"
	statService "cookdroogers/internal/statistics/service"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	"fmt"
)

type PublishRequestUseCase struct {
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

func (publishUseCase *PublishRequestUseCase) Apply(request base.IRequest) error {

	if err := request.Validate(publish.PubReq); err != nil {
		return err
	}
	pubReq := request.(*publish.PublishRequest)

	base.InitDateStatus(&pubReq.Request)

	if err := publishUseCase.repo.Create(context.Background(), pubReq); err != nil {
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

func (publishUseCase *PublishRequestUseCase) Accept(request base.IRequest) error {

	if err := request.Validate(publish.PubReq); err != nil {
		return err
	}
	pubReq := request.(*publish.PublishRequest)

	publication := models.Publication{
		ReleaseID: pubReq.ReleaseID,
		Date:      pubReq.ExpectedDate,
		ManagerID: pubReq.ManagerID,
	}

	ctx := context.Background()
	return publishUseCase.transactor.WithinTransaction(ctx, func(ctx context.Context) error {

		if err := publishUseCase.publicationRepo.Create(ctx, &publication); err != nil {
			return fmt.Errorf("can't create publication with err %w", err)
		}

		if err := publishUseCase.releaseRepo.UpdateStatus(ctx, publication.ReleaseID, models.PublishedRelease); err != nil {
			return fmt.Errorf("can't update publication with err %w", err)
		}

		pubReq.Status = base.ClosedRequest
		if err := publishUseCase.repo.Update(ctx, pubReq); err != nil {
			return fmt.Errorf("can't update request.go with err %w", err)
		}

		return nil
	})
}

func (publishUseCase *PublishRequestUseCase) Decline(request base.IRequest) error {

	if err := request.Validate(publish.PubReq); err != nil {
		return err
	}
	pubReq := request.(*publish.PublishRequest)

	pubReq.Status = base.ClosedRequest
	pubReq.Description = base.DescrDeclinedRequest

	return publishUseCase.repo.Update(context.Background(), pubReq)
}

func (publishUseCase *PublishRequestUseCase) Get(id uint64) (*publish.PublishRequest, error) {
	req, err := publishUseCase.repo.Get(context.Background(), id)

	if err != nil {
		return nil, fmt.Errorf("can't get publish request with err %w", err)
	}
	return req, nil
}
