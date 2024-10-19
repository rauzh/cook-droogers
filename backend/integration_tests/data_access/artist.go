package data_access

import (
	"context"
	"cookdroogers/integration_tests/utils"
	"cookdroogers/internal/repo"
	postgres "cookdroogers/internal/repo/pg"
	"cookdroogers/internal/transactor"
	transactor2 "cookdroogers/internal/transactor/trm"
	"cookdroogers/models/data_builders"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"fmt"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	at_manager "github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"os"
)

type ArtistIntegrationPgSuite struct {
	suite.Suite
	db         *sql.DB
	dbx        *sqlx.DB
	txResolver *trmsqlx.CtxGetter
	trm        transactor.Transactor
	artistRepo repo.ArtistRepo
	ctx        context.Context
}

func (s *ArtistIntegrationPgSuite) BeforeEach(t provider.T) {
	var err error
	pgInfo := utils.PostgresInfo{
		Host:     "postgres",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     "5432",
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	s.db, err = utils.InitDB(pgInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	if s.db == nil {
		return
	}

	s.ctx = context.Background()
	s.dbx = sqlx.NewDb(s.db, "pgx")

	s.artistRepo = postgres.NewArtistPgRepo(s.db)

	s.txResolver = trmsqlx.DefaultCtxGetter
	manager, _ := at_manager.New(trmsqlx.NewDefaultFactory(s.dbx))
	s.trm = transactor2.NewATtrm(manager)
}

func (s *ArtistIntegrationPgSuite) AfterEach(t provider.T) {
	_ = s.db.Close()
}

func (s *ArtistIntegrationPgSuite) TestCreateArtistSuccess(t provider.T) {
	t.Title("CreateArtist: Success")
	t.Tags("ArtistIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			err := s.artistRepo.Create(txCtx,
				data_builders.NewArtistBuilder().
					WithUserID(10).WithManagerID(1).Build())

			// Assert
			sCtx.Assert().NoError(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationPgSuite) TestCreateArtistFailure(t provider.T) {
	t.Title("CreateArtist: Failure")
	t.Tags("ArtistIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			err := s.artistRepo.Create(txCtx,
				data_builders.NewArtistBuilder().
					WithUserID(100).WithManagerID(1).Build())

			// Assert
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationPgSuite) TestGetArtistSuccess(t provider.T) {
	t.Title("GetArtist: Success")
	t.Tags("ArtistIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			artist, err := s.artistRepo.Get(txCtx, uint64(1))

			// Assert
			sCtx.Assert().Nil(err)
			sCtx.Assert().Equal(data_builders.NewArtistBuilder().
				WithID(1).WithUserID(6).
				WithActivity(true).WithContractTerm(cdtime.Date(2029, 10, 10)).
				WithNickname("kodak-black").WithManagerID(1).Build(), artist)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationPgSuite) TestGetArtistFailure(t provider.T) {
	t.Title("GetArtist: Failure")
	t.Tags("ArtistIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			artist, err := s.artistRepo.Get(txCtx, uint64(1000))

			// Assert
			sCtx.Assert().Nil(artist)
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationPgSuite) TestGetByUserIDArtistSuccess(t provider.T) {
	t.Title("GetByUserIDArtist: Success")
	t.Tags("ArtistIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			artist, err := s.artistRepo.GetByUserID(txCtx, uint64(6))

			// Assert
			sCtx.Assert().Nil(err)
			sCtx.Assert().Equal(data_builders.NewArtistBuilder().
				WithID(1).WithUserID(6).
				WithActivity(true).WithContractTerm(cdtime.Date(2029, 10, 10)).
				WithNickname("kodak-black").WithManagerID(1).Build(), artist)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationPgSuite) TestGetByUserIDArtistFailure(t provider.T) {
	t.Title("GetByUserIDArtist: Failure")
	t.Tags("ArtistIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			artist, err := s.artistRepo.GetByUserID(txCtx, uint64(10000))

			// Assert
			sCtx.Assert().Nil(artist)
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationPgSuite) TestUpdateArtistSuccess(t provider.T) {
	t.Title("UpdateArtist: Success")
	t.Tags("ArtistIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			newArtist := data_builders.NewArtistBuilder().
				WithID(1).WithUserID(6).
				WithActivity(true).WithContractTerm(cdtime.Date(2040, 10, 10)).
				WithNickname("kodak-black").WithManagerID(1).Build()

			// Act
			err := s.artistRepo.Update(txCtx, newArtist)

			// Assert
			sCtx.Assert().Nil(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationPgSuite) TestUpdateArtistFailure(t provider.T) {
	t.Title("UpdateArtist: Failure")
	t.Tags("ArtistIntegrationPg")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			newArtist := data_builders.NewArtistBuilder().
				WithID(1).WithUserID(0).
				WithActivity(true).WithContractTerm(cdtime.Date(2040, 10, 10)).
				WithNickname("kodak-black").WithManagerID(1).Build()

			// Act
			err := s.artistRepo.Update(txCtx, newArtist)

			// Assert
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}
