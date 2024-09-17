package usecase

import (
	rlsService "cookdroogers/internal/release/service"
	"cookdroogers/internal/repo/mocks"
	"cookdroogers/internal/requests/base"
	base_errors "cookdroogers/internal/requests/base/errors"
	broker_mocks "cookdroogers/internal/requests/broker/mocks"
	"cookdroogers/internal/requests/publish"
	pubReqErrors "cookdroogers/internal/requests/publish/errors"
	publishReqRepoMocks "cookdroogers/internal/requests/publish/repo/mocks"
	statFetcher "cookdroogers/internal/statistics/fetcher/mocks"
	statService "cookdroogers/internal/statistics/service"
	trackService "cookdroogers/internal/track/service"
	transacMock "cookdroogers/internal/transactor/mocks"
	cdtime "cookdroogers/pkg/time"
	"errors"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	_trackRepo *mocks.TrackRepo
	_statRepo  *mocks.StatisticsRepo

	statService     statService.IStatisticsService
	publicationRepo *mocks.PublicationRepo
	releaseRepo     *mocks.ReleaseRepo
	artistRepo      *mocks.ArtistRepo
	transactor      *transacMock.Transactor
	pbBroker        *broker_mocks.IBroker

	publishRepo *publishReqRepoMocks.PublishRequestRepo
}

func _newMockPublishReqDepFields(t *testing.T) *_depFields {
	transactionMock := transacMock.NewTransactor(t)
	pbcMockRepo := mocks.NewPublicationRepo(t)
	rlsMockRepo := mocks.NewReleaseRepo(t)
	artistMockRepo := mocks.NewArtistRepo(t)
	trkMockRepo := mocks.NewTrackRepo(t)
	statMockRepo := mocks.NewStatisticsRepo(t)
	publishMockRepo := publishReqRepoMocks.NewPublishRequestRepo(t)

	statMockFetcher := statFetcher.NewStatFetcher(t)

	trkSvc := trackService.NewTrackService(trkMockRepo, slog.Default())
	rlsSvc := rlsService.NewReleaseService(trkSvc, transactionMock, rlsMockRepo, slog.Default())
	statSvc := statService.NewStatisticsService(trkSvc, statMockFetcher, statMockRepo, rlsSvc, slog.Default())

	mockBroker := broker_mocks.NewIBroker(t)

	f := &_depFields{
		_statRepo:       statMockRepo,
		_trackRepo:      trkMockRepo,
		statService:     statSvc,
		publicationRepo: pbcMockRepo,
		releaseRepo:     rlsMockRepo,
		artistRepo:      artistMockRepo,
		transactor:      transactionMock,
		pbBroker:        mockBroker,
		publishRepo:     publishMockRepo,
	}

	return f
}

func TestPublishRequestUseCase_Decline(t *testing.T) {

	type args struct {
		pubReq *publish.PublishRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields)
	}{
		{
			name: "OK",
			in: &args{
				pubReq: &publish.PublishRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      publish.PubReq,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					ReleaseID:    777,
					Grade:        -3,
					ExpectedDate: cdtime.GetToday().AddDate(1, 0, 0),
					Description:  mock.Anything,
				},
			},
			out: nil,
			dependencies: func(df *_depFields) {

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
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
		{
			name: "InvalidDate",
			in: &args{
				pubReq: &publish.PublishRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      publish.PubReq,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					ReleaseID:    777,
					Grade:        -3,
					ExpectedDate: cdtime.GetToday().AddDate(0, 0, 3),
					Description:  mock.Anything,
				},
			},
			out: pubReqErrors.ErrInvalidDate,
			dependencies: func(df *_depFields) {

			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockPublishReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			publishReqUseCase, err := NewPublishRequestUseCase(f.statService, f.publicationRepo, f.releaseRepo, f.artistRepo, f.transactor, f.pbBroker, f.publishRepo, slog.Default())

			// act
			err = publishReqUseCase.Decline(tt.in.pubReq)

			// assert
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
			if tt.assert != nil {
				tt.assert(t, f)
			}
		})
	}
}

