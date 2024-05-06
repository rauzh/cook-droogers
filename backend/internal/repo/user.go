package repo

import (
	"context"
	"cookdroogers/models"
)

//go:generate mockery --name UserRepo --with-expecter
type UserRepo interface {
	Create(context.Context, *models.User) error
	GetByEmail(context.Context, string) (*models.User, error)
	Get(context.Context, uint64) (*models.User, error)
	Update(context.Context, *models.User) error
	UpdateType(ctx context.Context, userID uint64, typ models.UserType) error

	SetRole(ctx context.Context, role models.UserType) error
}
