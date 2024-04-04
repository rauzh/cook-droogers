package postgres

import "cookdroogers/models"

//go:generate mockery --name ManagerRepo --with-expecter
type ManagerRepo interface {
	Create(*models.Manager) error
	Get(uint64) (*models.Manager, error)
	GetRandManagerID() (uint64, error)
}
