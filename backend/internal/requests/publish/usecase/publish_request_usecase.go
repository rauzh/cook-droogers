package usecase

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/broker"
	criteria "cookdroogers/internal/requests/criteria_controller"
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
	pbBroker        broker.IBroker
	criterias       criteria.ICriteriaCollection

	repo publishReqRepo.PublishRequestRepo
}

func NewPublishRequestUseCase(
	statService statService.IStatisticsService,
	publicationRepo repo.PublicationRepo,
	releaseRepo repo.ReleaseRepo,
	artistRepo repo.ArtistRepo,
	transactor transactor.Transactor,
	pbBroker broker.IBroker,
	criterias criteria.ICriteriaCollection,
	repo publishReqRepo.PublishRequestRepo,
) (base.IRequestUseCase, error) {

	err := pbBroker.SignConsumerToTopic(PublishRequestProceedToManager)
	if err != nil {
		return nil, err
	}

	publishUseCase := &PublishRequestUseCase{
		statService:     statService,
		publicationRepo: publicationRepo,
		releaseRepo:     releaseRepo,
		artistRepo:      artistRepo,
		repo:            repo,
		transactor:      transactor,
		criterias:       criterias,
		pbBroker:        pbBroker,
	}

	err = publishUseCase.runProceedToManagerConsumer()
	if err != nil {
		return nil, err
	}

	return publishUseCase, nil
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

	if err := publishUseCase.sendProceedToManagerMSG(pubReq); err != nil {
		return err
	}

	return nil
}

func (publishUseCase *PublishRequestUseCase) proceedToManager(pubReq *publish.PublishRequest) error {

	ctx := context.Background()

	pubReq.Status = base.ProcessingRequest

	err := publishUseCase.repo.Update(ctx, pubReq)
	if err != nil {
		return fmt.Errorf("cant proceed publish request to manager with err %w", err)
	}

	publishUseCase.computeDegree(pubReq)

	artist, err := publishUseCase.artistRepo.Get(ctx, pubReq.ApplierID)
	if err != nil {
		return fmt.Errorf("cant proceed publish request to manager with err %w", err)
	}

	pubReq.ManagerID = artist.ManagerID
	pubReq.Status = base.OnApprovalRequest

	return publishUseCase.repo.Update(ctx, pubReq)
}

func (publishUseCase *PublishRequestUseCase) computeDegree(pubReq *publish.PublishRequest) {

	summaryDiff := publishUseCase.criterias.Apply(pubReq)

	pubReq.Grade = summaryDiff.ResultDiff
	for criteriaName, criteriaDiff := range summaryDiff.ResultExplanation {
		pubReq.Description += criteria.DiffToString(criteriaName, criteriaDiff.Explanation, criteriaDiff.Diff)
	}
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