func TestPublishRequestUseCase_Accept(t *testing.T) {

	type args struct {
		pubReq *publish.PublishRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields)
	}{
		{
			name: "OK",
			in: &args{
				pubReq: &publish.PublishRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      publish.PubReq,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					ReleaseID:    777,
					Grade:        -3,
					ExpectedDate: cdtime.GetToday().AddDate(1, 0, 0),
					Description:  mock.Anything,
				},
			},
			out: nil,
			dependencies: func(df *_depFields) {
				df.transactor.EXPECT().WithinTransaction(mock.AnythingOfType("context.backgroundCtx"),
					mock.Anything).Return(nil).Once()
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
		{
			name: "InvalidType",
			in: &args{
				pubReq: &publish.PublishRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      "",
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					ReleaseID:    777,
					Grade:        -3,
					ExpectedDate: cdtime.GetToday().AddDate(1, 0, 3),
					Description:  mock.Anything,
				},
			},
			out: base_errors.ErrInvalidType,
			dependencies: func(df *_depFields) {

			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
		{
			name: "InvalidAlreadyClosed",
			in: &args{
				pubReq: &publish.PublishRequest{
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
					ExpectedDate: cdtime.GetToday().AddDate(1, 0, 3),
					Description:  mock.Anything,
				},
			},
			out: base_errors.ErrAlreadyClosed,
			dependencies: func(df *_depFields) {
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockPublishReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			publishReqUseCase, err := NewPublishRequestUseCase(f.statService, f.publicationRepo, f.releaseRepo, f.artistRepo, f.transactor, f.pbBroker, f.publishRepo, slog.Default())

			// act
			err = publishReqUseCase.Accept(tt.in.pubReq)

			// assert
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
			if tt.assert != nil {
				tt.assert(t, f)
			}
		})
	}
}

func TestPublishRequestUseCase_Apply(t *testing.T) {

	type args struct {
		pubReq *publish.PublishRequest
	}

	tests := []struct {
		name string
		in   *args
		out  error

		dependencies func(*_depFields)
		assert       func(*testing.T, *_depFields)
	}{
		//{
		//	name: "OK",
		//	in: &args{
		//		pubReq: &publish.PublishRequest{
		//			Request: base.Request{
		//				RequestID: 1,
		//				Type:      publish.PubReq,
		//				Status:    base.OnApprovalRequest,
		//				Date:      cdtime.GetToday(),
		//				ApplierID: 12,
		//				ManagerID: 0,
		//			},
		//			ReleaseID:    777,
		//			Grade:        -3,
		//			ExpectedDate: cdtime.GetToday().AddDate(1, 0, 0),
		//			Description:  mock.Anything,
		//		},
		//	},
		//	out: nil,
		//	dependencies: func(df *_depFields) {
		//
		//		df.pbBroker.EXPECT().SendMessage(mock.Anything).Return(0, 0, nil)
		//
		//		df.publishRepo.EXPECT().Create(mock.AnythingOfType("context.backgroundCtx"), &publish.PublishRequest{
		//			Request: base.Request{
		//				RequestID: 1,
		//				Type:      publish.PubReq,
		//				Status:    base.NewRequest,
		//				Date:      cdtime.GetToday(),
		//				ApplierID: 12,
		//				ManagerID: 0,
		//			},
		//			ReleaseID:    777,
		//			Grade:        -3,
		//			ExpectedDate: cdtime.GetToday().AddDate(1, 0, 0),
		//			Description:  mock.Anything,
		//		}).Return(nil).Once()
		//	},
		//	assert: func(t *testing.T, df *_depFields) {
		//
		//	},
		//},
		{
			name: "InvalidDate",
			in: &args{
				pubReq: &publish.PublishRequest{
					Request: base.Request{
						RequestID: 1,
						Type:      publish.PubReq,
						Status:    base.OnApprovalRequest,
						Date:      cdtime.GetToday(),
						ApplierID: 12,
						ManagerID: 9,
					},
					ReleaseID:    777,
					Grade:        -3,
					ExpectedDate: cdtime.GetToday().AddDate(0, 0, 3),
					Description:  mock.Anything,
				},
			},
			out: pubReqErrors.ErrInvalidDate,
			dependencies: func(df *_depFields) {
			},
			assert: func(t *testing.T, df *_depFields) {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f := _newMockPublishReqDepFields(t)
			if tt.dependencies != nil {
				tt.dependencies(f)
			}

			publishReqUseCase, err := NewPublishRequestUseCase(f.statService, f.publicationRepo, f.releaseRepo, f.artistRepo, f.transactor, f.pbBroker, f.publishRepo, slog.Default())

			// act
			err = publishReqUseCase.Apply(tt.in.pubReq)

			// assert
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
			if tt.assert != nil {
				tt.assert(t, f)
			}
		})
	}
}
