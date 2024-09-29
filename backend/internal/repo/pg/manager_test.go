package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	transactor_impl "cookdroogers/internal/transactor/trm"
	"cookdroogers/models"
	"cookdroogers/models/data_builders"
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
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.ManagerRepo
	ctx  context.Context
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

func (s *ManagerPgRepoSuite) TestManagerPgRepo_GetSuccess(t provider.T) {
	t.Title("Get: Success")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание для первого запроса
		q1 := "SELECT manager_id, user_id FROM managers WHERE manager_id=$1"
		s.mock.ExpectQuery(q1).
			WithArgs(uint64(8)).
			WillReturnRows(sqlmock.NewRows([]string{"manager_id", "user_id"}).
				AddRow(8, 88))

		// Ожидание для второго запроса
		q2 := "SELECT artist_id FROM artists WHERE manager_id=$1"
		s.mock.ExpectQuery(q2).
			WithArgs(uint64(8)).
			WillReturnRows(sqlmock.NewRows([]string{"manager_id"}))

		expectedManager := data_builders.NewManagerBuilder().WithManagerID(8).WithUserID(88).WithArtists([]uint64{}).Build()

		Manager, err := s.repo.Get(s.ctx, uint64(8))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedManager, Manager)
	})
}

func (s *ManagerPgRepoSuite) TestRequestPgRepo_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание для первого запроса
		q1 := "SELECT manager_id, user_id FROM managers WHERE manager_id=$1"
		s.mock.ExpectQuery(q1).WillReturnError(sql.ErrConnDone)

		// Ожидание для второго запроса
		q2 := "SELECT artist_id FROM artists WHERE manager_id=$1"
		s.mock.ExpectQuery(q2).WillReturnError(sql.ErrConnDone)

		manager, err := s.repo.Get(s.ctx, uint64(8))

		sCtx.Assert().Nil(manager)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ManagerPgRepoSuite) TestManagerPgRepo_GetByUserIDSuccess(t provider.T) {
	t.Title("GetByUserID: Success")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание для первого запроса
		q1 := "SELECT manager_id, user_id FROM managers WHERE user_id=$1"
		s.mock.ExpectQuery(q1).
			WithArgs(uint64(8)).
			WillReturnRows(sqlmock.NewRows([]string{"manager_id", "user_id"}).
				AddRow(8, 88))

		// Ожидание для второго запроса
		q2 := "SELECT artist_id FROM artists WHERE manager_id=$1"
		s.mock.ExpectQuery(q2).
			WithArgs(uint64(8)).
			WillReturnRows(sqlmock.NewRows([]string{"manager_id"}))

		expectedManager := data_builders.NewManagerBuilder().WithManagerID(8).WithUserID(88).WithArtists([]uint64{}).Build()

		Manager, err := s.repo.GetByUserID(s.ctx, uint64(8))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedManager, Manager)
	})
}

func (s *ManagerPgRepoSuite) TestRequestPgRepo_GetByUserIDFailure(t provider.T) {
	t.Title("GetByUserID: Failure")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание для первого запроса
		q1 := "SELECT manager_id, user_id FROM managers WHERE user_id=$1"
		s.mock.ExpectQuery(q1).WillReturnError(sql.ErrConnDone)

		// Ожидание для второго запроса
		q2 := "SELECT artist_id FROM artists WHERE manager_id=$1"
		s.mock.ExpectQuery(q2).WillReturnError(sql.ErrConnDone)

		manager, err := s.repo.GetByUserID(s.ctx, uint64(8))

		sCtx.Assert().Nil(manager)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ManagerPgRepoSuite) TestManagerPgRepo_GetForAdminSuccess(t provider.T) {
	t.Title("GetForAdmin: Success")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание для первого запроса
		q1 := "SELECT manager_id, user_id FROM managers"
		s.mock.ExpectQuery(q1).
			WillReturnRows(sqlmock.NewRows([]string{"manager_id", "user_id"}).
				AddRow(8, 88))

		// Ожидание для второго запроса
		q2 := "SELECT artist_id FROM artists WHERE manager_id=$1"
		s.mock.ExpectQuery(q2).
			WithArgs(uint64(8)).
			WillReturnRows(sqlmock.NewRows([]string{"manager_id"}))

		expectedManager := data_builders.NewManagerBuilder().WithManagerID(8).WithUserID(88).WithArtists([]uint64{}).Build()
		expectedManagers := append(make([]models.Manager, 0), *expectedManager)

		manager, err := s.repo.GetForAdmin(s.ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedManagers, manager)
	})
}

