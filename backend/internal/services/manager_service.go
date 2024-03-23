package service

import "cookdroogers/internal/models"

type IManagerService interface {
	Create(*models.Manager) error
	Get(uint64) (*models.Manager, error)
	GetRandomManagerID() (uint64, error)
}
