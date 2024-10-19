//go:build unit
// +build unit

package service

import (
	"context"
	"cookdroogers/internal/repo/mocks"
	trackErrors "cookdroogers/internal/track/errors"
	"cookdroogers/models/data_builders"
	"fmt"
	"log/slog"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
)

type _depFields struct {
	trackRepo *mocks.TrackRepo
	logger    *slog.Logger
}

type TrackServiceSuite struct {
	suite.Suite
}

func _newMockTrackDepFields(t provider.T) *_depFields {
	mockTrackRepo := mocks.NewTrackRepo(t)

	f := &_depFields{
		trackRepo: mockTrackRepo,
		logger:    slog.Default(),
	}

	return f
}

func (s *TrackServiceSuite) TestTrackService_CreateOK(t provider.T) {
	t.Title("Create: OK")
	t.Tags("TrackService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockTrackDepFields(t)

		track := data_builders.NewTrackBuilder().WithArtists([]uint64{7}).Build()

		df.trackRepo.EXPECT().Create(mock.Anything, track).Return(uint64(1111), nil).Once()

		trackService := NewTrackService(df.trackRepo, df.logger)

		trackID, err := trackService.Create(context.Background(), track)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(uint64(1111), trackID)
	})
}

func (s *TrackServiceSuite) TestTrackService_CreateValidationError(t provider.T) {
	t.Title("Create: Validation Error")
	t.Tags("TrackService")
	t.Parallel()

	t.WithNewStep("Validation error", func(sCtx provider.StepCtx) {
		df := _newMockTrackDepFields(t)

		track := data_builders.NewTrackBuilder().Build()
		track.Genre = ""

		trackService := NewTrackService(df.trackRepo, df.logger)

		_, err := trackService.Create(context.Background(), track)

		sCtx.Assert().Error(err)
		sCtx.Assert().ErrorIs(trackErrors.ErrNoGenre, err)
	})
}

func (s *TrackServiceSuite) TestTrackService_CreateRepoError(t provider.T) {
	t.Title("Create: Repo Error")
	t.Tags("TrackService")
	t.Parallel()

	t.WithNewStep("Repo error", func(sCtx provider.StepCtx) {
		df := _newMockTrackDepFields(t)

		track := data_builders.NewTrackBuilder().WithArtists([]uint64{7}).Build()

		df.trackRepo.EXPECT().Create(mock.Anything, track).Return(uint64(0), fmt.Errorf("db error")).Once()

		trackService := NewTrackService(df.trackRepo, df.logger)

		_, err := trackService.Create(context.Background(), track)

		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "can't create track")
	})
}

func (s *TrackServiceSuite) TestTrackService_GetOK(t provider.T) {
	t.Title("Get: OK")
	t.Tags("TrackService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockTrackDepFields(t)

		track := data_builders.NewTrackBuilder().Build()

		df.trackRepo.EXPECT().Get(mock.Anything, uint64(1111)).Return(track, nil).Once()

		trackService := NewTrackService(df.trackRepo, df.logger)

		result, err := trackService.Get(context.Background(), uint64(1111))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(track, result)
	})
}

func (s *TrackServiceSuite) TestTrackService_GetError(t provider.T) {
	t.Title("Get: Repo Error")
	t.Tags("TrackService")
	t.Parallel()

	t.WithNewStep("Repo error", func(sCtx provider.StepCtx) {
		df := _newMockTrackDepFields(t)

		df.trackRepo.EXPECT().Get(mock.Anything, uint64(1111)).Return(nil, fmt.Errorf("db error")).Once()

		trackService := NewTrackService(df.trackRepo, df.logger)

		_, err := trackService.Get(context.Background(), uint64(1111))

		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "can't get track")
	})
}

func TestTrackServiceSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(TrackServiceSuite))
}
