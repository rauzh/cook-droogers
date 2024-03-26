package service

import "cookdroogers/models"

type IRequestService interface {
	Create(*models.Request) error
	Update(*models.Request) error
	GetAllByUserID(uint64) ([]*models.Request, error)
	GetAllByManagerID(uint64) ([]*models.Request, error)
	Get(uint64) (*models.Request, error)
}
