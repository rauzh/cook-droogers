//go:build unit
// +build unit

package service

import (
	"context"
	releaseErrors "cookdroogers/internal/release/errors"
	"cookdroogers/internal/repo/mocks"
	"cookdroogers/internal/track/service"
	transac_mocks "cookdroogers/internal/transactor/mocks"
	"cookdroogers/models"
	"cookdroogers/models/data_builders"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

// _depFields содержит моки и logger
type _depFields struct {
	trackRepo   *mocks.TrackRepo
	releaseRepo *mocks.ReleaseRepo
	transactor  *transac_mocks.Transactor
	logger      *slog.Logger
}

// ReleaseServiceSuite определяет набор тестов для ReleaseService
type ReleaseServiceSuite struct {
	suite.Suite
}

// _newMockReleaseDepFields создает моки и зависимости для тестов
func _newMockReleaseDepFields(t provider.T) *_depFields {
	mockTrackRepo := mocks.NewTrackRepo(t)
	mockReleaseRepo := mocks.NewReleaseRepo(t)
	mockTransactor := transac_mocks.NewTransactor(t)

	f := &_depFields{
		trackRepo:   mockTrackRepo,
		releaseRepo: mockReleaseRepo,
		transactor:  mockTransactor,
		logger:      slog.Default(),
	}

	return f
}

func (s *ReleaseServiceSuite) TestReleaseService_CreateOK(t provider.T) {
	t.Title("Create: OK")
	t.Tags("ReleaseService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockReleaseDepFields(t)

		release := data_builders.NewReleaseBuilder().Build()
		tracks := []*models.Track{
			{TrackID: 1, Genre: "Pop"},
			{TrackID: 2, Genre: "Rock"},
		}

		df.transactor.EXPECT().WithinTransaction(mock.Anything, mock.Anything).Return(nil).Once()

		releaseService := NewReleaseService(service.NewTrackService(df.trackRepo, df.logger), df.transactor, df.releaseRepo, df.logger)

		err := releaseService.Create(context.Background(), release, tracks)

		sCtx.Assert().NoError(err)
	})
}

func (s *ReleaseServiceSuite) TestReleaseService_CreateValidationError(t provider.T) {
	t.Title("Create: Validation Error")
	t.Tags("ReleaseService")
	t.Parallel()

	t.WithNewStep("Validation error", func(sCtx provider.StepCtx) {
		df := _newMockReleaseDepFields(t)

		release := data_builders.NewReleaseBuilder().WithTitle("").Build()
		tracks := []*models.Track{}

		releaseService := NewReleaseService(service.NewTrackService(df.trackRepo, df.logger), df.transactor, df.releaseRepo, df.logger)

		err := releaseService.Create(context.Background(), release, tracks)

		sCtx.Assert().ErrorIs(err, releaseErrors.ErrNoTitle)
	})
}

func (s *ReleaseServiceSuite) TestReleaseService_GetOK(t provider.T) {
	t.Title("Get: OK")
	t.Tags("ReleaseService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockReleaseDepFields(t)

		release := data_builders.NewReleaseBuilder().WithReleaseID(888).Build()

		df.releaseRepo.EXPECT().Get(mock.Anything, uint64(888)).Return(release, nil).Once()

		releaseService := NewReleaseService(service.NewTrackService(df.trackRepo, df.logger), df.transactor, df.releaseRepo, df.logger)

		result, err := releaseService.Get(context.Background(), uint64(888))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(release, result)
	})
}

func (s *ReleaseServiceSuite) TestReleaseService_GetAllByArtistOK(t provider.T) {
	t.Title("GetAllByArtist: OK")
	t.Tags("ReleaseService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockReleaseDepFields(t)

		releases := []models.Release{*data_builders.NewReleaseBuilder().WithArtistID(7).Build()}

		df.releaseRepo.EXPECT().GetAllByArtist(mock.Anything, uint64(7)).Return(releases, nil).Once()

		releaseService := NewReleaseService(service.NewTrackService(df.trackRepo, df.logger), df.transactor, df.releaseRepo, df.logger)

		result, err := releaseService.GetAllByArtist(context.Background(), uint64(7))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(releases, result)
	})
}

// TestReleaseService_GetAllTracksOK тестирует метод GetAllTracks в успешном сценарии
func (s *ReleaseServiceSuite) TestReleaseService_GetAllTracksOK(t provider.T) {
	t.Title("GetAllTracks: OK")
	t.Tags("ReleaseService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockReleaseDepFields(t)

		release := data_builders.NewReleaseBuilder().WithTracks([]uint64{1111}).Build()
		tracks := []models.Track{
			{TrackID: 1111, Genre: "Pop"},
		}

		df.releaseRepo.EXPECT().GetAllTracks(mock.Anything, release).Return(tracks, nil).Once()

		releaseService := NewReleaseService(service.NewTrackService(df.trackRepo, df.logger), df.transactor, df.releaseRepo, df.logger)

		result, err := releaseService.GetAllTracks(context.Background(), release)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(tracks, result)
	})
}

func (s *ReleaseServiceSuite) TestReleaseService_GetMainGenreOK(t provider.T) {
	t.Title("GetMainGenre: OK")
	t.Tags("ReleaseService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockReleaseDepFields(t)

		release := data_builders.NewReleaseBuilder().WithTracks([]uint64{1111, 2222}).Build()
		df.releaseRepo.EXPECT().Get(mock.Anything, uint64(888)).Return(release, nil).Once()

		df.trackRepo.EXPECT().Get(mock.Anything, uint64(1111)).Return(&models.Track{TrackID: 1111, Genre: "Pop"}, nil).Once()
		df.trackRepo.EXPECT().Get(mock.Anything, uint64(2222)).Return(&models.Track{TrackID: 2222, Genre: "Pop"}, nil).Once()

		releaseService := NewReleaseService(service.NewTrackService(df.trackRepo, df.logger), df.transactor, df.releaseRepo, df.logger)

		mainGenre, err := releaseService.GetMainGenre(context.Background(), uint64(888))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal("Pop", mainGenre)
	})
}

func TestReleaseServiceSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ReleaseServiceSuite))
}
