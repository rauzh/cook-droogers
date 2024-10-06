package service

import (
	"context"
	"cookdroogers/internal/repo/mocks"
	"cookdroogers/models/data_builders"
	"database/sql"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	artistRepo *mocks.ArtistRepo
	logger     *slog.Logger
}

type ArtistServiceSuite struct {
	suite.Suite
}

func _newMockArtistDepFields(t provider.T) *_depFields {
	mockArtistRepo := mocks.NewArtistRepo(t)

	f := &_depFields{
		artistRepo: mockArtistRepo,
		logger:     slog.Default(),
	}

	return f
}

func (s *ArtistServiceSuite) TestArtistService_CreateOK(t provider.T) {
	t.Title("Create: OK")
	t.Tags("ArtistService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockArtistDepFields(t)

		artist := data_builders.NewArtistBuilder().Build()

		df.artistRepo.EXPECT().Create(mock.Anything, artist).Return(nil).Once()

		artistService := NewArtistService(df.artistRepo, df.logger)

		err := artistService.Create(context.Background(), artist)

		sCtx.Assert().NoError(err)
	})
}

func (s *ArtistServiceSuite) TestArtistService_CreateDbErr(t provider.T) {
	t.Title("Create: DB error")
	t.Tags("ArtistService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockArtistDepFields(t)

		artist := data_builders.NewArtistBuilder().Build()

		df.artistRepo.EXPECT().Create(mock.Anything, artist).Return(sql.ErrConnDone).Once()

		artistService := NewArtistService(df.artistRepo, df.logger)

		err := artistService.Create(context.Background(), artist)

		sCtx.Assert().ErrorIs(err, CreateDbError)
	})
}

func (s *ArtistServiceSuite) TestArtistService_GetOK(t provider.T) {
	t.Title("Get: OK")
	t.Tags("ArtistService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockArtistDepFields(t)

		artist := data_builders.NewArtistBuilder().WithID(1).Build()

		df.artistRepo.EXPECT().Get(mock.Anything, uint64(1)).Return(artist, nil).Once()

		artistService := NewArtistService(df.artistRepo, df.logger)

		result, err := artistService.Get(context.Background(), uint64(1))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(artist, result)
	})
}

func (s *ArtistServiceSuite) TestArtistService_GetDbErr(t provider.T) {
	t.Title("Get: DB error")
	t.Tags("ArtistService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockArtistDepFields(t)

		df.artistRepo.EXPECT().Get(mock.Anything, uint64(1)).Return(nil, sql.ErrConnDone).Once()

		artistService := NewArtistService(df.artistRepo, df.logger)

		result, err := artistService.Get(context.Background(), uint64(1))

		sCtx.Assert().Nil(result)
		sCtx.Assert().ErrorIs(err, GetDbError)
	})
}

func (s *ArtistServiceSuite) TestArtistService_GetByUserIDOK(t provider.T) {
	t.Title("GetByUserID: OK")
	t.Tags("ArtistService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockArtistDepFields(t)

		artist := data_builders.NewArtistBuilder().WithUserID(7).Build()

		df.artistRepo.EXPECT().GetByUserID(mock.Anything, uint64(7)).Return(artist, nil).Once()

		artistService := NewArtistService(df.artistRepo, df.logger)

		result, err := artistService.GetByUserID(context.Background(), uint64(7))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(artist, result)
	})
}

func (s *ArtistServiceSuite) TestArtistService_GetByUserIDDbErr(t provider.T) {
	t.Title("GetByUserID: DB error")
	t.Tags("ArtistService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockArtistDepFields(t)

		df.artistRepo.EXPECT().GetByUserID(mock.Anything, uint64(7)).Return(nil, sql.ErrConnDone).Once()

		artistService := NewArtistService(df.artistRepo, df.logger)

		result, err := artistService.GetByUserID(context.Background(), uint64(7))

		sCtx.Assert().Nil(result)
		sCtx.Assert().ErrorIs(err, GetDbError)
	})
}

func (s *ArtistServiceSuite) TestArtistService_UpdateOK(t provider.T) {
	t.Title("Update: OK")
	t.Tags("ArtistService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockArtistDepFields(t)

		artist := data_builders.NewArtistBuilder().WithID(1).WithNickname("updated_nickname").Build()

		df.artistRepo.EXPECT().Update(mock.Anything, artist).Return(nil).Once()

		artistService := NewArtistService(df.artistRepo, df.logger)

		err := artistService.Update(context.Background(), artist)

		sCtx.Assert().NoError(err)
	})
}

func (s *ArtistServiceSuite) TestArtistService_UpdateDbErr(t provider.T) {
	t.Title("Update: DB error")
	t.Tags("ArtistService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockArtistDepFields(t)

		artist := data_builders.NewArtistBuilder().WithID(1).WithNickname("updated_nickname").Build()

		df.artistRepo.EXPECT().Update(mock.Anything, artist).Return(sql.ErrConnDone).Once()

		artistService := NewArtistService(df.artistRepo, df.logger)

		err := artistService.Update(context.Background(), artist)

		sCtx.Assert().ErrorIs(err, UpdateDbError)
	})
}

func TestArtistServiceSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ArtistServiceSuite))
}
