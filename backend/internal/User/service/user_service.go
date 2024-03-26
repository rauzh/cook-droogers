package service

import "cookdroogers/models"

type IUserService interface {
	Create(*models.User) error
	Login(login, password string) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	Get(uint64) (*models.User, error)
	Update(*models.User) error
	UpdateType(uint64, models.UserType) error
}
