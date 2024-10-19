//go:build unit
// +build unit

package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"cookdroogers/models/data_builders"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

var date = cdtime.GetToday().AddDate(-1, 0, 0)

type PublicationPgRepoSuit struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.PublicationRepo
	ctx  context.Context
}

func (s *PublicationPgRepoSuit) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewPublicationPgRepo(s.db)
	s.ctx = context.Background()
}

func (s *PublicationPgRepoSuit) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "INSERT INTO publications(manager_id, release_id, creation_date) VALUES($1, $2, $3) RETURNING publication_id"
		s.mock.ExpectQuery(q).
			WithArgs(uint64(8), uint64(888), date).
			WillReturnRows(sqlmock.NewRows([]string{"publication_id"}).
				AddRow(88))

		publication := data_builders.NewPublicationBuilder().WithManagerID(8).WithReleaseID(888).WithPublicationID(88).Build()

		err := s.repo.Create(s.ctx, publication)

		sCtx.Assert().NoError(err)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "INSERT INTO publications(manager_id, release_id, creation_date) VALUES($1, $2, $3) RETURNING publication_id"
		s.mock.ExpectQuery(q).WillReturnError(sql.ErrConnDone)

		publication := data_builders.NewPublicationBuilder().WithManagerID(8).WithReleaseID(888).WithPublicationID(88).Build()

		err := s.repo.Create(s.ctx, publication)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_GetSuccess(t provider.T) {
	t.Title("Get: Success")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE publication_id=$1"
		s.mock.ExpectQuery(q).
			WithArgs(uint64(88)).
			WillReturnRows(sqlmock.NewRows([]string{"publication_id", "creation_date", "manager_id", "release_id"}).
				AddRow(88, date, 8, 888))

		expectedPublication := data_builders.NewPublicationBuilder().WithManagerID(8).WithReleaseID(888).
			WithPublicationID(88).WithDate(date).Build()

		publication, err := s.repo.Get(s.ctx, uint64(88))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedPublication, publication)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE publication_id=$1"
		s.mock.ExpectQuery(q).WillReturnError(sql.ErrConnDone)

		publication, err := s.repo.Get(s.ctx, uint64(88))

		sCtx.Assert().Nil(publication)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_GetAllByDateSuccess(t provider.T) {
	t.Title("GetAllByDate: Success")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE creation_date=$1"
		s.mock.ExpectQuery(q).
			WithArgs(date).
			WillReturnRows(sqlmock.NewRows([]string{"publication_id", "creation_date", "manager_id", "release_id"}).
				AddRow(88, date, 8, 888))

		expectedPublications := make([]models.Publication, 0)
		expectedPublication := data_builders.NewPublicationBuilder().WithManagerID(8).WithReleaseID(888).
			WithPublicationID(88).WithDate(date).Build()
		expectedPublications = append(expectedPublications, *expectedPublication)

		publications, err := s.repo.GetAllByDate(s.ctx, date)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedPublications, publications)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_GetAllByDateFailure(t provider.T) {
	t.Title("GetAllByDate: Failure")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE creation_date=$1"
		s.mock.ExpectQuery(q).WillReturnError(sql.ErrConnDone)

		publications, err := s.repo.GetAllByDate(s.ctx, date)

		sCtx.Assert().Nil(publications)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_GetAllByManagerSuccess(t provider.T) {
	t.Title("GetAllByManager: Success")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE manager_id=$1"
		s.mock.ExpectQuery(q).
			WithArgs(8).
			WillReturnRows(sqlmock.NewRows([]string{"publication_id", "creation_date", "manager_id", "release_id"}).
				AddRow(88, date, 8, 888))

		expectedPublications := make([]models.Publication, 0)
		expectedPublication := data_builders.NewPublicationBuilder().WithManagerID(8).WithReleaseID(888).
			WithPublicationID(88).WithDate(date).Build()
		expectedPublications = append(expectedPublications, *expectedPublication)

		publications, err := s.repo.GetAllByManager(s.ctx, 8)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedPublications, publications)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_GetAllByManagerFailure(t provider.T) {
	t.Title("GetAllByManager: Failure")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE manager_id=$1"
		s.mock.ExpectQuery(q).WillReturnError(sql.ErrConnDone)

		publications, err := s.repo.GetAllByManager(s.ctx, 8)

		sCtx.Assert().Nil(publications)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_GetAllByArtistSinceDateSuccess(t provider.T) {
	t.Title("GetAllByArtistSinceDate: Success")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "SELECT p.publication_id, p.creation_date, p.manager_id, p.release_id " +
			"FROM publications p JOIN releases r ON p.release_id = r.release_id" +
			"WHERE r.artist_id=$1 AND p.creation_date>=$2;"
		s.mock.ExpectQuery(q).
			WithArgs(7, date).
			WillReturnRows(sqlmock.NewRows([]string{"publication_id", "creation_date", "manager_id", "release_id"}).
				AddRow(88, date, 8, 888))

		expectedPublications := make([]models.Publication, 0)
		expectedPublication := data_builders.NewPublicationBuilder().WithManagerID(8).WithReleaseID(888).
			WithPublicationID(88).WithDate(date).Build()
		expectedPublications = append(expectedPublications, *expectedPublication)

		publications, err := s.repo.GetAllByArtistSinceDate(s.ctx, date, 7)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedPublications, publications)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_GetAllByArtistSinceDateFailure(t provider.T) {
	t.Title("GetAllByArtistSinceDate: Failure")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "SELECT p.publication_id, p.creation_date, p.manager_id, p.release_id " +
			"FROM publications p JOIN releases r ON p.release_id = r.release_id" +
			"WHERE r.artist_id=$1 AND p.creation_date>=$2;"
		s.mock.ExpectQuery(q).WillReturnError(sql.ErrConnDone)

		publications, err := s.repo.GetAllByArtistSinceDate(s.ctx, date, 7)

		sCtx.Assert().Nil(publications)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_UpdateSuccess(t provider.T) {
	t.Title("Update: Success")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := "UPDATE publications SET creation_date=$1, manager_id=$2, release_id=$3 WHERE publication_id=$4 RETURNING *"
		s.mock.ExpectQuery(q).
			WithArgs(date, 8, 888, 88).
			WillReturnRows(sqlmock.NewRows([]string{"creation_date", "manager_id", "release_id", "publication_id"}).
				AddRow(date, 8, 888, 88))

		publication := data_builders.NewPublicationBuilder().WithManagerID(8).WithReleaseID(888).
			WithPublicationID(88).WithDate(date).Build()

		err := s.repo.Update(s.ctx, publication)

		sCtx.Assert().NoError(err)
	})
}

func (s *PublicationPgRepoSuit) TestPublicationPgRepo_UpdateFailure(t provider.T) {
	t.Title("Update: Failure")
	t.Tags("PublicationPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := "UPDATE publications SET creation_date=$1, manager_id=$2, release_id=$3 WHERE publication_id=$4 RETURNING *"
		s.mock.ExpectQuery(q).WillReturnError(sql.ErrConnDone)

		publication := data_builders.NewPublicationBuilder().WithManagerID(8).WithReleaseID(888).WithPublicationID(88).Build()
		err := s.repo.Update(s.ctx, publication)

		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func TestPublicationPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(PublicationPgRepoSuit))
}
