//go:build unit
// +build unit

package service

import (
	"context"
	"cookdroogers/internal/repo/mocks"
	"cookdroogers/models"
	"cookdroogers/models/data_builders"
	cdtime "cookdroogers/pkg/time"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	publicationRepo *mocks.PublicationRepo
	logger          *slog.Logger
}

type PublicationServiceSuite struct {
	suite.Suite
}

func _newMockPublicationDepFields(t provider.T) *_depFields {
	mockPublicationRepo := mocks.NewPublicationRepo(t)

	f := &_depFields{
		publicationRepo: mockPublicationRepo,
		logger:          slog.Default(),
	}

	return f
}

func (s *PublicationServiceSuite) TestPublicationService_CreateOK(t provider.T) {
	t.Title("Create: OK")
	t.Tags("PublicationService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockPublicationDepFields(t)

		publication := data_builders.NewPublicationBuilder().Build()

		df.publicationRepo.EXPECT().Create(mock.Anything, publication).Return(nil).Once()

		publicationService := NewPublicationService(df.publicationRepo, df.logger)

		err := publicationService.Create(context.Background(), publication)

		sCtx.Assert().NoError(err)
	})
}

func (s *PublicationServiceSuite) TestPublicationService_CreateError(t provider.T) {
	t.Title("Create: Error")
	t.Tags("PublicationService")
	t.Parallel()

	t.WithNewStep("Repo error", func(sCtx provider.StepCtx) {
		df := _newMockPublicationDepFields(t)

		publication := data_builders.NewPublicationBuilder().Build()

		df.publicationRepo.EXPECT().Create(mock.Anything, publication).Return(fmt.Errorf("db error")).Once()

		publicationService := NewPublicationService(df.publicationRepo, df.logger)

		err := publicationService.Create(context.Background(), publication)

		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "can't create publication info")
	})
}

func (s *PublicationServiceSuite) TestPublicationService_GetOK(t provider.T) {
	t.Title("Get: OK")
	t.Tags("PublicationService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockPublicationDepFields(t)

		publication := data_builders.NewPublicationBuilder().WithPublicationID(88).Build()

		df.publicationRepo.EXPECT().Get(mock.Anything, uint64(88)).Return(publication, nil).Once()

		publicationService := NewPublicationService(df.publicationRepo, df.logger)

		result, err := publicationService.Get(context.Background(), uint64(88))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(publication, result)
	})
}

func (s *PublicationServiceSuite) TestPublicationService_GetError(t provider.T) {
	t.Title("Get: Error")
	t.Tags("PublicationService")
	t.Parallel()

	t.WithNewStep("Repo error", func(sCtx provider.StepCtx) {
		df := _newMockPublicationDepFields(t)

		df.publicationRepo.EXPECT().Get(mock.Anything, uint64(88)).Return(nil, fmt.Errorf("db error")).Once()

		publicationService := NewPublicationService(df.publicationRepo, df.logger)

		_, err := publicationService.Get(context.Background(), uint64(88))

		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "can't get publication info")
	})
}

func (s *PublicationServiceSuite) TestPublicationService_GetAllByDateOK(t provider.T) {
	t.Title("GetAllByDate: OK")
	t.Tags("PublicationService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockPublicationDepFields(t)

		publications := []models.Publication{
			*data_builders.NewPublicationBuilder().WithDate(cdtime.GetToday()).Build(),
		}

		df.publicationRepo.EXPECT().GetAllByDate(mock.Anything, cdtime.GetToday().AddDate(-1, 0, 0)).Return(publications, nil).Once()

		publicationService := NewPublicationService(df.publicationRepo, df.logger)

		result, err := publicationService.GetAllByDate(context.Background(), cdtime.GetToday().AddDate(-1, 0, 0))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(publications, result)
	})
}

func (s *PublicationServiceSuite) TestPublicationService_GetAllByDateError(t provider.T) {
	t.Title("GetAllByDate: Error")
	t.Tags("PublicationService")
	t.Parallel()

	t.WithNewStep("Repo error", func(sCtx provider.StepCtx) {
		df := _newMockPublicationDepFields(t)

		df.publicationRepo.EXPECT().GetAllByDate(mock.Anything, cdtime.GetToday().AddDate(-1, 0, 0)).Return(nil, fmt.Errorf("db error")).Once()

		publicationService := NewPublicationService(df.publicationRepo, df.logger)

		_, err := publicationService.GetAllByDate(context.Background(), cdtime.GetToday().AddDate(-1, 0, 0))

		sCtx.Assert().Error(err)
		sCtx.Assert().Contains(err.Error(), "can't get publications info")
	})
}

func (s *PublicationServiceSuite) TestPublicationService_GetAllByManagerOK(t provider.T) {
	t.Title("GetAllByManager: OK")
	t.Tags("PublicationService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockPublicationDepFields(t)

		publications := []models.Publication{
			*data_builders.NewPublicationBuilder().WithManagerID(8).Build(),
		}

		df.publicationRepo.EXPECT().GetAllByManager(mock.Anything, uint64(8)).Return(publications, nil).Once()

		publicationService := NewPublicationService(df.publicationRepo, df.logger)

		result, err := publicationService.GetAllByManager(context.Background(), uint64(8))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(publications, result)
	})
}

func (s *PublicationServiceSuite) TestPublicationService_GetAllByArtistSinceDateOK(t provider.T) {
	t.Title("GetAllByArtistSinceDate: OK")
	t.Tags("PublicationService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockPublicationDepFields(t)

		publications := []models.Publication{
			*data_builders.NewPublicationBuilder().WithManagerID(8).Build(),
		}

		df.publicationRepo.EXPECT().GetAllByArtistSinceDate(mock.Anything, cdtime.GetToday().AddDate(-1, 0, 0), uint64(8)).Return(publications, nil).Once()

		publicationService := NewPublicationService(df.publicationRepo, df.logger)

		result, err := publicationService.GetAllByArtistSinceDate(context.Background(), cdtime.GetToday().AddDate(-1, 0, 0), uint64(8))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(publications, result)
	})
}

func TestPublicationServiceSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(PublicationServiceSuite))
}
