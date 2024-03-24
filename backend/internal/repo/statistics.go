package repo

import (
	"cookdroogers/internal/models"
	"time"
)

type StatisticsRepo interface {
	Create(*models.Statistics) error
	GetForTrack(uint64) ([]models.Statistics, error)
	GetByID(uint64) (*models.Statistics, error)
	GetAllGroupByTracksSince(date time.Time) (*map[uint64][]models.Statistics, error)
	CreateMany([]models.Statistics) error
}
