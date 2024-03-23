package service

import "cookdroogers/internal/models"

type IStatisticsService interface {
	Create(*models.Statistics) error
	GetForTrack(uint64) ([]models.Statistics, error)
	GetByID(uint64) (*models.Statistics, error)
	GetRelevantGenre() (string, error)
}
