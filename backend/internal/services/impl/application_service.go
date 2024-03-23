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

func (as *ApplicationService) Create(application *models.Application) error {
	if err := as.repo.Create(application); err != nil {
		return fmt.Errorf("can't create application info with error %w", err)
	}
	return nil
}

func (as *ApplicationService) Update(application *models.Application) error {
	if err := as.repo.Update(application); err != nil {
		return fmt.Errorf("can't update application info with error %w", err)
	}
	return nil
}

func (as *ApplicationService) Get(id uint64) (*models.Application, error) {
	app, err := as.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get application info with error %w", err)
	}
	return app, nil
}

func (as *ApplicationService) GetAllByManagerID(id uint64) ([]*models.Application, error) {
	applicatoins, err := as.repo.GetAllByManagerID(id)
	if err != nil {
		return nil, fmt.Errorf("can't get applications info with error %w", err)
	}
	return applicatoins, nil
}

func (as *ApplicationService) GetAllByUserID(id uint64) ([]*models.Application, error) {
	applicatoins, err := as.repo.GetAllByUserID(id)
	if err != nil {
		return nil, fmt.Errorf("can't get applications info with error %w", err)
	}
	return applicatoins, nil
}
