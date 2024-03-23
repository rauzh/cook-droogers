package service

import "cookdroogers/internal/models"

type IApplicationService interface {
	Create(*models.Application) error
	Update(*models.Application) error
	GetAllByUserID(uint64) ([]*models.Application, error)
	GetAllByManagerID(uint64) ([]*models.Application, error)
	Get(uint64) (*models.Application, error)
	ApplyPublication(*models.Application) error
	DeclinePublication(*models.Application) error
}
