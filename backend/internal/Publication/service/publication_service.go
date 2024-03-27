package service

import (
	"cookdroogers/models"
	"time"
)

type IPublicationService interface {
	Create(*models.Publication) error
	Update(*models.Publication) error
	Get(uint64) (*models.Publication, error)
	GetAllByDate(date time.Time) ([]models.Publication, error)
	GetAllByManager(managerID uint64) ([]models.Publication, error)
	GetAllByArtistSinceDate(date time.Time, artistID uint64) ([]models.Publication, error)
}
