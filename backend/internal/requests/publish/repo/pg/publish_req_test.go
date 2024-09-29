package pg

import (
	"context"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/publish"
	"cookdroogers/internal/requests/publish/data_builder"
	"cookdroogers/internal/requests/publish/repo"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type PublishRequestPgRepoSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.PublishRequestRepo
	ctx  context.Context
}

func (s *PublishRequestPgRepoSuite) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewPublishRequestPgRepo(s.db)
	s.ctx = context.Background()
}

func (s *PublishRequestPgRepoSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *PublishRequestPgRepoSuite) TestPublishRequestPgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("PublishRequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание запроса (INSERT INTO requests)
		q := "INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)" +
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING request_id"

		meta := PublishRequestMetaPgDTO{
			ReleaseID:    100,
			Grade:        5,
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Description:  "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		s.mock.ExpectQuery(q).
			WithArgs(base.ProcessingRequest, publish.PubReq,
				cdtime.GetToday(), metaJson, uint64(8), uint64(88)).
			WillReturnRows(sqlmock.NewRows([]string{"request_id"}).AddRow(123))

		req := data_builder.NewPublishRequestBuilder().Build()

		// Выполнение метода Create
		err := s.repo.Create(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(uint64(123), req.RequestID) // Проверяем, что ID был правильно присвоен
	})
}

func (s *PublishRequestPgRepoSuite) TestPublishRequestPgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("PublishRequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание запроса (INSERT INTO requests)
		q := "INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)" +
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING request_id"

		meta := PublishRequestMetaPgDTO{
			ReleaseID:    100,
			Grade:        5,
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Description:  "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		s.mock.ExpectQuery(q).
			WithArgs(base.ProcessingRequest, publish.PubReq,
				cdtime.GetToday(), metaJson, uint64(8), uint64(88)).
			WillReturnError(sql.ErrConnDone)

		req := data_builder.NewPublishRequestBuilder().Build()

		// Выполнение метода Create
		err := s.repo.Create(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
	})
}

func (s *PublishRequestPgRepoSuite) TestPublishRequestPgRepo_UpdateSuccess(t provider.T) {
	t.Title("Update: Success")
	t.Tags("PublishRequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание запроса (UPDATE requests)
		q := "UPDATE requests SET status=$1, type=$2, creation_date=$3, meta=$4, manager_id=$5, user_id=$6 WHERE request_id=$7"

		meta := PublishRequestMetaPgDTO{
			ReleaseID:    100,
			Grade:        5,
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Description:  "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		s.mock.ExpectExec(q).
			WithArgs(base.ProcessingRequest, publish.PubReq,
									cdtime.GetToday(), metaJson, uint64(8), uint64(88), uint64(123)).
			WillReturnResult(sqlmock.NewResult(1, 1)) // Симулируем успешное выполнение запроса

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(123).
			Build()

		// Выполнение метода Update
		err := s.repo.Update(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().NoError(err)
	})
}

func (s *PublishRequestPgRepoSuite) TestPublishRequestPgRepo_UpdateFailure(t provider.T) {
	t.Title("Update: Failure")
	t.Tags("PublishRequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание запроса с ошибкой
		q := "UPDATE requests SET status=$1, type=$2, creation_date=$3, meta=$4, manager_id=$5, user_id=$6 WHERE request_id=$7"

		meta := PublishRequestMetaPgDTO{
			ReleaseID:    100,
			Grade:        5,
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Description:  "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		s.mock.ExpectExec(q).
			WithArgs(base.ProcessingRequest, publish.PubReq,
								cdtime.GetToday(), metaJson, uint64(8), uint64(88), uint64(123)).
			WillReturnError(sql.ErrConnDone) // Симуляция ошибки при выполнении запроса

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(123).
			Build()

		// Выполнение метода Update
		err := s.repo.Update(s.ctx, req)

		// Проверка на наличие ошибки
		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
	})
}

func (s *PublishRequestPgRepoSuite) TestPublishRequestPgRepo_GetSuccess(t provider.T) {
	t.Title("Get: Success")
	t.Tags("PublishRequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание запроса SELECT
		q := "SELECT request_id, status, type, creation_date, meta, manager_id, user_id FROM requests WHERE request_id=$1"

		meta := PublishRequestMetaPgDTO{
			ReleaseID:    100,
			Grade:        5,
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Description:  "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		// Эмулируем успешный ответ от базы данных
		s.mock.ExpectQuery(q).
			WithArgs(uint64(123)).
			WillReturnRows(sqlmock.NewRows([]string{"request_id", "status", "type", "creation_date", "meta", "manager_id", "user_id"}).
				AddRow(uint64(123), base.ProcessingRequest, publish.PubReq, cdtime.GetToday(), metaJson, uint64(8), uint64(88)))

		// Выполнение метода Get
		req, err := s.repo.Get(s.ctx, 123)

		// Проверка результатов
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(req)
		sCtx.Assert().Equal(uint64(123), req.RequestID)
		sCtx.Assert().Equal("Test description", req.Description)
	})
}

func (s *PublishRequestPgRepoSuite) TestPublishRequestPgRepo_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("PublishRequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание запроса SELECT с ошибкой
		q := "SELECT request_id, status, type, creation_date, meta, manager_id, user_id FROM requests WHERE request_id=$1"

		s.mock.ExpectQuery(q).
			WithArgs(uint64(123)).
			WillReturnError(sql.ErrNoRows)

		// Выполнение метода Get
		req, err := s.repo.Get(s.ctx, 123)

		// Проверка результатов
		sCtx.Assert().ErrorIs(err, sql.ErrNoRows)
		sCtx.Assert().Nil(req)
	})
}

func (s *PublishRequestPgRepoSuite) TestPublishRequestPgRepo_SetMetaSuccess(t provider.T) {
	t.Title("SetMeta: Success")
	t.Tags("PublishRequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание запроса UPDATE
		q := "UPDATE requests SET meta=$1 WHERE request_id=$2"

		meta := PublishRequestMetaPgDTO{
			ReleaseID:    100,
			Grade:        5,
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Description:  "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		s.mock.ExpectExec(q).
			WithArgs(metaJson, uint64(123)).
			WillReturnResult(sqlmock.NewResult(1, 1)) // Симулируем успешное выполнение запроса

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(123).
			WithReleaseID(100).
			WithGrade(5).
			WithDescription("Test description").
			Build()

		// Выполнение метода SetMeta
		err := s.repo.SetMeta(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().NoError(err)
	})
}

func (s *PublishRequestPgRepoSuite) TestPublishRequestPgRepo_SetMetaFailure(t provider.T) {
	t.Title("SetMeta: Failure")
	t.Tags("PublishRequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание запроса UPDATE с ошибкой
		q := "UPDATE requests SET meta=$1 WHERE request_id=$2"

		meta := PublishRequestMetaPgDTO{
			ReleaseID:    100,
			Grade:        5,
			ExpectedDate: cdtime.GetToday().AddDate(0, 0, 14),
			Description:  "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		s.mock.ExpectExec(q).
			WithArgs(metaJson, uint64(123)).
			WillReturnError(sql.ErrConnDone) // Симулируем ошибку при выполнении запроса

		req := data_builder.NewPublishRequestBuilder().
			WithRequestID(123).
			WithReleaseID(100).
			WithGrade(5).
			WithDescription("Test description").
			Build()

		// Выполнение метода SetMeta
		err := s.repo.SetMeta(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
	})
}

func TestPublishRequestPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(PublishRequestPgRepoSuite))
}
