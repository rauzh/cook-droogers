package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models/data_builders"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type ArtistPgRepoSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.ArtistRepo
	ctx  context.Context
}

func (s *ArtistPgRepoSuite) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewArtistPgRepo(s.db)
	s.ctx = context.Background()
}

func (s *ArtistPgRepoSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *ArtistPgRepoSuite) Test_ArtistPgRepoGetSuccess(t provider.T) {
	t.Title("Get: Success")
	t.Tags("ArtistPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "SELECT artist_id, nickname, contract_due, activity, user_id, manager_id FROM artists WHERE artist_id=$1"
		s.mock.ExpectQuery(q).
			WithArgs(uint64(1)).
			WillReturnRows(sqlmock.NewRows([]string{"artist_id", "nickname", "contract_due", "activity", "user_id", "manager_id"}).
				AddRow(1, "uzi", cdtime.GetEndOfContract(), true, uint64(7), uint64(9)))

		expectedArtist := data_builders.NewArtistBuilder().Build()

		artist, err := s.repo.Get(s.ctx, uint64(1))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedArtist, artist)
	})
}

func (s *ArtistPgRepoSuite) TestRequestPgRepo_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("ArtistPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "SELECT artist_id, nickname, contract_due, activity, user_id, manager_id FROM artists WHERE artist_id=$1"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		artist, err := s.repo.Get(s.ctx, uint64(1))

		sCtx.Assert().Nil(artist)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ArtistPgRepoSuite) TestArtistPgRepo_GetByIDSuccess(t provider.T) {
	t.Title("GetByID: Success")
	t.Tags("ArtistPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "SELECT artist_id, nickname, contract_due, activity, user_id, manager_id FROM artists WHERE user_id=$1"
		s.mock.ExpectQuery(q).
			WithArgs(uint64(7)).
			WillReturnRows(sqlmock.NewRows([]string{"artist_id", "nickname", "contract_due", "activity", "user_id", "manager_id"}).
				AddRow(1, "uzi", cdtime.GetEndOfContract(), true, uint64(7), uint64(9)))

		expectedArtist := data_builders.NewArtistBuilder().WithUserID(7).Build()

		artist, err := s.repo.GetByUserID(s.ctx, uint64(7))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedArtist, artist)
	})
}

func (s *ArtistPgRepoSuite) TestRequestPgRepo_GetByIDFailure(t provider.T) {
	t.Title("GetByID: Failure")
	t.Tags("ArtistPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "SELECT artist_id, nickname, contract_due, activity, user_id, manager_id FROM artists WHERE user_id=$1"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		artist, err := s.repo.GetByUserID(s.ctx, uint64(7))

		sCtx.Assert().Nil(artist)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ArtistPgRepoSuite) TestArtistPgRepo_UpdateSuccess(t provider.T) {
	t.Title("Update: Success")
	t.Tags("ArtistPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "UPDATE artists SET user_id=$1, nickname=$2, contract_due=$3, activity=$4, manager_id=$5 WHERE artist_id=$6 RETURNING *"
		s.mock.ExpectQuery(q).
			WithArgs(uint64(7), "carti", cdtime.GetEndOfContract(), true, uint64(9), uint64(1)).
			WillReturnRows(sqlmock.NewRows([]string{"artist_id", "nickname", "contract_due", "activity", "user_id", "manager_id"}).
				AddRow(1, "carti", cdtime.GetEndOfContract(), true, uint64(7), uint64(9)))

		expectedArtist := data_builders.NewArtistBuilder().WithNickname("carti").Build()
		artist := data_builders.NewArtistBuilder().WithNickname("carti").Build()

		err := s.repo.Update(s.ctx, artist)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedArtist, artist)
	})
}

func (s *ArtistPgRepoSuite) TestRequestPgRepo_UpdateFailure(t provider.T) {
	t.Title("Update: Failure")
	t.Tags("ArtistPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "UPDATE artists SET user_id=$1, nickname=$2, contract_due=$3, activity=$4, manager_id=$5 WHERE artist_id=$6 RETURNING *"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		expectedArtist := data_builders.NewArtistBuilder().WithNickname("carti").Build()
		artist := data_builders.NewArtistBuilder().WithNickname("carti").Build()

		err := s.repo.Update(s.ctx, artist)

		sCtx.Assert().ErrorIs(err, PgDbErr)
		sCtx.Assert().Equal(expectedArtist, artist)
	})
}

func (s *ArtistPgRepoSuite) TestArtistPgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("ArtistPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "INSERT INTO artists(user_id, nickname, contract_due, activity, manager_id)" +
			"VALUES($1, $2, $3, $4, $5) RETURNING artist_id"
		s.mock.ExpectQuery(q).
			WithArgs(uint64(7), "uzi", cdtime.GetEndOfContract(), true, uint64(9)).
			WillReturnRows(sqlmock.NewRows([]string{"artist_id"}).AddRow(27))

		expectedArtist := data_builders.NewArtistBuilder().WithID(uint64(27)).Build()

		artist := data_builders.NewArtistBuilder().WithID(0).Build()

		err := s.repo.Create(s.ctx, artist)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedArtist, artist)
	})
}

func (s *ArtistPgRepoSuite) TestRequestPgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("ArtistPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "INSERT INTO artists(user_id, nickname, contract_due, activity, manager_id)" +
			"VALUES($1, $2, $3, $4, $5) RETURNING artist_id"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		expectedArtist := data_builders.NewArtistBuilder().WithID(0).Build()
		artist := data_builders.NewArtistBuilder().WithID(0).Build()

		err := s.repo.Create(s.ctx, artist)

		sCtx.Assert().ErrorIs(err, PgDbErr)
		sCtx.Assert().Equal(expectedArtist, artist)
	})
}

func TestArtistPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ArtistPgRepoSuite))
}
