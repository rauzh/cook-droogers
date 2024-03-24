package repo

import "cookdroogers/models"

type ManagerRepo interface {
	Create(*models.Manager) error
	Get(uint64) (*models.Manager, error)
	GetRandManagerID() (uint64, error)
}
