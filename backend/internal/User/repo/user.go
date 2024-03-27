package repo

import "cookdroogers/models"

//go:generate mockery --name UserRepo --with-expecter
type UserRepo interface {
	Create(*models.User) error
	GetByEmail(string) (*models.User, error)
	Get(uint64) (*models.User, error)
	Update(*models.User) error
	UpdateType(userID uint64, typ models.UserType) error
}
