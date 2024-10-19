//go:build unit
// +build unit

package pg

import (
	"context"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/base/repo"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

// TEST_HW: default db access test with before each & after each
type RequestPgRepoSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.RequestRepo
	ctx  context.Context
}

func (s *RequestPgRepoSuite) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewRequestPgRepo(s.db)
	s.ctx = context.Background()
}

func (s *RequestPgRepoSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *RequestPgRepoSuite) TestRequestPgRepo_GetAllByManagerIDSuccess(t provider.T) {
	t.Title("GetAllByManagerID: Success")
	t.Tags("RequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := `SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE manager_id=$1`
		s.mock.ExpectQuery(q).
			WithArgs(uint64(9)).
			WillReturnRows(sqlmock.NewRows([]string{"request_id", "status", "type", "creation_date", "manager_id", "user_id"}).
				AddRow(1, base.OnApprovalRequest, "Sign", cdtime.GetToday(), uint64(9), uint64(12)))

		expectedReqs := []base.Request{*base.GetBaseRequestObject()}

		reqs, err := s.repo.GetAllByManagerID(s.ctx, uint64(9))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedReqs, reqs)
	})
}

func (s *RequestPgRepoSuite) TestRequestPgRepo_GetAllByManagerIDFailure(t provider.T) {
	t.Title("GetAllByManagerID: Failure")
	t.Tags("RequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := `SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE manager_id=$1`
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		var expectedReqs []base.Request

		reqs, err := s.repo.GetAllByManagerID(s.ctx, uint64(9))

		sCtx.Assert().Equal(expectedReqs, reqs)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *RequestPgRepoSuite) TestRequestPgRepo_GetAllByUserIDSuccess(t provider.T) {
	t.Title("GetAllByUserID: Success")
	t.Tags("RequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := `SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE user_id=$1`
		s.mock.ExpectQuery(q).
			WithArgs(uint64(12)).
			WillReturnRows(sqlmock.NewRows([]string{"request_id", "status", "type", "creation_date", "manager_id", "user_id"}).
				AddRow(1, base.OnApprovalRequest, "Sign", cdtime.GetToday(), uint64(9), uint64(12)))

		expectedReqs := []base.Request{*base.GetBaseRequestObject()}

		reqs, err := s.repo.GetAllByUserID(s.ctx, uint64(12))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedReqs, reqs)
	})
}

func (s *RequestPgRepoSuite) TestRequestPgRepo_GetAllByUserIDFailure(t provider.T) {
	t.Title("GetAllByUserID: Failure")
	t.Tags("RequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := `SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE user_id=$1`
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		var expectedReqs []base.Request

		reqs, err := s.repo.GetAllByUserID(s.ctx, uint64(12))

		sCtx.Assert().Equal(expectedReqs, reqs)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func (s *RequestPgRepoSuite) TestRequestPgRepo_GetByIDSuccess(t provider.T) {
	t.Title("GetByID: Success")
	t.Tags("RequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		q := `SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE request_id=$1`
		s.mock.ExpectQuery(q).
			WithArgs(uint64(1)).
			WillReturnRows(sqlmock.NewRows([]string{"request_id", "status", "type", "creation_date", "manager_id", "user_id"}).
				AddRow(1, base.OnApprovalRequest, "Sign", cdtime.GetToday(), uint64(9), uint64(12)))

		expectedReq := base.GetBaseRequestObject()

		req, err := s.repo.GetByID(s.ctx, uint64(1))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(expectedReq, req)
	})
}

func (s *RequestPgRepoSuite) TestRequestPgRepo_GetByIDFailure(t provider.T) {
	t.Title("GetByID: Failure")
	t.Tags("RequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		q := `SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE request_id=$1`
		s.mock.ExpectQuery(q).
			WillReturnError(sql.ErrConnDone)

		reqs, err := s.repo.GetByID(s.ctx, uint64(12))

		sCtx.Assert().Nil(reqs)
		sCtx.Assert().ErrorIs(err, PgDbErr)
	})
}

func TestRequestPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(RequestPgRepoSuite))
}
