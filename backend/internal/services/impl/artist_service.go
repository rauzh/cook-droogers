package service

import (
	"cookdroogers/internal/models"
	"cookdroogers/internal/repo"
	service "cookdroogers/internal/services"
	"fmt"
)

type ArtistService struct {
	applicationService service.IApplicationService
	managerService     service.IManagerService
	repo               repo.ArtistRepo
}

func (ars *ArtistService) CreateSignApplication(userID uint64, nickname string) error {

	if nickname == "" {
		return fmt.Errorf("no nickname provided")
	}

	application := models.Application{
		Type:      models.SignApplication,
		Meta:      nickname,
		ApplierID: userID,
	}

	if err := ars.applicationService.Create(&application); err != nil {
		return fmt.Errorf("can't create application with err %s", err)
	}

	go func(managerService service.IManagerService, applicationService service.IApplicationService) {

		application.Status = models.OnApprovalApplication
		managerID, err := managerService.GetRandomManagerID()
		if err != nil {
			application.Status = models.ClosedApplication
			application.Meta = "Can't find manager"
		}

		application.ManagerID = managerID

		applicationService.Update(&application)

	}(ars.managerService, ars.applicationService)

	return nil
}
