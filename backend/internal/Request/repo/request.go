package repo

import "cookdroogers/models"

//go:generate mockery --name RequestRepo --with-expecter
type RequestRepo interface {
	Create(*models.Request) error
	Get(uint64) (*models.Request, error)
	Update(*models.Request) error
	GetAllByManagerID(uint64) ([]*models.Request, error)
	GetAllByUserID(uint64) ([]*models.Request, error)
}
