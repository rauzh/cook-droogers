package service

import "cookdroogers/models"

type IStatisticsService interface {
	Create(*models.Statistics) error
	Fetch(tracks []uint64) error
	GetForTrack(uint64) ([]models.Statistics, error)
	GetByID(uint64) (*models.Statistics, error)
	GetRelevantGenre() (string, error)
}
