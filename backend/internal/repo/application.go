package repo

import "cookdroogers/internal/models"

type ApplicationRepo interface {
	Create(*models.Application) error
	Get(*models.Application, uint64) error
	Update(*models.Application) error
	GetAllByManagerID(uint64) ([]*models.Application, error)
	GetAllByUserID(uint64) ([]*models.Application, error)
}
