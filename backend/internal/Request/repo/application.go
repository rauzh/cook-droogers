package repo

import "cookdroogers/models"

type RequestRepo interface {
	Create(*models.Request) error
	Get(uint64) (*models.Request, error)
	Update(*models.Request) error
	GetAllByManagerID(uint64) ([]*models.Request, error)
	GetAllByUserID(uint64) ([]*models.Request, error)
}
