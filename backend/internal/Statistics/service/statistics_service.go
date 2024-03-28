package service

import (
	"cookdroogers/models"
)

type IStatisticsService interface {
	Create(*models.Statistics) error
	FetchByRelease(release *models.Release) error
	GetForTrack(uint64) ([]models.Statistics, error)
	GetByID(uint64) (*models.Statistics, error)
	GetRelevantGenre() (string, error)
	GetLatestStatForTrack(trackID uint64) (*models.Statistics, error)
}
