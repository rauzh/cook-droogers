//go:build unit
// +build unit

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

type ReleasePgRepoSuit struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.ReleaseRepo
	ctx  context.Context
}

func (s *ReleasePgRepoSuit) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewReleasePgRepo(s.db, transactor_impl.NewATtrm(manager.Must(trmsqlx.NewDefaultFactory(sqlx.NewDb(s.db, "pgx")))))
	s.ctx = context.Background()
}

func (s *ReleasePgRepoSuit) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		s.mock.ExpectBegin()

		q1 := "INSERT INTO releases (title, status, creation_date, artist_id) VALUES ($1, $2, $3, $4) RETURNING release_id"
		s.mock.ExpectQuery(q1).
			WithArgs("title", models.PublishedRelease, date, 7).
			WillReturnRows(sqlmock.NewRows([]string{"release_id"}).
				AddRow(888))

		q2 := "UPDATE tracks SET release_id=$1 WHERE track_id=$2"
		s.mock.ExpectExec(q2).
			WithArgs(888, 1111).
			WillReturnResult(sqlmock.NewResult(1, 1))

		s.mock.ExpectCommit()

		release := data_builders.NewReleaseBuilder().WithReleaseID(888).WithArtistID(7).WithTitle("title").
			WithStatus(models.PublishedRelease).WithDate(date).Build()

		err := s.repo.Create(s.ctx, release)

		sCtx.Assert().NoError(err)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		s.mock.ExpectBegin()

		q1 := "INSERT INTO releases (title, status, creation_date, artist_id) VALUES ($1, $2, $3, $4) RETURNING release_id"
		s.mock.ExpectQuery(q1).
			WillReturnError(sql.ErrConnDone)

		q2 := "UPDATE tracks SET release_id=$1 WHERE track_id=$2"
		s.mock.ExpectExec(q2).
			WillReturnError(sql.ErrConnDone)

		s.mock.ExpectCommit()

		release := data_builders.NewReleaseBuilder().WithReleaseID(888).WithArtistID(7).WithTitle("title").
			WithStatus(models.PublishedRelease).WithDate(date).Build()

		err := s.repo.Create(s.ctx, release)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_GetSuccess(t provider.T) {
	t.Title("Get: Success")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q1 := "SELECT title, status, creation_date, artist_id FROM releases WHERE release_id=$1"
		s.mock.ExpectQuery(q1).
			WithArgs(888).
			WillReturnRows(sqlmock.NewRows([]string{"title", "status", "creation_date", "artist_id"}).
				AddRow("title", models.PublishedRelease, date, 7))

		q2 := "SELECT track_id FROM tracks WHERE release_id=$1"
		s.mock.ExpectQuery(q2).
			WithArgs(888).
			WillReturnRows(sqlmock.NewRows([]string{"track_id"}).
				AddRow(1111))

		expectedRelease := data_builders.NewReleaseBuilder().WithReleaseID(888).WithArtistID(7).WithTitle("title").
			WithStatus(models.PublishedRelease).WithDate(date).Build()

		release, err := s.repo.Get(s.ctx, 888)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedRelease, release)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q1 := "SELECT title, status, creation_date, artist_id FROM releases WHERE release_id=$1"
		s.mock.ExpectQuery(q1).
			WillReturnError(sql.ErrConnDone)

		q2 := "SELECT track_id FROM tracks WHERE release_id=$1"
		s.mock.ExpectQuery(q2).
			WillReturnError(sql.ErrConnDone)

		release, err := s.repo.Get(s.ctx, 888)

		sCtx.Assert().Nil(release)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_GetAllByArtistSuccess(t provider.T) {
	t.Title("GetAllByArtist: Success")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q1 := "SELECT title, status, creation_date, release_id FROM releases WHERE artist_id=$1"
		s.mock.ExpectQuery(q1).
			WithArgs(7).
			WillReturnRows(sqlmock.NewRows([]string{"title", "status", "creation_date", "release_id"}).
				AddRow("title", models.PublishedRelease, date, 888))

		q2 := "SELECT track_id FROM tracks WHERE release_id=$1"
		s.mock.ExpectQuery(q2).
			WithArgs(888).
			WillReturnRows(sqlmock.NewRows([]string{"track_id"}).
				AddRow(1111))

		expectedRelease := data_builders.NewReleaseBuilder().WithReleaseID(888).WithArtistID(7).WithTitle("title").
			WithStatus(models.PublishedRelease).WithDate(date).Build()
		expectedReleases := append(make([]models.Release, 0), *expectedRelease)

		releases, err := s.repo.GetAllByArtist(s.ctx, 7)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedReleases, releases)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_GetAllByArtistFailure(t provider.T) {
	t.Title("GetAllByArtist: Failure")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q1 := "SELECT title, status, creation_date, release_id FROM releases WHERE artist_id=$1"
		s.mock.ExpectQuery(q1).
			WillReturnError(sql.ErrConnDone)

		q2 := "SELECT track_id FROM tracks WHERE release_id=$1"
		s.mock.ExpectQuery(q2).
			WillReturnError(sql.ErrConnDone)

		releases, err := s.repo.GetAllByArtist(s.ctx, 7)

		sCtx.Assert().Nil(releases)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_GetAllTracksSuccess(t provider.T) {
	t.Title("GetAllTracks: Success")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q1 := "SELECT track_id, title, genre, duration, type FROM tracks WHERE track_id=$1"
		s.mock.ExpectQuery(q1).
			WithArgs(1111).
			WillReturnRows(sqlmock.NewRows([]string{"track_id", "title", "genre", "duration", "type"}).
				AddRow("1111", "title", "trap", 60, "intro"))

		q2 := "SELECT artist_id FROM track_artist WHERE track_id=$1"
		s.mock.ExpectQuery(q2).
			WithArgs(1111).
			WillReturnRows(sqlmock.NewRows([]string{"artist_id"}).
				AddRow(7))

		release := data_builders.NewReleaseBuilder().Build()
		expectedTrack := data_builders.NewTrackBuilder().WithArtists([]uint64{7}).Build()
		expectedTracks := append(make([]models.Track, 0), *expectedTrack)

		tracks, err := s.repo.GetAllTracks(s.ctx, release)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedTracks, tracks)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_GetAllTracksFailure(t provider.T) {
	t.Title("GetAllTracks: Failure")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q1 := "SELECT track_id, title, genre, duration, type FROM tracks WHERE track_id=$1"
		s.mock.ExpectQuery(q1).
			WillReturnError(sql.ErrConnDone)

		q2 := "SELECT artist_id FROM track_artist WHERE track_id=$1"
		s.mock.ExpectQuery(q2).
			WillReturnError(sql.ErrConnDone)

		release := data_builders.NewReleaseBuilder().Build()
		tracks, err := s.repo.GetAllTracks(s.ctx, release)

		sCtx.Assert().Nil(tracks)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_UpdateSuccess(t provider.T) {
	t.Title("Update: Success")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "UPDATE releases SET title=$1, status=$2, creation_date=$3, artist_id=$4 WHERE release_id=$5"
		s.mock.ExpectExec(q).
			WithArgs("title", models.PublishedRelease, date, 7, 888).
			WillReturnResult(sqlmock.NewResult(1, 1))

		release := data_builders.NewReleaseBuilder().Build()

		err := s.repo.Update(s.ctx, release)

		sCtx.Assert().NoError(err)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_UpdateFailure(t provider.T) {
	t.Title("Update: Failure")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "UPDATE releases SET title=$1, status=$2, creation_date=$3, artist_id=$4 WHERE release_id=$5"
		s.mock.ExpectExec(q).
			WillReturnError(sql.ErrConnDone)

		release := data_builders.NewReleaseBuilder().Build()

		err := s.repo.Update(s.ctx, release)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_UpdateStatusSuccess(t provider.T) {
	t.Title("UpdateStatus: Success")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		q := "UPDATE releases SET status=$1 WHERE release_id=$2"
		s.mock.ExpectExec(q).
			WithArgs(models.PublishedRelease, 888).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := s.repo.UpdateStatus(s.ctx, 888, models.PublishedRelease)

		sCtx.Assert().NoError(err)
	})
}

func (s *ReleasePgRepoSuit) TestReleasePgRepo_UpdateStatusFailure(t provider.T) {
	t.Title("UpdateStatus: Failure")
	t.Tags("ReleasePgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {

		q := "UPDATE releases SET status=$1 WHERE release_id=$2"
		s.mock.ExpectExec(q).
			WillReturnError(sql.ErrConnDone)

		err := s.repo.UpdateStatus(s.ctx, 888, models.PublishedRelease)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func TestReleasePgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ReleasePgRepoSuit))
}
