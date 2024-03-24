package repo

import "cookdroogers/models"

type ApplicationRepo interface {
	Create(*models.Application) error
	Get(uint64) (*models.Application, error)
	Update(*models.Application) error
	GetAllByManagerID(uint64) ([]*models.Application, error)
	GetAllByUserID(uint64) ([]*models.Application, error)
}
