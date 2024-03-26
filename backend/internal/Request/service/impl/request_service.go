package service

import (
	requestRepo "cookdroogers/internal/Request/repo"
	s "cookdroogers/internal/Request/service"
	"cookdroogers/models"
	"fmt"
	"time"
)

type RequestService struct {
	repo requestRepo.RequestRepo
}

func NewRequestServiceImpl(repo requestRepo.RequestRepo) s.IRequestService {
	return &RequestService{
		repo: repo,
	}
}

func (as *RequestService) Create(request *models.Request) error {

	request.Status = models.NewRequest
	request.Date = time.Now()

	if err := as.repo.Create(request); err != nil {
		return fmt.Errorf("can't create request info with error %w", err)
	}
	return nil
}

func (as *RequestService) Update(request *models.Request) error {
	if err := as.repo.Update(request); err != nil {
		return fmt.Errorf("can't update request info with error %w", err)
	}
	return nil
}

func (as *RequestService) Get(id uint64) (*models.Request, error) {
	app, err := as.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get request info with error %w", err)
	}
	return app, nil
}

func (as *RequestService) GetAllByManagerID(id uint64) ([]*models.Request, error) {
	applicatoins, err := as.repo.GetAllByManagerID(id)
	if err != nil {
		return nil, fmt.Errorf("can't get requests info with error %w", err)
	}
	return applicatoins, nil
}

func (as *RequestService) GetAllByUserID(id uint64) ([]*models.Request, error) {
	applicatoins, err := as.repo.GetAllByUserID(id)
	if err != nil {
		return nil, fmt.Errorf("can't get requests info with error %w", err)
	}
	return applicatoins, nil
}
