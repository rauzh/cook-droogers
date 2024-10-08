package service

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"fmt"
	"log/slog"
)

type IManagerService interface {
	Create(*models.Manager) error
	Get(uint64) (*models.Manager, error)
	GetByUserID(uint64) (*models.Manager, error)
	GetRandomManagerID() (uint64, error)
	GetForAdmin() ([]models.Manager, error)
}

type ManagerService struct {
	repo   repo.ManagerRepo
	logger *slog.Logger
}

func NewManagerService(r repo.ManagerRepo, logger *slog.Logger) IManagerService {
	return &ManagerService{repo: r, logger: logger}
}

func (mngSvc *ManagerService) Create(manager *models.Manager) error {
	if err := mngSvc.repo.Create(context.Background(), manager); err != nil {
		mngSvc.logger.Error("MANAGER SERVICE: Create", "error", err.Error())
		return fmt.Errorf("can't create manager with err %w", err)
	}
	return nil
}

func (mngSvc *ManagerService) Get(id uint64) (*models.Manager, error) {
	manager, err := mngSvc.repo.Get(context.Background(), id)

	if err != nil {
		mngSvc.logger.Error("MANAGER SERVICE: Get", "error", err.Error())
		return nil, fmt.Errorf("can't get manager with err %w", err)
	}
	return manager, nil
}

func (mngSvc *ManagerService) GetForAdmin() ([]models.Manager, error) {
	managers, err := mngSvc.repo.GetForAdmin(context.Background())

	if err != nil {
		mngSvc.logger.Error("MANAGER SERVICE: GetForAdmin", "error", err.Error())
		return nil, fmt.Errorf("can't get managers with err %w", err)
	}
	return managers, nil
}

func (mngSvc *ManagerService) GetByUserID(id uint64) (*models.Manager, error) {
	manager, err := mngSvc.repo.GetByUserID(context.Background(), id)

	if err != nil {
		mngSvc.logger.Error("MANAGER SERVICE: GetByUserID", "error", err.Error())
		return nil, fmt.Errorf("can't get manager with err %w", err)
	}
	return manager, nil
}

func (mngSvc *ManagerService) GetRandomManagerID() (uint64, error) {
	id, err := mngSvc.repo.GetRandManagerID(context.Background())

	if err != nil {
		return 0, fmt.Errorf("can't get manager with err %w", err)
	}
	return id, nil
}
