package service

import "cookdroogers/models"

type IManagerService interface {
	Create(*models.Manager) error
	Get(uint64) (*models.Manager, error)
	GetRandomManagerID() (uint64, error)
}
