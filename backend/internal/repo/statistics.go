package repo

import (
	"cookdroogers/models"
	"time"
)

//go:generate mockery --name StatisticsRepo --with-expecter
type StatisticsRepo interface {
	Create(*models.Statistics) error
	GetForTrack(uint64) ([]models.Statistics, error)
	GetByID(uint64) (*models.Statistics, error)
	GetAllGroupByTracksSince(date time.Time) (*map[uint64][]models.Statistics, error)
	CreateMany([]models.Statistics) error
}
