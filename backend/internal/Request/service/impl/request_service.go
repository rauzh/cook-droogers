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
	y, m, d := time.Now().UTC().Date()
	request.Date = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	if err := as.repo.Create(request); err != nil {
		return fmt.Errorf("can't create request.go info with error %w", err)
	}
	return nil
}

func (as *RequestService) Update(request *models.Request) error {
	if err := as.repo.Update(request); err != nil {
		return fmt.Errorf("can't update request.go info with error %w", err)
	}
	return nil
}

func (as *RequestService) Get(id uint64) (*models.Request, error) {
	app, err := as.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get request.go info with error %w", err)
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
