package data_access

import (
	"context"
	"cookdroogers/integration_tests/utils"
	"cookdroogers/internal/repo"
	postgres "cookdroogers/internal/repo/pg"
	"cookdroogers/internal/transactor"
	transactor2 "cookdroogers/internal/transactor/trm"
	"cookdroogers/models/data_builders"
	"database/sql"
	"fmt"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	at_manager "github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"os"
)

type ManagerIntegrationPgSuite struct {
	suite.Suite
	db         *sql.DB
	dbx        *sqlx.DB
	txResolver *trmsqlx.CtxGetter
	trm        transactor.Transactor
	ctx        context.Context
	repo       repo.ManagerRepo
	builder    *data_builders.ManagerBuilder
}

func (s *ManagerIntegrationPgSuite) BeforeEach(t provider.T) {
	var err error
	pgInfo := utils.PostgresInfo{
		Host:     "postgres",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     "5432",
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	s.db, err = utils.InitDB(pgInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	if s.db == nil {
		return
	}

	s.ctx = context.Background()
	s.dbx = sqlx.NewDb(s.db, "pgx")

	s.repo = postgres.NewManagerPgRepo(s.db, s.trm)

	s.txResolver = trmsqlx.DefaultCtxGetter
	manager, _ := at_manager.New(trmsqlx.NewDefaultFactory(s.dbx))
	s.trm = transactor2.NewATtrm(manager)

	s.builder = data_builders.NewManagerBuilder()
}

func (s *ManagerIntegrationPgSuite) AfterEach(t provider.T) {
	_ = s.db.Close()
}

func (s *ManagerIntegrationPgSuite) TestManagerPgRepoGetSuccess(t provider.T) {
	t.Title("ManagerPgRepoGet: Success")
	t.Tags("ManagerIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Assign
			expectedManager := s.builder.WithManagerID(2).WithUserID(2).WithArtists(make([]uint64, 0)).Build()

			// Act
			manager, err := s.repo.Get(txCtx, uint64(2))

			// Assert
			sCtx.Assert().NoError(err)
			sCtx.Assert().Equal(expectedManager, manager)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ManagerIntegrationPgSuite) TestManagerPgRepoGetFailure(t provider.T) {
	t.Title("ManagerPgRepoGet: Failure")
	t.Tags("ManagerIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			_, err := s.repo.Get(txCtx, uint64(10))

			// Assert
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ManagerIntegrationPgSuite) TestManagerPgRepoGetByUserIDSuccess(t provider.T) {
	t.Title("ManagerPgRepoGetByUserID: Success")
	t.Tags("ManagerIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Assign
			expectedManager := s.builder.WithManagerID(2).WithUserID(2).WithArtists(make([]uint64, 0)).Build()

			// Act
			manager, err := s.repo.GetByUserID(txCtx, uint64(2))

			// Assert
			sCtx.Assert().NoError(err)
			sCtx.Assert().Equal(expectedManager, manager)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ManagerIntegrationPgSuite) TestManagerPgRepoGetByUserIDFailure(t provider.T) {
	t.Title("ManagerPgRepoGetByUserID: Failure")
	t.Tags("ManagerIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			_, err := s.repo.GetByUserID(txCtx, uint64(10))

			// Assert
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}
