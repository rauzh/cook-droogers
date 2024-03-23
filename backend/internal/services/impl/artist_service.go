package service

import (
	"cookdroogers/internal/models"
	"cookdroogers/internal/repo"
	service "cookdroogers/internal/services"
	"fmt"
	"time"
)

type ArtistService struct {
	applicationService service.IApplicationService
	managerService     service.IManagerService
	userService        service.IUserService
	repo               repo.ArtistRepo
}

func NewArtistService(
	as service.IApplicationService,
	ms service.IManagerService,
	us service.IUserService,
	r repo.ArtistRepo) *ArtistService {
	return &ArtistService{
		applicationService: as,
		managerService:     ms,
		userService:        us,
		repo:               r,
	}
}

func (ars *ArtistService) Create(artist *models.Artist) error {
	if err := ars.repo.Create(artist); err != nil {
		return fmt.Errorf("can't create artist with err %w", err)
	}
	return nil
}

func (ars *ArtistService) Get(id uint64) (*models.Artist, error) {
	artist, err := ars.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get artist with err %w", err)
	}
	return artist, nil
}

func (ars *ArtistService) GetAll() ([]models.Artist, error) {
	artists, err := ars.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("can't get artists with err %w", err)
	}
	return artists, nil
}

func (ars *ArtistService) Update(artist *models.Artist) error {
	if err := ars.repo.Update(artist); err != nil {
		return fmt.Errorf("can't update artist with err %w", err)
	}
	return nil
}

func (ars *ArtistService) CreateSignApplication(userID uint64, nickname string) error {

	if nickname == "" {
		return fmt.Errorf("no nickname provided")
	}

	application := models.Application{
		Type:      models.SignApplication,
		Meta:      map[string]string{"nickname": nickname},
		ApplierID: userID,
	}

	if err := ars.applicationService.Create(&application); err != nil {
		return fmt.Errorf("can't create application with err %w", err)
	}

	go func(managerService service.IManagerService, applicationService service.IApplicationService) {

		application.Status = models.OnApprovalApplication

		managerID, err := managerService.GetRandomManagerID()
		if err == nil {
			application.ManagerID = managerID
		} else {
			application.Status = models.ClosedApplication
			application.Meta["descr"] = "Can't find manager"
		}

		applicationService.Update(&application)

	}(ars.managerService, ars.applicationService)

	return nil
}

func (ars *ArtistService) ApplySignApplication(applicationID uint64) error {

	application, err := ars.applicationService.Get(applicationID)
	if err != nil {
		return fmt.Errorf("can't get application %d with err %w", applicationID, err)
	}

	artist := models.Artist{
		UserID:       application.ApplierID,
		Nickname:     application.Meta["nickname"],
		ContractTerm: time.Now().AddDate(1, 0, 0),
		Activity:     true,
		ManagerID:    application.ManagerID,
	}

	if err := ars.Create(&artist); err != nil {
		return fmt.Errorf("can't create artist %s with err %w", artist.Nickname, err)
	}

	go func() {
		user, err := ars.userService.Get(artist.UserID)
		if err != nil {
			panic("APPLY-SIGN-APPL: Can't get USER")
		}
		user.Type = models.ArtistUser
		ars.userService.Update(user)

		application.Status = models.ClosedApplication
		ars.applicationService.Update(application)
	}()

	return nil
}

func (ars *ArtistService) DeclineSignApplication(applicationID uint64) error {
	application, err := ars.applicationService.Get(applicationID)
	if err != nil {
		return fmt.Errorf("can't get application %d with err %w", applicationID, err)
	}

	application.Status = models.ClosedApplication
	application.Meta["descr"] = "The application is declined."
	ars.applicationService.Update(application)

	return nil
}
