package usecase

import (
	"context"
	"cookdroogers/internal/repo/mocks"
	"cookdroogers/internal/requests/base"
	base_errors "cookdroogers/internal/requests/base/errors"
	broker_mocks "cookdroogers/internal/requests/broker/mocks"
	"cookdroogers/internal/requests/publish"
	"cookdroogers/internal/requests/publish/data_builder"
	pubReqErrors "cookdroogers/internal/requests/publish/errors"
	publishReqRepoMocks "cookdroogers/internal/requests/publish/repo/mocks"
	transacMock "cookdroogers/internal/transactor/mocks"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	_trackRepo      *mocks.TrackRepo
	publicationRepo *mocks.PublicationRepo
	releaseRepo     *mocks.ReleaseRepo
	artistRepo      *mocks.ArtistRepo
	transactor      *transacMock.Transactor
	pbBroker        *broker_mocks.IBroker
	publishRepo     *publishReqRepoMocks.PublishRequestRepo
	logger          *slog.Logger
}

func _newMockPublishReqDepFields(t provider.T) *_depFields {
	transactionMock := transacMock.NewTransactor(t)
	pbcMockRepo := mocks.NewPublicationRepo(t)
	rlsMockRepo := mocks.NewReleaseRepo(t)
	artistMockRepo := mocks.NewArtistRepo(t)
	trkMockRepo := mocks.NewTrackRepo(t)
	publishMockRepo := publishReqRepoMocks.NewPublishRequestRepo(t)
	mockBroker := broker_mocks.NewIBroker(t)

	return &_depFields{
		_trackRepo:      trkMockRepo,
		publicationRepo: pbcMockRepo,
		releaseRepo:     rlsMockRepo,
		artistRepo:      artistMockRepo,
		transactor:      transactionMock,
		pbBroker:        mockBroker,
		publishRepo:     publishMockRepo,
		logger:          slog.Default(),
	}
}

