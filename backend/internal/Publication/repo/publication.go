package repo

import (
	"cookdroogers/models"
	"time"
)

type PublicationRepo interface {
	Create(*models.Publication) error
	Get(uint64) (*models.Publication, error)
	GetAllByDate(date time.Time) ([]models.Publication, error)
	GetAllByArtistSinceDate(date time.Time, artistID uint64) ([]models.Publication, error)
	Update(*models.Publication) error
}
