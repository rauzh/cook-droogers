//go:build unit
// +build unit

package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models/data_builders"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type TrackPgRepoSuit struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.TrackRepo
	ctx  context.Context
}

func (s *TrackPgRepoSuit) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewTrackPgRepo(s.db)
	s.ctx = context.Background()
}

func (s *TrackPgRepoSuit) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *TrackPgRepoSuit) TestTrackPgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("TrackPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q1 := "INSERT INTO tracks (title, genre, duration, type) VALUES ($1, $2, $3, $4) RETURNING track_id"
		s.mock.ExpectQuery(q1).
			WithArgs("title", "trap", 60, "intro").
			WillReturnRows(sqlmock.NewRows([]string{"track_id"}).
				AddRow(1111))

		q2 := "INSERT INTO track_artist (artist_id, track_id) VALUES ($1, $2) RETURNING track_artist_id"
		s.mock.ExpectQuery(q2).
			WithArgs(7, 1111).
			WillReturnRows(sqlmock.NewRows([]string{"track_artist_id"}).
				AddRow(71111))

		track := data_builders.NewTrackBuilder().Build()

		_, err := s.repo.Create(s.ctx, track)

		sCtx.Assert().NoError(err)
	})
}

func (s *TrackPgRepoSuit) TestTrackPgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("TrackPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		s.mock.ExpectBegin()

		q1 := "INSERT INTO tracks (title, genre, duration, type) VALUES ($1, $2, $3, $4) RETURNING track_id"
		s.mock.ExpectQuery(q1).
			WillReturnError(sql.ErrConnDone)

		q2 := "INSERT INTO track_artist (artist_id, track_id) VALUES ($1, $2) RETURNING track_artist_id"
		s.mock.ExpectExec(q2).
			WillReturnError(sql.ErrConnDone)

		s.mock.ExpectCommit()

		track := data_builders.NewTrackBuilder().Build()

		_, err := s.repo.Create(s.ctx, track)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *TrackPgRepoSuit) TestTrackPgRepo_GetSuccess(t provider.T) {
	t.Title("Get: Success")
	t.Tags("TrackPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "SELECT track_id, title, genre, duration, type FROM tracks WHERE track_id=$1"
		s.mock.ExpectQuery(q).
			WithArgs(1111).
			WillReturnRows(sqlmock.NewRows([]string{"track_id", "title", "genre", "duration", "type"}).
				AddRow(1111, "title", "trap", 60, "intro"))

		expectedTrack := data_builders.NewTrackBuilder().Build()

		track, err := s.repo.Get(s.ctx, 1111)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedTrack, track)
	})
}

func (s *TrackPgRepoSuit) TestTrackPgRepo_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("TrackPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "SELECT track_id, title, genre, duration, type FROM tracks WHERE track_id=$1"
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		track, err := s.repo.Get(s.ctx, 1111)

		sCtx.Assert().Nil(track)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func TestTrackPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(TrackPgRepoSuit))
}
