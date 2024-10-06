package service

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"github.com/pkg/errors"
	"log/slog"
)

type IManagerService interface {
	Create(context.Context, *models.Manager) error
	Get(context.Context, uint64) (*models.Manager, error)
	GetByUserID(context.Context, uint64) (*models.Manager, error)
	GetRandomManagerID(context.Context) (uint64, error)
	GetForAdmin(context.Context) ([]models.Manager, error)
}

var CreateDbError error = errors.New("can't create manager")
var GetDbError error = errors.New("can't get manager")

type ManagerService struct {
	repo   repo.ManagerRepo
	logger *slog.Logger
}

func NewManagerService(r repo.ManagerRepo, logger *slog.Logger) IManagerService {
	return &ManagerService{repo: r, logger: logger}
}

func (mngSvc *ManagerService) Create(ctx context.Context, manager *models.Manager) error {
	if err := mngSvc.repo.Create(ctx, manager); err != nil {
		mngSvc.logger.Error("MANAGER SERVICE: Create", "error", err.Error())
		return errors.Wrap(CreateDbError, err.Error())
	}
	return nil
}

func (mngSvc *ManagerService) Get(ctx context.Context, id uint64) (*models.Manager, error) {
	manager, err := mngSvc.repo.Get(ctx, id)

	if err != nil {
		mngSvc.logger.Error("MANAGER SERVICE: Get", "error", err.Error())
		return nil, errors.Wrap(GetDbError, err.Error())
	}
	return manager, nil
}

func (mngSvc *ManagerService) GetForAdmin(ctx context.Context) ([]models.Manager, error) {
	managers, err := mngSvc.repo.GetForAdmin(ctx)

	if err != nil {
		mngSvc.logger.Error("MANAGER SERVICE: GetForAdmin", "error", err.Error())
		return nil, errors.Wrap(GetDbError, err.Error())
	}
	return managers, nil
}

func (mngSvc *ManagerService) GetByUserID(ctx context.Context, id uint64) (*models.Manager, error) {
	manager, err := mngSvc.repo.GetByUserID(ctx, id)

	if err != nil {
		mngSvc.logger.Error("MANAGER SERVICE: GetByUserID", "error", err.Error())
		return nil, errors.Wrap(GetDbError, err.Error())
	}
	return manager, nil
}

func (mngSvc *ManagerService) GetRandomManagerID(ctx context.Context) (uint64, error) {
	id, err := mngSvc.repo.GetRandManagerID(ctx)

	if err != nil {
		return 0, errors.Wrap(GetDbError, err.Error())
	}
	return id, nil
}