type PublishRequestUseCaseSuite struct {
	suite.Suite
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_DeclineOK(t provider.T) {
	t.Title("Decline: OK")
	t.Tags("PublishRequestUseCase")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockPublishReqDepFields(t)

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.OnApprovalRequest).
			WithReleaseID(777).
			WithGrade(-3).
			WithExpexctedDate(cdtime.GetToday().AddDate(1, 0, 0)).
			Build()

		df.publishRepo.EXPECT().Update(mock.AnythingOfType("context.backgroundCtx"), &publish.PublishRequest{
			Request: base.Request{
				RequestID: 1,
				Type:      publish.PubReq,
				Status:    base.ClosedRequest,
				Date:      cdtime.GetToday(),
				ApplierID: 12,
				ManagerID: 9,
			},
			ReleaseID:    777,
			Grade:        -3,
			ExpectedDate: cdtime.GetToday().AddDate(1, 0, 0),
			Description:  base.DescrDeclinedRequest,
		}).Return(nil).Once()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		err := publishUseCase.Decline(context.Background(), req)

		sCtx.Assert().NoError(err)
	})
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_DeclineInvalidDate(t provider.T) {
	t.Title("Decline: InvalidDate")
	t.Tags("PublishRequestUseCase")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockPublishReqDepFields(t)

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.OnApprovalRequest).
			WithReleaseID(777).
			WithGrade(-3).
			WithExpexctedDate(cdtime.GetToday().AddDate(0, 0, 3)).
			Build()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		err := publishUseCase.Decline(context.Background(), req)

		sCtx.Assert().ErrorIs(err, pubReqErrors.ErrInvalidDate)
	})
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_AcceptOK(t provider.T) {
	t.Title("Accept: OK")
	t.Tags("PublishRequestUseCase")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockPublishReqDepFields(t)

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.OnApprovalRequest).
			WithReleaseID(777).
			WithGrade(-3).
			WithExpexctedDate(cdtime.GetToday().AddDate(1, 0, 0)).
			Build()

		//publication := models.Publication{
		//	ReleaseID: req.ReleaseID,
		//	Date:      req.ExpectedDate,
		//	ManagerID: req.ManagerID,
		//}

		df.transactor.EXPECT().WithinTransaction(mock.AnythingOfType("context.backgroundCtx"), mock.Anything).Return(nil).Once()

		//RunAndReturn(func(_ context.Context, fn func(ctx context.Context) error) error {
		//	return fn(context.Background())
		//}).Once()

		//df.publicationRepo.EXPECT().Create(mock.Anything, &publication).Return(nil).Once()

		//df.releaseRepo.EXPECT().UpdateStatus(mock.Anything, publication.ReleaseID, models.PublishedRelease).Return(nil).Once()

		//df.publishRepo.EXPECT().Update(mock.Anything, req).Return(nil).Once()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		err := publishUseCase.Accept(context.Background(), req)

		sCtx.Assert().NoError(err)
	})
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_AcceptInvalidType(t provider.T) {
	t.Title("Accept: InvalidType")
	t.Tags("PublishRequestUseCase")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockPublishReqDepFields(t)

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.OnApprovalRequest).
			WithReleaseID(777).
			WithGrade(-3).
			WithExpexctedDate(cdtime.GetToday().AddDate(1, 0, 3)).
			WithType("").
			Build()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		err := publishUseCase.Accept(context.Background(), req)

		sCtx.Assert().ErrorIs(err, base_errors.ErrInvalidType)
	})
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_AcceptAlreadyClosed(t provider.T) {
	t.Title("Accept: AlreadyClosed")
	t.Tags("PublishRequestUseCase")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockPublishReqDepFields(t)

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.ClosedRequest).
			WithReleaseID(777).
			WithGrade(-3).
			WithExpexctedDate(cdtime.GetToday().AddDate(1, 0, 3)).
			Build()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		err := publishUseCase.Accept(context.Background(), req)

		sCtx.Assert().ErrorIs(err, base_errors.ErrAlreadyClosed)
	})
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_GetOK(t provider.T) {
	t.Title("Get: OK")
	t.Tags("PublishRequestUseCase")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		df := _newMockPublishReqDepFields(t)
		df.publishRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(1)).
			Return(data_builder.NewPublishRequestBuilder().WithRequestID(1).Build(), nil).Once()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		req, err := publishUseCase.(*PublishRequestUseCase).Get(context.Background(), uint64(1))

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(req)
		sCtx.Assert().Equal(data_builder.NewPublishRequestBuilder().WithRequestID(1).Build(), req)
	})
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_GetDbErr(t provider.T) {
	t.Title("Get: DBerr")
	t.Tags("PublishRequestUseCase")
	t.Parallel()
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		df := _newMockPublishReqDepFields(t)
		df.publishRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(1)).
			Return(nil, sql.ErrConnDone).Once()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		req, err := publishUseCase.(*PublishRequestUseCase).Get(context.Background(), uint64(1))

		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
		sCtx.Assert().Nil(req)
	})
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_ApplyOK(t provider.T) {
	t.Title("Apply: OK")
	t.Tags("PublishRequestUseCase")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockPublishReqDepFields(t)

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(8).
			WithStatus(base.OnApprovalRequest).
			WithReleaseID(777).
			WithGrade(-3).
			WithExpexctedDate(cdtime.GetToday().AddDate(1, 0, 0)).
			Build()

		df.releaseRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(777)).
			Return(&models.Release{
				ReleaseID:    777,
				Title:        "luv2",
				Status:       models.UnpublishedRelease,
				DateCreation: cdtime.GetToday().AddDate(-1, 0, 0),
				Tracks:       []uint64{7, 8, 9},
				ArtistID:     228,
			}, nil).Once()

		df.artistRepo.EXPECT().GetByUserID(mock.AnythingOfType("context.backgroundCtx"), uint64(12)).
			Return(&models.Artist{
				ArtistID:     228,
				UserID:       12,
				Nickname:     "uzi",
				ContractTerm: cdtime.GetToday().AddDate(10, 0, 0),
				Activity:     true,
				ManagerID:    8,
			}, nil).Once()

		df.publishRepo.EXPECT().Create(mock.AnythingOfType("context.backgroundCtx"), req).Return(nil).Once()
		df.pbBroker.EXPECT().SendMessage(mock.AnythingOfType("*sarama.ProducerMessage")).Return(0, 0, nil).Once()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		err := publishUseCase.Apply(context.Background(), req)

		sCtx.Assert().NoError(err)
	})
}

func (s *PublishRequestUseCaseSuite) TestPublishRequestUseCase_ApplyInvalidDate(t provider.T) {
	t.Title("Apply: InvalidDate")
	t.Tags("PublishRequestUseCase")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockPublishReqDepFields(t)

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(1).
			WithApplierID(12).
			WithManagerID(9).
			WithStatus(base.OnApprovalRequest).
			WithReleaseID(777).
			WithGrade(-3).
			WithExpexctedDate(cdtime.GetToday().AddDate(0, 0, 3)).
			Build()

		publishUseCase, _ := NewPublishRequestUseCase(df.publicationRepo, df.releaseRepo, df.artistRepo, df.transactor, df.pbBroker, df.publishRepo, df.logger)

		err := publishUseCase.Apply(context.Background(), req)

		sCtx.Assert().ErrorIs(err, pubReqErrors.ErrInvalidDate)
	})
}

func TestPublishRequestUseCaseSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(PublishRequestUseCaseSuite))
}
