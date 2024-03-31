package service

import (
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"fmt"
)

type IManagerService interface {
	Create(*models.Manager) error
	Get(uint64) (*models.Manager, error)
	GetRandomManagerID() (uint64, error)
}

type ManagerService struct {
	repo repo.ManagerRepo
}

func NewManagerService(r repo.ManagerRepo) IManagerService {
	return &ManagerService{repo: r}
}

func (mngSvc *ManagerService) Create(artist *models.Manager) error {
	if err := mngSvc.repo.Create(artist); err != nil {
		return fmt.Errorf("can't create manager with err %w", err)
	}
	return nil
}

func (mngSvc *ManagerService) Get(id uint64) (*models.Manager, error) {
	manager, err := mngSvc.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get manager with err %w", err)
	}
	return manager, nil
}

func (mngSvc *ManagerService) GetRandomManagerID() (uint64, error) {
	id, err := mngSvc.repo.GetRandManagerID()
	if err != nil {
		return 0, fmt.Errorf("can't get manager with err %w", err)
	}
	return id, nil
}
