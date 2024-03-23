package service

import "cookdroogers/internal/models"

type IUserService interface {
	Create(*models.User) error
	GetByEmail(string) (*models.User, error)
	Get(uint64) (*models.User, error)
	Update(*models.User) error
}
