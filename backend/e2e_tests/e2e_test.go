//go:build e2e
// +build e2e

package integration_tests

import (
	"context"
	"cookdroogers/integration_tests/utils"
	artist_service "cookdroogers/internal/artist/service"
	release_service "cookdroogers/internal/release/service"
	postgres "cookdroogers/internal/repo/pg"
	"cookdroogers/internal/track/service"
	"cookdroogers/internal/transactor"
	transactor2 "cookdroogers/internal/transactor/trm"
	user_service "cookdroogers/internal/user/service"
	"cookdroogers/models"
	"cookdroogers/models/data_builders"
	"cookdroogers/pkg/time"
	"database/sql"
	"fmt"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	at_manager "github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"log/slog"
	"os"
	"testing"
)

type E2ESuite struct {
	suite.Suite
	db             *sql.DB
	dbx            *sqlx.DB
	txResolver     *trmsqlx.CtxGetter
	trm            transactor.Transactor
	userService    user_service.IUserService
	artistService  artist_service.IArtistService
	releaseService release_service.IReleaseService
	ctx            context.Context
}

func (s *E2ESuite) BeforeEach(t provider.T) {
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

	s.artistService = artist_service.NewArtistService(artistRepo, slog.Default())

	userRepo := postgres.NewUserPgRepo(s.db)

	s.userService = user_service.NewUserService(userRepo, slog.Default())

	trackRepo := postgres.NewTrackPgRepo(s.db)

	trackService := service.NewTrackService(trackRepo, slog.Default())

	releaseRepo := postgres.NewReleasePgRepo(s.db, s.trm)

	s.releaseService = release_service.NewReleaseService(trackService, s.trm, releaseRepo, slog.Default())
}

func (s *E2ESuite) AfterEach(t provider.T) {
	_ = s.db.Close()
}

func BeforeAll() {
	pgInfo := utils.PostgresInfo{
		Host:     "postgres",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     "5432",
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	db, err := utils.InitDB(pgInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := os.ReadFile("/builds/rpp21u198/test-cd-73/backend/integration_tests/init.sql")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(string(text))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *E2ESuite) TestE2EScenarioLoginGetUploadSuccess(t provider.T) {
	t.Title("E2E_ScenarioLoginGetUpload: Success")
	t.Tags("E2E")
	if utils.IsUnitTestsFailed() || utils.IsIntegrationTestsFailed() {
		t.Skip()
	}
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		_ = s.trm.WithinTransaction(s.ctx, func(txCtx context.Context) error {

			user, err := s.userService.Login(txCtx, "uzi@ppo.ru", "123")
			sCtx.Assert().NoError(err)
			sCtx.Assert().Equal(
				data_builders.NewUserBuilder().
					WithUserID(7).
					WithEmail("uzi@ppo.ru").
					WithName("uzi").
					WithPassword("123").
					WithType(models.ArtistUser).Build(), user)

			err = s.userService.SetRole(txCtx, user.Type)
			sCtx.Assert().NoError(err)

			artist, err := s.artistService.GetByUserID(txCtx, user.UserID)
			sCtx.Assert().NoError(err)
			sCtx.Assert().Equal(
				data_builders.NewArtistBuilder().
					WithUserID(7).
					WithNickname("lil-uzi-vert").
					WithID(2).
					WithContractTerm(time.Date(2029, 12, 12)).
					WithActivity(true).
					WithManagerID(1).Build(), artist)

			releases, err := s.releaseService.GetAllByArtist(txCtx, artist.ArtistID)

			sCtx.Assert().Nil(err)
			sCtx.Assert().Len(releases, 0)

			utils.Rollback(txCtx, s.txResolver, s.dbx)
			return nil
		})
	})
}

func TestE2ERunner(t *testing.T) {
	BeforeAll()

	suite.RunSuite(t, new(E2ESuite))
}
