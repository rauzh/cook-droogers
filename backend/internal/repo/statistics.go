package repo

import "cookdroogers/internal/models"

type StatisticsRepo interface {
	Create(*models.Statistics) error
	GetForTrack(uint64) ([]models.Statistics, error)
	GetByID(uint64) (*models.Statistics, error)
	GetAll() ([]models.Statistics, error)
	CreateMany([]models.Statistics) error
}
