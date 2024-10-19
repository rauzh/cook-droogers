package buisness_logic

import (
	"context"
	"cookdroogers/integration_tests/utils"
	"cookdroogers/internal/artist/service"
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
	"log/slog"
	"os"
)

type ArtistIntegrationCoreSuite struct {
	suite.Suite
	db            *sql.DB
	dbx           *sqlx.DB
	txResolver    *trmsqlx.CtxGetter
	trm           transactor.Transactor
	artistService service.IArtistService
	ctx           context.Context
}

func (s *ArtistIntegrationCoreSuite) BeforeEach(t provider.T) {
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

	artistRepo := postgres.NewArtistPgRepo(s.db)

	s.txResolver = trmsqlx.DefaultCtxGetter
	manager, _ := at_manager.New(trmsqlx.NewDefaultFactory(s.dbx))
	s.trm = transactor2.NewATtrm(manager)

	s.artistService = service.NewArtistService(artistRepo, slog.Default())
}

func (s *ArtistIntegrationCoreSuite) AfterEach(t provider.T) {
	_ = s.db.Close()
}

func (s *ArtistIntegrationCoreSuite) TestCreateArtistSuccess(t provider.T) {
	t.Title("CreateArtist: Success")
	t.Tags("ArtistIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			err := s.artistService.Create(txCtx,
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

func (s *ArtistIntegrationCoreSuite) TestCreateArtistFailure(t provider.T) {
	t.Title("CreateArtist: Failure")
	t.Tags("ArtistIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			err := s.artistService.Create(txCtx,
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

func (s *ArtistIntegrationCoreSuite) TestGetArtistSuccess(t provider.T) {
	t.Title("GetArtist: Success")
	t.Tags("ArtistIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			artist, err := s.artistService.Get(txCtx, uint64(1))

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

func (s *ArtistIntegrationCoreSuite) TestGetArtistFailure(t provider.T) {
	t.Title("GetArtist: Failure")
	t.Tags("ArtistIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			artist, err := s.artistService.Get(txCtx, uint64(1000))

			// Assert
			sCtx.Assert().Nil(artist)
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationCoreSuite) TestGetByUserIDArtistSuccess(t provider.T) {
	t.Title("GetByUserIDArtist: Success")
	t.Tags("ArtistIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			artist, err := s.artistService.GetByUserID(txCtx, uint64(6))

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

func (s *ArtistIntegrationCoreSuite) TestGetByUserIDArtistFailure(t provider.T) {
	t.Title("GetByUserIDArtist: Failure")
	t.Tags("ArtistIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			artist, err := s.artistService.GetByUserID(txCtx, uint64(10000))

			// Assert
			sCtx.Assert().Nil(artist)
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationCoreSuite) TestUpdateArtistSuccess(t provider.T) {
	t.Title("UpdateArtist: Success")
	t.Tags("ArtistIntegration")
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
			err := s.artistService.Update(txCtx, newArtist)

			// Assert
			sCtx.Assert().Nil(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *ArtistIntegrationCoreSuite) TestUpdateArtistFailure(t provider.T) {
	t.Title("UpdateArtist: Failure")
	t.Tags("ArtistIntegration")
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
			err := s.artistService.Update(txCtx, newArtist)

			// Assert
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}
