package repo

import (
	"context"
	"cookdroogers/models"
)

//go:generate mockery --name ManagerRepo --with-expecter
type ManagerRepo interface {
	Create(context.Context, *models.Manager) error
	Get(context.Context, uint64) (*models.Manager, error)
	GetRandManagerID(context.Context) (uint64, error)
	GetByUserID(ctx context.Context, userID uint64) (*models.Manager, error)
	GetForAdmin(ctx context.Context) ([]models.Manager, error)
}
