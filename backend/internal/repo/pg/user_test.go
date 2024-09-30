package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"cookdroogers/models/data_builders"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type UserPgRepoSuit struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.UserRepo
	ctx  context.Context
}

func (s *UserPgRepoSuit) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewUserPgRepo(s.db)
	s.ctx = context.Background()
}

func (s *UserPgRepoSuit) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *UserPgRepoSuit) TestUserPgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("UserPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "INSERT INTO users (name, email, type, password) VALUES ($1, $2, $3, $4) RETURNING user_id"
		s.mock.ExpectQuery(q).
			WithArgs("uzi", "uzi@gmail.com", models.NonMemberUser, "password").
			WillReturnRows(sqlmock.NewRows([]string{"user_id"}).
				AddRow(7))

		user := data_builders.NewUserBuilder().Build()

		err := s.repo.Create(s.ctx, user)

		sCtx.Assert().NoError(err)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("UserPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "INSERT INTO users (name, email, type, password) VALUES ($1, $2, $3, $4) RETURNING user_id"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		user := data_builders.NewUserBuilder().Build()

		err := s.repo.Create(s.ctx, user)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_GetByEmailSuccess(t provider.T) {
	t.Title("GetByEmail: Success")
	t.Tags("UserPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "SELECT user_id, name, email, password, type FROM users WHERE email=$1"
		s.mock.ExpectQuery(q).
			WithArgs("uzi@gmail.com").
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "name", "email", "password", "type"}).
				AddRow(7, "uzi", "uzi@gmail.com", "password", models.NonMemberUser))

		expectedUser := data_builders.NewUserBuilder().Build()

		user, err := s.repo.GetByEmail(s.ctx, "uzi@gmail.com")

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedUser, user)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_GetByEmailFailure(t provider.T) {
	t.Title("GetByEmail: Failure")
	t.Tags("UserPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "SELECT user_id, name, email, password, type FROM users WHERE email=$1"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		user, err := s.repo.GetByEmail(s.ctx, "uzi@gmail.com")

		sCtx.Assert().Nil(user)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_GetSuccess(t provider.T) {
	t.Title("Get: Success")
	t.Tags("UserPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "SELECT user_id, name, email, password, type FROM users WHERE user_id=$1"
		s.mock.ExpectQuery(q).
			WithArgs(7).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "name", "email", "password", "type"}).
				AddRow(7, "uzi", "uzi@gmail.com", "password", models.NonMemberUser))

		expectedUser := data_builders.NewUserBuilder().Build()

		user, err := s.repo.Get(s.ctx, 7)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedUser, user)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("UserPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "SELECT user_id, name, email, password, type FROM users WHERE user_id=$1"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		user, err := s.repo.Get(s.ctx, 7)

		sCtx.Assert().Nil(user)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_UpdateSuccess(t provider.T) {
	t.Title("Update: Success")
	t.Tags("UserPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "UPDATE users SET name=$1, email=$2, type=$3, password=$4 WHERE user_id=$5"
		s.mock.ExpectExec(q).
			WithArgs("uzi", "uzi@gmail.com", models.NonMemberUser, "password", 7).
			WillReturnResult(sqlmock.NewResult(1, 1))

		user := data_builders.NewUserBuilder().Build()
		err := s.repo.Update(s.ctx, user)

		sCtx.Assert().NoError(err)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_UpdateFailure(t provider.T) {
	t.Title("Update: Failure")
	t.Tags("UserPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "UPDATE users SET name=$1, email=$2, type=$3, password=$4 WHERE user_id=$5"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		user := data_builders.NewUserBuilder().Build()
		err := s.repo.Update(s.ctx, user)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_UpdateTypeSuccess(t provider.T) {
	t.Title("UpdateType: Success")
	t.Tags("UserPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "UPDATE users SET type=$1 WHERE user_id=$2"
		s.mock.ExpectExec(q).
			WithArgs(models.ArtistUser, 7).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := s.repo.UpdateType(s.ctx, 7, models.ArtistUser)

		sCtx.Assert().NoError(err)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_UpdateTypeFailure(t provider.T) {
	t.Title("UpdateType: Failure")
	t.Tags("UserPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "UPDATE users SET type=$1 WHERE user_id=$2"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		err := s.repo.UpdateType(s.ctx, 7, models.ArtistUser)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_GetForAdminSuccess(t provider.T) {
	t.Title("GetForAdmin: Success")
	t.Tags("UserPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "SELECT user_id, name, email, password, type FROM users"
		s.mock.ExpectQuery(q).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "name", "email", "password", "type"}).
				AddRow(7, "uzi", "uzi@gmail.com", "password", models.NonMemberUser))

		expectedUser := data_builders.NewUserBuilder().Build()
		expectedUsers := append(make([]models.User, 0), *expectedUser)

		users, err := s.repo.GetForAdmin(s.ctx)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedUsers, users)
	})
}

func (s *UserPgRepoSuit) TestUserPgRepo_GetForAdminFailure(t provider.T) {
	t.Title("GetForAdmin: Failure")
	t.Tags("UserPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "SELECT user_id, name, email, password, type FROM users"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		_, err := s.repo.GetForAdmin(s.ctx)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func TestUserPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(UserPgRepoSuit))
}
