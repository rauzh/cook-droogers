package repo

import "cookdroogers/internal/models"

type UserRepo interface {
	Create(*models.User) error
	GetByEmail(string) (*models.User, error)
	Get(uint64) (*models.User, error)
	Update(*models.User) error
}
