package service

import (
	as "cookdroogers/internal/Application/service"
	ars "cookdroogers/internal/Artist/service"
	ms "cookdroogers/internal/Manager/service"
	"cookdroogers/internal/Publication/repo"
	s "cookdroogers/internal/Publication/service"
	rs "cookdroogers/internal/Release/service"
	ss "cookdroogers/internal/Statistics/service"
	"cookdroogers/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type PublicationService struct {
	applicationService as.IApplicationService
	releaseService     rs.IReleaseService
	managerService     ms.IManagerService
	artistService      ars.IArtistService
	statService        ss.IStatisticsService
	repo               repo.PublicationRepo
}

func NewPublicationService(
	as as.IApplicationService,
	rs rs.IReleaseService,
	ms ms.IManagerService,
	ars ars.IArtistService,
	ss ss.IStatisticsService,
	repo repo.PublicationRepo) s.IPublicationService {
	return &PublicationService{
		applicationService: as,
		releaseService:     rs,
		managerService:     ms,
		artistService:      ars,
		statService:        ss,
		repo:               repo,
	}
}

func (ps *PublicationService) Create(publication *models.Publication) error {
	if err := ps.repo.Create(publication); err != nil {
		return fmt.Errorf("can't create publication info with error %w", err)
	}
	return nil
}

func (ps *PublicationService) Update(publication *models.Publication) error {
	if err := ps.repo.Update(publication); err != nil {
		return fmt.Errorf("can't update publication info with error %w", err)
	}
	return nil
}

func (ps *PublicationService) Get(id uint64) (*models.Publication, error) {
	publication, err := ps.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get publication info with error %w", err)
	}
	return publication, nil
}

func (ps *PublicationService) CreatePublApplication(applierID, releaseID uint64, date time.Time) error {

	application := models.Application{
		Type: models.PublishApplication,
		Meta: map[string]string{
			"release": fmt.Sprintf("%d", releaseID),
			"grade":   "5",
			"date":    date.String(),
			"descr":   ""},
		ApplierID: applierID,
	}

	if err := ps.applicationService.Create(&application); err != nil {
		return fmt.Errorf("can't create application with err %w", err)
	}

	go func() {

		application.Status = models.ProcessingApplication
		err := ps.applicationService.Update(&application)
		if err != nil {
			panic("CREATE-PUBL-APPL: Can't update APPL")
		}

		ps.getDegree(releaseID, date, &application)

		artist, err := ps.artistService.Get(applierID)
		if err != nil {
			panic("CREATE-PUBL-APPL: Can't get ARTIST")
		}

		application.ManagerID = artist.ManagerID
		application.Status = models.OnApprovalApplication

		err = ps.applicationService.Update(&application)
		if err != nil {
			panic("CREATE-PUBL-APPL: Can't update APPL")
		}

	}()

	return nil
}

func (ps *PublicationService) getDegree(releaseID uint64, date time.Time, application *models.Application) {
	var grade uint8 = 5

	pubsThatDay, err := ps.repo.GetAllByDate(date)
	if err != nil {
		panic("CREATE-PUBL-APPL: Can't get ALL PUBL BY DATE")
	}

	if len(pubsThatDay) > 1 {
		grade--
		application.Meta["descr"] += "\n Too many releases that day"
	}

	pubsFromThatArtistLastSeason, err := ps.repo.GetAllByArtistSinceDate(date.AddDate(0, -3, 1), application.ApplierID)
	if err != nil {
		panic("CREATE-PUBL-APPL: Can't get ALL PUBL BY ARTIST SINCE DATE")
	}

	if len(pubsFromThatArtistLastSeason) > 2 {
		grade--
		application.Meta["descr"] += "\n Too many releases from this artist"
	}

	relevantGenre, err := ps.statService.GetRelevantGenre()
	if err != nil {
		panic("CREATE-PUBL-APPL: Can't get RELEVANT GENRE")
	}

	currentGenre, err := ps.releaseService.GetMainGenre(releaseID)
	if err != nil {
		panic("CREATE-PUBL-APPL: Can't get CURRENT GENRE")
	}

	if currentGenre != relevantGenre {
		grade--
		application.Meta["descr"] += "\n Not relevant genre"
	}

	application.Meta["grade"] = fmt.Sprintf("%d", grade)
}

func (ps *PublicationService) ApplyPublApplication(applicationID uint64) error {

	application, err := ps.applicationService.Get(applicationID)
	if err != nil {
		return fmt.Errorf("can't get application %d with err %w", applicationID, err)
	}

	releaseID, err := strconv.ParseUint(application.Meta["release"], 10, 64)
	if err != nil {
		return errors.New("can't get release id from publication application")
	}

	date, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", application.Meta["date"])
	if err != nil {
		return errors.New("can't get date from publication application")
	}

	publication := models.Publication{
		ReleaseID: releaseID,
		Date:      date,
		ManagerID: application.ManagerID,
	}

	if err := ps.Create(&publication); err != nil {
		return fmt.Errorf("can't create publication with err %w", err)
	}

	go func() {
		release, err := ps.releaseService.Get(releaseID)
		if err != nil {
			panic("APPLY-SIGN-APPL: Can't get RELEASE")
		}
		release.Status = models.PublishedRelease
		ps.releaseService.Update(release)

		application.Status = models.ClosedApplication
		ps.applicationService.Update(application)
	}()

	return nil
}

func (ps *PublicationService) DeclinePublApplication(applicationID uint64) error {
	application, err := ps.applicationService.Get(applicationID)
	if err != nil {
		return fmt.Errorf("can't get application %d with err %w", applicationID, err)
	}

	application.Status = models.ClosedApplication
	application.Meta["descr"] = "The application is declined."
	ps.applicationService.Update(application)

	return nil
}
