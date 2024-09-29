package postgres

import (
	"context"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/sign_contract"
	"cookdroogers/internal/requests/sign_contract/data_builder"
	"cookdroogers/internal/requests/sign_contract/repo"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type SignContractRequestPgRepoSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repo.SignContractRequestRepo
	ctx  context.Context
}

func (s *SignContractRequestPgRepoSuite) BeforeEach(t provider.T) {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	s.repo = NewSignContractRequestPgRepo(s.db)
	s.ctx = context.Background()
}

func (s *SignContractRequestPgRepoSuite) AfterEach(t provider.T) {
	s.db.Close()
}

func (s *SignContractRequestPgRepoSuite) TestSignContractRequestPgRepo_CreateSuccess(t provider.T) {
	t.Title("Create: Success")
	t.Tags("SignContractRequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание запроса (INSERT INTO requests)
		q := "INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING request_id"

		meta := SignContractReqMetaPgDTO{
			Nickname:    "leclerc",
			Description: "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		s.mock.ExpectQuery(q).
			WithArgs(base.ProcessingRequest, sign_contract.SignRequest,
				cdtime.GetToday(), metaJson, uint64(8), uint64(88)).
			WillReturnRows(sqlmock.NewRows([]string{"request_id"}).AddRow(123))

		req := data_builder.NewSignContractRequestBuilder().Build()

		// Выполнение метода Create
		err := s.repo.Create(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(uint64(123), req.RequestID) // Проверяем, что ID был правильно присвоен
	})
}

func (s *SignContractRequestPgRepoSuite) TestSignContractRequestPgRepo_CreateFailure(t provider.T) {
	t.Title("Create: Failure")
	t.Tags("SignContractRequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание запроса (INSERT INTO requests)
		q := "INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING request_id"

		meta := SignContractReqMetaPgDTO{
			Nickname:    "leclerc",
			Description: "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		s.mock.ExpectQuery(q).
			WithArgs(base.ProcessingRequest, sign_contract.SignRequest,
				cdtime.GetToday(), metaJson, uint64(8), uint64(88)).
			WillReturnError(sql.ErrConnDone)

		req := data_builder.NewSignContractRequestBuilder().Build()

		// Выполнение метода Create
		err := s.repo.Create(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
	})
}

func (s *SignContractRequestPgRepoSuite) TestSignContractRequestPgRepo_UpdateSuccess(t provider.T) {
	t.Title("Update: Success")
	t.Tags("SignContractRequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание запроса (UPDATE requests)
		q := "UPDATE requests SET status=$2, type=$3, meta=$4, manager_id=$5, user_id=$6 WHERE request_id=$1"

		meta := SignContractReqMetaPgDTO{
			Nickname:    "leclerc",
			Description: "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		// Эмулируем успешное выполнение запроса
		s.mock.ExpectQuery(q).
			WithArgs(uint64(123), base.ProcessingRequest, sign_contract.SignRequest,
									metaJson, uint64(8), uint64(88)).
			WillReturnRows(sqlmock.NewRows([]string{})) // Успешный результат без строк

		req := data_builder.NewSignContractRequestBuilder().Build()

		// Выполнение метода Update
		err := s.repo.Update(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().NoError(err)
	})
}

func (s *SignContractRequestPgRepoSuite) TestSignContractRequestPgRepo_UpdateFailure(t provider.T) {
	t.Title("Update: Failure")
	t.Tags("SignContractRequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание запроса (UPDATE requests) с ошибкой
		q := "UPDATE requests SET status=$2, type=$3, meta=$4, manager_id=$5, user_id=$6 WHERE request_id=$1"

		meta := SignContractReqMetaPgDTO{
			Nickname:    "leclerc",
			Description: "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		// Эмулируем ошибку при выполнении запроса
		s.mock.ExpectQuery(q).
			WithArgs(uint64(123), base.ProcessingRequest, sign_contract.SignRequest,
								metaJson, uint64(8), uint64(88)).
			WillReturnError(sql.ErrConnDone) // Симулируем ошибку соединения

		req := data_builder.NewSignContractRequestBuilder().Build()

		// Выполнение метода Update
		err := s.repo.Update(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
	})
}

func (s *SignContractRequestPgRepoSuite) TestSignContractRequestPgRepo_SetMetaSuccess(t provider.T) {
	t.Title("SetMeta: Success")
	t.Tags("SignContractRequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание запроса (UPDATE requests SET meta)
		q := "UPDATE requests SET meta=$1 WHERE request_id=$2"

		meta := SignContractReqMetaPgDTO{
			Nickname:    "leclerc",
			Description: "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		// Эмулируем успешное выполнение запроса
		s.mock.ExpectQuery(q).
			WithArgs(metaJson, uint64(123)).WillReturnRows(sqlmock.NewRows([]string{}))

		req := data_builder.NewSignContractRequestBuilder().Build()

		// Выполнение метода SetMeta
		err := s.repo.SetMeta(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().NoError(err)
	})
}

func (s *SignContractRequestPgRepoSuite) TestSignContractRequestPgRepo_SetMetaFailure(t provider.T) {
	t.Title("SetMeta: Failure")
	t.Tags("SignContractRequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание запроса (UPDATE requests SET meta)
		q := "UPDATE requests SET meta=$1 WHERE request_id=$2"

		meta := SignContractReqMetaPgDTO{
			Nickname:    "leclerc",
			Description: "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		// Эмулируем ошибку при выполнении запроса
		s.mock.ExpectQuery(q).
			WithArgs(metaJson, uint64(123)).
			WillReturnError(sql.ErrConnDone)

		req := data_builder.NewSignContractRequestBuilder().Build()

		// Выполнение метода SetMeta
		err := s.repo.SetMeta(s.ctx, req)

		// Проверка результатов
		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
	})
}

func (s *SignContractRequestPgRepoSuite) TestSignContractRequestPgRepo_GetSuccess(t provider.T) {
	t.Title("Get: Success")
	t.Tags("SignContractRequestPgRepo")
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		// Ожидание запроса (SELECT * FROM requests WHERE request_id)
		q := "SELECT * FROM requests WHERE request_id=$1"

		meta := SignContractReqMetaPgDTO{
			Nickname:    "leclerc",
			Description: "Test description",
		}
		metaJson, _ := json.Marshal(meta)

		mngID := sql.NullInt64{Int64: int64(8), Valid: true}

		// Эмулируем успешное выполнение запроса
		s.mock.ExpectQuery(q).
			WithArgs(uint64(123)).
			WillReturnRows(sqlmock.NewRows([]string{
				"request_id", "status", "type", "creation_date", "meta", "manager_id", "user_id",
			}).AddRow(uint64(123), base.ProcessingRequest, sign_contract.SignRequest,
				cdtime.GetToday(), metaJson, mngID, uint64(88)))

		// Выполнение метода Get
		req, err := s.repo.Get(s.ctx, uint64(123))

		// Проверка результатов
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(req)
		sCtx.Assert().Equal(uint64(123), req.RequestID)
		sCtx.Assert().Equal("leclerc", req.Nickname)
		sCtx.Assert().Equal("Test description", req.Description)
	})
}

func (s *SignContractRequestPgRepoSuite) TestSignContractRequestPgRepo_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("SignContractRequestPgRepo")
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		// Ожидание запроса (SELECT * FROM requests WHERE request_id)
		q := "SELECT * FROM requests WHERE request_id=$1"

		// Эмулируем ошибку при выполнении запроса
		s.mock.ExpectQuery(q).
			WithArgs(uint64(123)).
			WillReturnError(sql.ErrConnDone)

		// Выполнение метода Get
		req, err := s.repo.Get(s.ctx, uint64(123))

		// Проверка результатов
		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
		sCtx.Assert().Nil(req)
	})
}

func TestSignContractRequestPgRepoSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(SignContractRequestPgRepoSuite))
}
