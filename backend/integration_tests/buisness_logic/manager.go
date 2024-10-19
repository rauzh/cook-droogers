package buisness_logic

import (
	"context"
	"cookdroogers/integration_tests/utils"
	"cookdroogers/internal/manager/service"
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
	"log/slog"
	"os"
)

type ManagerIntegrationCoreSuite struct {
	suite.Suite
	db             *sql.DB
	dbx            *sqlx.DB
	txResolver     *trmsqlx.CtxGetter
	trm            transactor.Transactor
	managerService service.IManagerService
	ctx            context.Context
	builder        *data_builders.ManagerBuilder
}

func (s *ManagerIntegrationCoreSuite) BeforeEach(t provider.T) {
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

	managerRepo := postgres.NewManagerPgRepo(s.db, s.trm)

	s.txResolver = trmsqlx.DefaultCtxGetter
	manager, _ := at_manager.New(trmsqlx.NewDefaultFactory(s.dbx))
	s.trm = transactor2.NewATtrm(manager)

	s.managerService = service.NewManagerService(managerRepo, slog.Default())
}

func (s *ManagerIntegrationCoreSuite) AfterEach(t provider.T) {
	_ = s.db.Close()
}

func (s *ManagerIntegrationCoreSuite) TestManagerGetFailure(t provider.T) {
	t.Title("ManagerGet: Failure")
	t.Tags("ManagerIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			_, err := s.managerService.Get(txCtx, uint64(10))

			// Assert
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ManagerIntegrationCoreSuite) TestManagerGetByUserIDFailure(t provider.T) {
	t.Title("ManagerGetByUserID: Failure")
	t.Tags("ManagerIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			_, err := s.managerService.GetByUserID(txCtx, uint64(10))

			// Assert
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}
