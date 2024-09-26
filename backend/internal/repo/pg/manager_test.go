package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/transactor"
	transactor_impl "cookdroogers/internal/transactor/trm"
	"cookdroogers/models/data_builders"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type ManagerPgRepoSuite struct {
	suite.Suite
	db         *sql.DB
	mock       sqlmock.Sqlmock
	repo       repo.ManagerRepo
	transactor transactor.Transactor
	ctx        context.Context
}

func (s *ManagerPgRepoSuite) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewManagerPgRepo(s.db, transactor_impl.NewATtrm(manager.Must(trmsqlx.NewDefaultFactory(sqlx.NewDb(s.db, "pgx")))))
	s.ctx = context.Background()
}

func (s *ManagerPgRepoSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *ManagerPgRepoSuite) TestManagerPgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "INSERT INTO managers(user_id) VALUES($1) RETURNING manager_id"
		s.mock.ExpectQuery(q).
			WithArgs(uint64(7), "uzi", cdtime.GetEndOfContract(), true, uint64(9)).
			WillReturnRows(sqlmock.NewRows([]string{"Manager_id"}).AddRow(1))

		expectedManager := data_builders.NewManagerBuilder().Build()
		Manager := data_builders.NewManagerBuilder().WithID(0).Build()

		err := s.repo.Create(s.ctx, Manager)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedManager, Manager)
	})
}

func (s *ManagerPgRepoSuite) TestRequestPgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "INSERT INTO Managers(user_id, nickname, contract_due, activity, manager_id)" +
			"VALUES($1, $2, $3, $4, $5) RETURNING Manager_id"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		expectedManager := data_builders.NewManagerBuilder().WithID(0).Build()
		Manager := data_builders.NewManagerBuilder().WithID(0).Build()

		err := s.repo.Create(s.ctx, Manager)

		sCtx.Assert().ErrorIs(err, PgDbErr)
		sCtx.Assert().Equal(expectedManager, Manager)
	})
}

func TestManagerPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ManagerPgRepoSuite))
}
