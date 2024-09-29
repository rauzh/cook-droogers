package publish

//
//import (
//	rlsService "cookdroogers/internal/release/service"
//	"cookdroogers/internal/repo/mocks"
//	"cookdroogers/internal/requests/base"
//	broker_mocks "cookdroogers/internal/requests/broker/mocks"
//	criteria "cookdroogers/internal/requests/criteria_controller"
//	publish_criteria "cookdroogers/internal/requests/criteria_controller/publish"
//	"cookdroogers/internal/requests/publish"
//	pubReqErrors "cookdroogers/internal/requests/publish/errors"
//	publishReqRepoMocks "cookdroogers/internal/requests/publish/repo/mocks"
//	statFetcher "cookdroogers/internal/statistics/fetcher/mocks"
//	statService "cookdroogers/internal/statistics/service"
//	trackService "cookdroogers/internal/track/service"
//	transacMock "cookdroogers/internal/transactor/mocks"
//	"cookdroogers/models"
//	cdtime "cookdroogers/pkg/time"
//	"errors"
//	"log/slog"
//	"testing"
//
//	"github.com/stretchr/testify/mock"
//)
//
//type _depFields struct {
//	_trackRepo *mocks.TrackRepo
//	_statRepo  *mocks.StatisticsRepo
//
//	statService     statService.IStatisticsService
//	publicationRepo *mocks.PublicationRepo
//	releaseRepo     *mocks.ReleaseRepo
//	artistRepo      *mocks.ArtistRepo
//
//	pbBroker  *broker_mocks.IBroker
//	criterias criteria.ICriteriaCollection
//
//	publishRepo *publishReqRepoMocks.PublishRequestRepo
//}
//
//func _newMockPublishReqDepFields(t *testing.T) *_depFields {
//	transactionMock := transacMock.NewTransactor(t)
//	mockArtRepo := mocks.NewArtistRepo(t)
//	pbcMockRepo := mocks.NewPublicationRepo(t)
//	rlsMockRepo := mocks.NewReleaseRepo(t)
//	trkMockRepo := mocks.NewTrackRepo(t)
//	statMockRepo := mocks.NewStatisticsRepo(t)
//	publishMockRepo := publishReqRepoMocks.NewPublishRequestRepo(t)
//
//	statMockFetcher := statFetcher.NewStatFetcher(t)
//
//	trkSvc := trackService.NewTrackService(trkMockRepo, slog.Default())
//	rlsSvc := rlsService.NewReleaseService(trkSvc, transactionMock, rlsMockRepo, slog.Default())
//	statSvc := statService.NewStatisticsService(trkSvc, statMockFetcher, statMockRepo, rlsSvc, slog.Default())
//
//	critCollection, _ := criteria.BuildCollection(
//		&publish_criteria.ArtistReleaseLimitPerSeasonCriteriaFabric{PublicationRepo: pbcMockRepo, ArtistRepo: mockArtRepo},
//		&publish_criteria.RelevantGenreCriteriaFabric{ReleaseService: rlsSvc, StatService: statSvc},
//		&publish_criteria.OneReleasePerDayCriteriaFabric{PublicationRepo: pbcMockRepo})
//
//	mockBroker := broker_mocks.NewIBroker(t)
//
//	f := &_depFields{
//		_statRepo:       statMockRepo,
//		_trackRepo:      trkMockRepo,
//		statService:     statSvc,
//		publicationRepo: pbcMockRepo,
//		releaseRepo:     rlsMockRepo,
//		artistRepo:      mockArtRepo,
//		pbBroker:        mockBroker,
//		criterias:       critCollection,
//		publishRepo:     publishMockRepo,
//	}
//
//	return f
//}
//
//func TestPublishRequestUseCase_proceedToManager(t *testing.T) {
//
//	type args struct {
//		pubReq *publish.PublishRequest
//	}
//
//	tests := []struct {
//		name string
//		in   *args
//		out  error
//
//		dependencies func(*_depFields)
//		assert       func(*testing.T, *_depFields)
//	}{
//		{
//			name: "OK",
//			in: &args{
//				pubReq: &publish.PublishRequest{
//					Request: base.Request{
//						RequestID: 1,
//						Type:      publish.PubReq,
//						Status:    base.OnApprovalRequest,
//						Date:      cdtime.GetToday(),
//						ApplierID: 12,
//						ManagerID: 0,
//					},
//					ReleaseID:    777,
//					Grade:        0,
//					ExpectedDate: cdtime.GetToday().AddDate(1, 0, 0),
//					Description:  "",
//				},
//			},
//			out: nil,
//			dependencies: func(df *_depFields) {
//				df.publishRepo.EXPECT().Update(mock.AnythingOfType("context.backgroundCtx"), &publish.PublishRequest{
//					Request: base.Request{
//						RequestID: 1,
//						Type:      publish.PubReq,
//						Status:    base.ProcessingRequest,
//						Date:      cdtime.GetToday(),
//						ApplierID: 12,
//						ManagerID: 0,
//					},
//					ReleaseID:    777,
//					Grade:        0,
//					ExpectedDate: cdtime.GetToday().AddDate(1, 0, 0),
//					Description:  "",
//				}).Return(nil).Once()
//
//				df.artistRepo.EXPECT().GetByUserID(mock.AnythingOfType("context.backgroundCtx"), uint64(12)).Return(
//					&models.Artist{ManagerID: 9, ArtistID: 199}, nil).Once()
//
//				df.publicationRepo.EXPECT().GetAllByArtistSinceDate(mock.AnythingOfType("context.backgroundCtx"),
//					cdtime.RelevantPeriod(), uint64(199)).Return([]models.Publication{{}, {}, {}, {}}, nil).Once()
//
//				df.publicationRepo.EXPECT().GetAllByDate(mock.AnythingOfType("context.backgroundCtx"),
//					cdtime.GetToday().AddDate(1, 0, 0)).Return([]models.Publication{}, nil).Once()
//
//				df.releaseRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(777)).Return(
//					&models.Release{Tracks: []uint64{999, 721}}, nil).Once()
//
//				df._trackRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(999)).Return(
//					&models.Track{Genre: "rock"}, nil).Once()
//
//				df._trackRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(721)).Return(
//					&models.Track{Genre: "rock"}, nil).Once()
//
//				df._statRepo.EXPECT().GetAllGroupByTracksSince(mock.AnythingOfType("context.backgroundCtx"),
//					cdtime.RelevantPeriod()).Return(nil, pubReqErrors.ErrInvalidDate).Once()
//
//				df.artistRepo.EXPECT().Get(mock.AnythingOfType("context.backgroundCtx"), uint64(12)).Return(
//					&models.Artist{ManagerID: 9}, nil).Once()
//
//				df.publishRepo.EXPECT().Update(mock.AnythingOfType("context.backgroundCtx"), &publish.PublishRequest{
//					Request: base.Request{
//						RequestID: 1,
//						Type:      publish.PubReq,
//						Status:    base.OnApprovalRequest,
//						Date:      cdtime.GetToday(),
//						ApplierID: 12,
//						ManagerID: 9,
//					},
//					ReleaseID:    777,
//					Grade:        -1,
//					ExpectedDate: cdtime.GetToday().AddDate(1, 0, 0),
//					Description:  "**No releases from artist more than limit** diff: -1\n**No releases from artist more than limit** reason: More than limit releases per season**Genre should be relevant** diff: 0\n**Genre should be relevant** reason: Can't apply criteria**No releases that day** diff: 0\n**No releases that day** reason: OK",
//				}).Return(nil).Once()
//			},
//			assert: func(t *testing.T, df *_depFields) {
//
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			f := _newMockPublishReqDepFields(t)
//			if tt.dependencies != nil {
//				tt.dependencies(f)
//			}
//
//			publishReqHandler := InitPublishProceedToManagerConsumerHandler(f.pbBroker, f.publishRepo, f.artistRepo, f.criterias)
//
//			// act
//			err := publishReqHandler.(*PublishProceedToManagerConsumerHandler).proceedToManager(tt.in.pubReq)
//
//			// assert
//			if !errors.Is(err, tt.out) {
//				t.Errorf("got %v, want %v", err, tt.out)
//			}
//			if tt.assert != nil {
//				tt.assert(t, f)
//			}
//		})
//	}
//}
