package service

import (
	"cookdroogers/internal/models"
	"cookdroogers/internal/repo"
	"fmt"
)

type ApplicationService struct {
	repo repo.ApplicationRepo
}

func NewApplicationServiceImpl(repo repo.ApplicationRepo) *ApplicationService {
	return &ApplicationService{
		repo: repo,
	}
}

func (as *ApplicationService) Get(app *models.Application, id uint64) error {
	err := as.repo.Get(app, id)
	if err != nil {
		return fmt.Errorf("can't get application info with error %s", err)
	}
	return nil
}

func (as *ApplicationService) GetAllByManagerID(id uint64) ([]*models.Application, error) {
	applicatoins, err := as.repo.GetAllByManagerID(id)
	if err != nil {
		return nil, fmt.Errorf("can't get applications info with error %s", err)
	}
	return applicatoins, nil
}

func (as *ApplicationService) GetAllByUserID(id uint64) ([]*models.Application, error) {
	applicatoins, err := as.repo.GetAllByUserID(id)
	if err != nil {
		return nil, fmt.Errorf("can't get applications info with error %s", err)
	}
	return applicatoins, nil
}
