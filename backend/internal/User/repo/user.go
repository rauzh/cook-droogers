package repo

import "cookdroogers/models"

type UserRepo interface {
	Create(*models.User) error
	GetByEmail(string) (*models.User, error)
	Get(uint64) (*models.User, error)
	Update(*models.User) error
}