func (s *ManagerPgRepoSuite) TestManagerPgRepo_GetForAdminFailure(t provider.T) {
	t.Title("GetForAdmin: Failure")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание для первого запроса
		q1 := "SELECT manager_id, user_id FROM managers"
		s.mock.ExpectQuery(q1).WillReturnError(sql.ErrConnDone)

		// Ожидание для второго запроса
		q2 := "SELECT artist_id FROM artists WHERE manager_id=$1"
		s.mock.ExpectQuery(q2).WillReturnError(sql.ErrConnDone)

		manager, err := s.repo.GetForAdmin(s.ctx)

		sCtx.Assert().Nil(manager)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ManagerPgRepoSuite) TestManagerPgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание начала транзакции
		s.mock.ExpectBegin()

		// Ожидание первого запроса (INSERT INTO managers)
		q1 := "INSERT INTO managers(user_id) VALUES($1) RETURNING manager_id"
		s.mock.ExpectQuery(q1).
			WithArgs(uint64(88)). // Обратите внимание на значение UserID
			WillReturnRows(sqlmock.NewRows([]string{"manager_id"}).AddRow(8))

		// Ожидание второго запроса (UPDATE artists)
		q2 := "UPDATE artists SET manager_id=$1 WHERE artist_id=$2"
		s.mock.ExpectExec(q2).
			WithArgs(uint64(8), uint64(4)).
			WillReturnResult(sqlmock.NewResult(1, 1)) // Успешное выполнение UPDATE

		// Ожидание коммита транзакции
		s.mock.ExpectCommit()

		// Ожидаемый результат
		expectedManager := data_builders.NewManagerBuilder().WithManagerID(8).WithUserID(88).WithArtists([]uint64{4}).Build()
		manager := data_builders.NewManagerBuilder().WithUserID(88).WithArtists([]uint64{4}).Build() // Добавление артиста для теста

		err := s.repo.Create(s.ctx, manager)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedManager, manager)
	})
}

func (s *ManagerPgRepoSuite) TestManagerPgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		s.mock.ExpectBegin()

		// Ожидание первого запроса (INSERT INTO managers)
		q1 := "INSERT INTO managers(user_id) VALUES($1) RETURNING manager_id"
		s.mock.ExpectQuery(q1).
			WillReturnError(sql.ErrConnDone)

		// Ожидание второго запроса (UPDATE artists)
		q2 := "UPDATE artists SET manager_id=$1 WHERE artist_id=$2"
		s.mock.ExpectExec(q2).
			WillReturnError(sql.ErrConnDone)

		s.mock.ExpectCommit()

		expectedManager := data_builders.NewManagerBuilder().WithManagerID(8).WithUserID(88).WithArtists([]uint64{4}).Build()
		manager := data_builders.NewManagerBuilder().WithUserID(88).WithArtists([]uint64{4}).Build() // Добавление артиста для теста

		err := s.repo.Create(s.ctx, manager)

		sCtx.Assert().ErrorIs(err, PgDbErr)
		sCtx.Assert().Equal(expectedManager, manager)
	})
}

func (s *ManagerPgRepoSuite) TestManagerPgRepo_GetRandManagerIDSuccess(t provider.T) {
	t.Title("GetRandManagerID: Success")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание для первого запроса
		q := "SELECT manager_id FROM managers ORDER BY random() LIMIT 1"
		s.mock.ExpectQuery(q).
			WillReturnRows(sqlmock.NewRows([]string{"manager_id"}).
				AddRow(8))

		var expectedManagerID uint64 = 8
		managerID, err := s.repo.GetRandManagerID(s.ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedManagerID, managerID)
	})
}

func (s *ManagerPgRepoSuite) TestRequestPgRepo_GetRandManagerIDFailure(t provider.T) {
	t.Title("GetRandManagerID: Failure")
	t.Tags("ManagerPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание для первого запроса
		q := "SELECT manager_id FROM managers ORDER BY random() LIMIT 1"
		s.mock.ExpectQuery(q).WillReturnError(sql.ErrConnDone)

		managerID, err := s.repo.GetRandManagerID(s.ctx)

		sCtx.Assert().Zero(managerID)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func TestManagerPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ManagerPgRepoSuite))
}