package service

import (
	"cookdroogers/internal/models"
	"cookdroogers/internal/repo"
	"fmt"
)

type ManagerService struct {
	repo repo.ManagerRepo
}

func NewManagerService(r repo.ManagerRepo) *ManagerService {
	return &ManagerService{repo: r}
}

func (ms *ManagerService) Create(artist *models.Manager) error {
	if err := ms.repo.Create(artist); err != nil {
		return fmt.Errorf("can't create manager with err %w", err)
	}
	return nil
}

func (ms *ManagerService) Get(id uint64) (*models.Manager, error) {
	manager, err := ms.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get manager with err %w", err)
	}
	return manager, nil
}

func (ms *ManagerService) GetRandomManagerID() (uint64, error) {
	id, err := ms.repo.GetRandManagerID()
	if err != nil {
		return 0, fmt.Errorf("can't get manager with err %w", err)
	}
	return id, nil
}
