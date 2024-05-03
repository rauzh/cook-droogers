package usecase

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/broker"
	"cookdroogers/internal/requests/broker/broker_dto"
	publish_req_broker "cookdroogers/internal/requests/broker/publish"
	"cookdroogers/internal/requests/publish"
	"cookdroogers/internal/requests/publish/errors"
	publishReqRepo "cookdroogers/internal/requests/publish/repo"
	statService "cookdroogers/internal/statistics/service"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	"fmt"
	"log/slog"
)

type PublishRequestUseCase struct {
	statService     statService.IStatisticsService
	publicationRepo repo.PublicationRepo
	releaseRepo     repo.ReleaseRepo
	artistRepo      repo.ArtistRepo
	transactor      transactor.Transactor
	broker          broker.IBroker

	repo publishReqRepo.PublishRequestRepo

	logger *slog.Logger
}

func NewPublishRequestUseCase(
	statService statService.IStatisticsService,
	publicationRepo repo.PublicationRepo,
	releaseRepo repo.ReleaseRepo,
	artistRepo repo.ArtistRepo,
	transactor transactor.Transactor,
	pbBroker broker.IBroker,
	repo publishReqRepo.PublishRequestRepo,
	logger *slog.Logger,
) (base.IRequestUseCase, error) {

	publishUseCase := &PublishRequestUseCase{
		statService:     statService,
		publicationRepo: publicationRepo,
		releaseRepo:     releaseRepo,
		artistRepo:      artistRepo,
		repo:            repo,
		transactor:      transactor,
		broker:          pbBroker,
		logger:          logger,
	}

	return publishUseCase, nil
}

func (publishUseCase *PublishRequestUseCase) Apply(request base.IRequest) error {

	if err := request.Validate(publish.PubReq); err != nil {
		return err
	}
	pubReq := request.(*publish.PublishRequest)

	base.InitDateStatus(&pubReq.Request)

	if err := publishUseCase.checkRelease(pubReq); err != nil {
		return fmt.Errorf("can't apply publish request with err %w", err)
	}

	if err := publishUseCase.repo.Create(context.Background(), pubReq); err != nil {
		return fmt.Errorf("can't apply publish request with err %w", err)
	}

	if err := publishUseCase.sendProceedToManagerMSG(pubReq); err != nil {
		return err
	}

	return nil
}

func (publishUseCase *PublishRequestUseCase) checkRelease(pubReq *publish.PublishRequest) error {

	ctx := context.Background()

	release, err := publishUseCase.releaseRepo.Get(ctx, pubReq.ReleaseID)
	if err != nil {
		return err
	}

	if release.Status != models.UnpublishedRelease {
		return errors.ErrReleaseAlreadyPublished
	}

	artist, err := publishUseCase.artistRepo.GetByUserID(ctx, pubReq.ApplierID)
	if err != nil {
		return err
	}

	if release.ArtistID != artist.ArtistID {
		return errors.ErrNotOwner
	}

	if artist.ContractTerm.Before(pubReq.ExpectedDate) {
		return errors.ErrEndContract
	}

	pubReq.ManagerID = artist.ManagerID

	return nil
}

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

func (publishUseCase *PublishRequestUseCase) sendProceedToManagerMSG(pubReq *publish.PublishRequest) error {

	msg, err := broker_dto.NewPublishRequestProducerMsg(publish_req_broker.PublishRequestProceedToManager, pubReq)
	if err != nil {
		return fmt.Errorf("can't apply publish request: can't proceed to manager with err %w", err)
	}

	_, _, err = publishUseCase.broker.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("can't apply publish request: can't proceed to manager with err %w", err)
	}

	return nil
}
