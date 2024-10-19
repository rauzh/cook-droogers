package data_access

import (
	"context"
	"cookdroogers/integration_tests/utils"
	"cookdroogers/internal/repo"
	postgres "cookdroogers/internal/repo/pg"
	"cookdroogers/internal/transactor"
	transactor2 "cookdroogers/internal/transactor/trm"
	"cookdroogers/models"
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

type UserIntegrationPgSuite struct {
	suite.Suite
	db         *sql.DB
	dbx        *sqlx.DB
	txResolver *trmsqlx.CtxGetter
	trm        transactor.Transactor
	userRepo   repo.UserRepo
	ctx        context.Context
}

func (s *UserIntegrationPgSuite) BeforeEach(t provider.T) {
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

	s.userRepo = postgres.NewUserPgRepo(s.db)

	s.txResolver = trmsqlx.DefaultCtxGetter
	manager, _ := at_manager.New(trmsqlx.NewDefaultFactory(s.dbx))
	s.trm = transactor2.NewATtrm(manager)
}

func (s *UserIntegrationPgSuite) AfterEach(t provider.T) {
	_ = s.db.Close()
}

func (s *UserIntegrationPgSuite) TestCreateUserSuccess(t provider.T) {
	t.Title("CreateUser: Success")
	t.Tags("UserIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			err := s.userRepo.Create(txCtx,
				data_builders.NewUserBuilder().WithUserID(0).WithEmail("hehehehehe@mail.ru").Build())

			// Assert
			sCtx.Assert().NoError(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *UserIntegrationPgSuite) TestGetUserSuccess(t provider.T) {
	t.Title("GetUser: Success")
	t.Tags("UserIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			user, err := s.userRepo.Get(txCtx, uint64(1))

			// Assert
			sCtx.Assert().Nil(err)
			sCtx.Assert().Equal(data_builders.NewUserBuilder().
				WithUserID(1).
				WithEmail("pavel@ppo.ru").
				WithType(models.ManagerUser).
				WithPassword("123123").
				WithName("pavel-manager").Build(), user)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *UserIntegrationPgSuite) TestGetUserFailure(t provider.T) {
	t.Title("GetUser: Failure")
	t.Tags("UserIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			User, err := s.userRepo.Get(txCtx, uint64(1000))

			// Assert
			sCtx.Assert().Nil(User)
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *UserIntegrationPgSuite) TestGetByEmailUserSuccess(t provider.T) {
	t.Title("GetByEmailUser: Success")
	t.Tags("UserIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			User, err := s.userRepo.GetByEmail(txCtx, "pavel@ppo.ru")

			// Assert
			sCtx.Assert().Nil(err)
			sCtx.Assert().Equal(data_builders.NewUserBuilder().
				WithUserID(1).
				WithEmail("pavel@ppo.ru").
				WithType(models.ManagerUser).
				WithName("pavel-manager").WithPassword("123123").Build(), User)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *UserIntegrationPgSuite) TestGetByEmailUserFailure(t provider.T) {
	t.Title("GetByEmailUser: Failure")
	t.Tags("UserIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			User, err := s.userRepo.GetByEmail(txCtx, "notexistssss@ppo.ru")

			// Assert
			sCtx.Assert().Nil(User)
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *UserIntegrationPgSuite) TestUpdateUserSuccess(t provider.T) {
	t.Title("UpdateUser: Success")
	t.Tags("UserIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			newUser := data_builders.NewUserBuilder().
				WithUserID(1).
				WithEmail("pavel@popo.ru").
				WithType(models.ManagerUser).
				WithName("pavel-manager").WithPassword("123123").Build()

			// Act
			err := s.userRepo.Update(txCtx, newUser)

			// Assert
			sCtx.Assert().Nil(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}
