package app

import (
	cd_config "cookdroogers/config"
	artistService "cookdroogers/internal/artist/service"
	managerService "cookdroogers/internal/manager/service"
	publicationService "cookdroogers/internal/publication/service"
	releaseService "cookdroogers/internal/release/service"
	"cookdroogers/internal/repo"
	postgres "cookdroogers/internal/repo/pg"
	"cookdroogers/internal/reporter/service"
	"cookdroogers/internal/requests/base"
	repo2 "cookdroogers/internal/requests/base/repo"
	"cookdroogers/internal/requests/base/repo/pg"
	requestService "cookdroogers/internal/requests/base/service"
	"cookdroogers/internal/requests/broker"
	"cookdroogers/internal/requests/broker/sync_broker"
	publishReqRepo "cookdroogers/internal/requests/publish/repo"
	pg2 "cookdroogers/internal/requests/publish/repo/pg"
	usecase2 "cookdroogers/internal/requests/publish/usecase"
	repo3 "cookdroogers/internal/requests/sign_contract/repo"
	postgres2 "cookdroogers/internal/requests/sign_contract/repo/pg"
	"cookdroogers/internal/requests/sign_contract/usecase"
	"cookdroogers/internal/statistics/fetcher/adapters"
	statService "cookdroogers/internal/statistics/service"
	trackService "cookdroogers/internal/track/service"
	"cookdroogers/internal/transactor"
	transactor_impl "cookdroogers/internal/transactor/trm"
	userService "cookdroogers/internal/user/service"
	"database/sql"
	"fmt"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"
)

type App struct {
	postgresDB *sql.DB
	repos      *AppRepositories
	Services   *AppServices
	UseCases   *AppUseCases
	Broker     broker.IBroker
	Transactor transactor.Transactor
	Config     *cd_config.Config
}

type AppUseCases struct {
	SignContractReqUC base.IRequestUseCase
	PublishReqUC      base.IRequestUseCase
}

type AppServices struct {
	ArtistService      artistService.IArtistService
	ManagerService     managerService.IManagerService
	PublicationService publicationService.IPublicationService
	ReleaseService     releaseService.IReleaseService
	ReportService      service.IReportService
	RequestService     requestService.IRequestService
	StatService        statService.IStatisticsService
	TrackService       trackService.ITrackService
	UserService        userService.IUserService
}

type AppRepositories struct {
	artistRepo      repo.ArtistRepo
	managerRepo     repo.ManagerRepo
	publicationRepo repo.PublicationRepo
	releaseRepo     repo.ReleaseRepo
	requestRepo     repo2.RequestRepo
	pubReqRepo      publishReqRepo.PublishRequestRepo
	signReqRepo     repo3.SignContractRequestRepo
	userRepo        repo.UserRepo
	statRepo        repo.StatisticsRepo
	trackRepo       repo.TrackRepo
}

func (a *App) initRepositories() *AppRepositories {

	repos := &AppRepositories{
		artistRepo:      postgres.NewArtistPgRepo(a.postgresDB),
		managerRepo:     postgres.NewManagerPgRepo(a.postgresDB, a.Transactor),
		publicationRepo: postgres.NewPublicationPgRepo(a.postgresDB),
		releaseRepo:     postgres.NewReleasePgRepo(a.postgresDB, a.Transactor),
		requestRepo:     pg.NewRequestPgRepo(a.postgresDB),
		pubReqRepo:      pg2.NewPublishRequestPgRepo(a.postgresDB),
		signReqRepo:     postgres2.NewSignContractRequestPgRepo(a.postgresDB),
		userRepo:        postgres.NewUserPgRepo(a.postgresDB),
		statRepo:        postgres.NewStatisticsPgRepo(a.postgresDB),
		trackRepo:       postgres.NewTrackPgRepo(a.postgresDB),
	}

	return repos
}

func (a *App) initServices() *AppServices {

	trackSvc := trackService.NewTrackService(a.repos.trackRepo)
	rlsSvc := releaseService.NewReleaseService(trackSvc, a.Transactor, a.repos.releaseRepo)
	statFetcher := adapters.NewStatFetcherAdapter(a.Config.StatFetchURLrauzh, a.repos.artistRepo, a.repos.releaseRepo)
	statSvc := statService.NewStatisticsService(trackSvc, statFetcher, a.repos.statRepo, rlsSvc)

	artSvc := artistService.NewArtistService(a.repos.artistRepo)
	mngSvc := managerService.NewManagerService(a.repos.managerRepo)
	pbcSvc := publicationService.NewPublicationService(a.repos.publicationRepo)

	svcs := &AppServices{
		ArtistService:      artSvc,
		ManagerService:     mngSvc,
		PublicationService: pbcSvc,
		TrackService:       trackSvc,
		ReleaseService:     rlsSvc,
		StatService:        statSvc,
		UserService:        userService.NewUserService(a.repos.userRepo),
		RequestService:     requestService.NewRequestService(a.repos.requestRepo),
		ReportService:      service.NewReportService(mngSvc, statSvc, artSvc, pbcSvc, rlsSvc),
	}

	return svcs
}

func (a *App) initUseCases() (*AppUseCases, error) {

	signUC, err := usecase.NewSignContractRequestUseCase(a.repos.userRepo, a.repos.artistRepo,
		a.Transactor, a.Broker, a.repos.signReqRepo)
	if err != nil {
		return nil, err
	}

	pubUC, err := usecase2.NewPublishRequestUseCase(a.Services.StatService, a.repos.publicationRepo, a.repos.releaseRepo,
		a.Transactor, a.Broker, a.repos.pubReqRepo)
	if err != nil {
		return nil, err
	}

	ucs := &AppUseCases{
		SignContractReqUC: signUC,
		PublishReqUC:      pubUC,
	}

	return ucs, nil
}

func (a *App) initDB() (*sql.DB, error) {

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		a.Config.Postgres.User, a.Config.Postgres.DBName, a.Config.Postgres.Password,
		a.Config.Postgres.Host, a.Config.Postgres.Port)
	fmt.Println(dsnPGConn)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		return nil, ErrInitDB
	}

	err = db.Ping()
	if err != nil {
		return nil, ErrInitDB
	}

	db.SetMaxOpenConns(10)

	return db, nil
}

func (a *App) Init() error {

	db, err := a.initDB()
	if err != nil {
		return err
	}

	a.postgresDB = db

	var trmng transactor.Transactor
	trmng = transactor_impl.NewATtrm(manager.Must(trmsqlx.NewDefaultFactory(sqlx.NewDb(db, "pgx"))))

	a.Transactor = trmng

	syncbroker, err := sync_broker.NewSyncBroker(a.Config.Kafka.KafkaEndpoints, a.Config.Kafka.KafkaSettings)
	if err != nil {
		return err
	}

	a.Broker = syncbroker

	a.repos = a.initRepositories()
	a.Services = a.initServices()

	ucs, err := a.initUseCases()
	if err != nil {
		return err
	}

	a.UseCases = ucs

	return nil
}
