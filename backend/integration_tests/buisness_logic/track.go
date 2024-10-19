package buisness_logic

import (
	"context"
	"cookdroogers/integration_tests/utils"
	postgres "cookdroogers/internal/repo/pg"
	"cookdroogers/internal/track/service"
	"cookdroogers/internal/transactor"
	transactor2 "cookdroogers/internal/transactor/trm"
	"cookdroogers/models/data_builders"
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

type TrackIntegrationCoreSuite struct {
	suite.Suite
	db           *sql.DB
	dbx          *sqlx.DB
	txResolver   *trmsqlx.CtxGetter
	trm          transactor.Transactor
	trackService service.ITrackService
	ctx          context.Context
}

func (s *TrackIntegrationCoreSuite) BeforeEach(t provider.T) {
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

	TrackRepo := postgres.NewTrackPgRepo(s.db)

	s.txResolver = trmsqlx.DefaultCtxGetter
	manager, _ := at_manager.New(trmsqlx.NewDefaultFactory(s.dbx))
	s.trm = transactor2.NewATtrm(manager)

	s.trackService = service.NewTrackService(TrackRepo, slog.Default())
}

func (s *TrackIntegrationCoreSuite) AfterEach(t provider.T) {
	_ = s.db.Close()
}

func (s *TrackIntegrationCoreSuite) TestCreateTrackSuccess(t provider.T) {
	t.Title("CreateTrack: Success")
	t.Tags("TrackIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			_, err := s.trackService.Create(txCtx,
				data_builders.NewTrackBuilder().WithArtists([]uint64{1}).Build())

			// Assert
			sCtx.Assert().NoError(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *TrackIntegrationCoreSuite) TestCreateTrackFailure(t provider.T) {
	t.Title("CreateTrack: Failure")
	t.Tags("TrackIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			_, err := s.trackService.Create(txCtx,
				data_builders.NewTrackBuilder().WithTitle("").Build())

			// Assert
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *TrackIntegrationCoreSuite) TestGetTrackSuccess(t provider.T) {
	t.Title("GetTrack: Success")
	t.Tags("TrackIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			Track, err := s.trackService.Get(txCtx, uint64(1))

			// Assert
			sCtx.Assert().Nil(err)
			sCtx.Assert().Equal(data_builders.NewTrackBuilder().
				WithID(1).WithTitle("oga-boga-1").WithGenre("rock").WithDuration(222).
				WithType("song").WithArtists([]uint64(nil)).Build(), Track)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func (s *TrackIntegrationCoreSuite) TestGetTrackFailure(t provider.T) {
	t.Title("GetTrack: Failure")
	t.Tags("TrackIntegration")
	if utils.IsUnitTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			// Act
			Track, err := s.trackService.Get(txCtx, uint64(1000))

			// Assert
			sCtx.Assert().Nil(Track)
			sCtx.Assert().Error(err)

			// Fixture
			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}
