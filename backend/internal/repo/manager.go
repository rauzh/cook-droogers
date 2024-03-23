package repo

import "cookdroogers/internal/models"

type ManagerRepo interface {
	Create(*models.Manager) error
	Get(uint64) (*models.Manager, error)
	GetRandManagerID() (uint64, error)
}
