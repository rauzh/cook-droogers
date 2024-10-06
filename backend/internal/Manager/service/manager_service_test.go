package service

import (
	"context"
	"cookdroogers/internal/repo/mocks"
	"cookdroogers/models"
	"cookdroogers/models/data_builders"
	"errors"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	managerRepo *mocks.ManagerRepo
	logger      *slog.Logger
}

type ManagerServiceSuite struct {
	suite.Suite
}

func _newMockManagerDepFields(t provider.T) *_depFields {
	mockManagerRepo := mocks.NewManagerRepo(t)

	f := &_depFields{
		managerRepo: mockManagerRepo,
		logger:      slog.Default(),
	}

	return f
}

func (s *ManagerServiceSuite) TestManagerService_CreateOK(t provider.T) {
	t.Title("Create: OK")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockManagerDepFields(t)

		manager := data_builders.NewManagerBuilder().Build()

		df.managerRepo.EXPECT().Create(mock.Anything, manager).Return(nil).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		err := managerService.Create(context.Background(), manager)

		sCtx.Assert().NoError(err)
	})
}

func (s *ManagerServiceSuite) TestManagerService_CreateDbErr(t provider.T) {
	t.Title("Create: DB error")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		someDBErr := errors.New("some db error")
		df := _newMockManagerDepFields(t)

		manager := data_builders.NewManagerBuilder().Build()

		df.managerRepo.EXPECT().Create(mock.Anything, manager).Return(someDBErr).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		err := managerService.Create(context.Background(), manager)

		sCtx.Assert().ErrorIs(err, CreateDbError)
	})
}

func (s *ManagerServiceSuite) TestManagerService_GetOK(t provider.T) {
	t.Title("Get: OK")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockManagerDepFields(t)

		manager := data_builders.NewManagerBuilder().WithManagerID(1).Build()

		df.managerRepo.EXPECT().Get(mock.Anything, uint64(1)).Return(manager, nil).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		result, err := managerService.Get(context.Background(), uint64(1))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(manager, result)
	})
}

func (s *ManagerServiceSuite) TestManagerService_GetDbErr(t provider.T) {
	t.Title("Get: DB error")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		someDBErr := errors.New("some db error")
		df := _newMockManagerDepFields(t)

		df.managerRepo.EXPECT().Get(mock.Anything, uint64(1)).Return(nil, someDBErr).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		result, err := managerService.Get(context.Background(), uint64(1))

		sCtx.Assert().Nil(result)
		sCtx.Assert().ErrorIs(err, GetDbError)
	})
}

func (s *ManagerServiceSuite) TestManagerService_GetByUserIDOK(t provider.T) {
	t.Title("GetByUserID: OK")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockManagerDepFields(t)

		manager := data_builders.NewManagerBuilder().WithUserID(88).Build()

		df.managerRepo.EXPECT().GetByUserID(mock.Anything, uint64(88)).Return(manager, nil).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		result, err := managerService.GetByUserID(context.Background(), uint64(88))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(manager, result)
	})
}

func (s *ManagerServiceSuite) TestManagerService_GetByUserIDDbErr(t provider.T) {
	t.Title("GetByUserID: DB error")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		someDBErr := errors.New("some db error")
		df := _newMockManagerDepFields(t)

		df.managerRepo.EXPECT().GetByUserID(mock.Anything, uint64(88)).Return(nil, someDBErr).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		result, err := managerService.GetByUserID(context.Background(), uint64(88))

		sCtx.Assert().Nil(result)
		sCtx.Assert().ErrorIs(err, GetDbError)
	})
}

func (s *ManagerServiceSuite) TestManagerService_GetForAdminOK(t provider.T) {
	t.Title("GetForAdmin: OK")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockManagerDepFields(t)

		managers := []models.Manager{*data_builders.NewManagerBuilder().WithManagerID(8).Build()}

		df.managerRepo.EXPECT().GetForAdmin(mock.Anything).Return(managers, nil).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		result, err := managerService.GetForAdmin(context.Background())

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(managers, result)
	})
}

func (s *ManagerServiceSuite) TestManagerService_GetForAdminDbErr(t provider.T) {
	t.Title("GetForAdmin: DB error")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		someDBErr := errors.New("some db error")
		df := _newMockManagerDepFields(t)

		df.managerRepo.EXPECT().GetForAdmin(mock.Anything).Return(nil, someDBErr).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		result, err := managerService.GetForAdmin(context.Background())

		sCtx.Assert().Nil(result)
		sCtx.Assert().ErrorIs(err, GetDbError)
	})
}

func (s *ManagerServiceSuite) TestManagerService_GetRandomManagerIDOK(t provider.T) {
	t.Title("GetRandomManagerID: OK")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockManagerDepFields(t)

		df.managerRepo.EXPECT().GetRandManagerID(mock.Anything).Return(uint64(8), nil).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		id, err := managerService.GetRandomManagerID(context.Background())

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(uint64(8), id)
	})
}

func (s *ManagerServiceSuite) TestManagerService_GetRandomManagerIDDbErr(t provider.T) {
	t.Title("GetRandomManagerID: DB error")
	t.Tags("ManagerService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		someDBErr := errors.New("some db error")
		df := _newMockManagerDepFields(t)

		df.managerRepo.EXPECT().GetRandManagerID(mock.Anything).Return(uint64(0), someDBErr).Once()

		managerService := NewManagerService(df.managerRepo, df.logger)

		id, err := managerService.GetRandomManagerID(context.Background())

		sCtx.Assert().Equal(uint64(0), id)
		sCtx.Assert().ErrorIs(err, GetDbError)
	})
}

func TestManagerServiceSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ManagerServiceSuite))
}
