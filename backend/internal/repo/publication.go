package repo

import (
	"cookdroogers/models"
	"time"
)

//go:generate mockery --name PublicationRepo --with-expecter
type PublicationRepo interface {
	Create(*models.Publication) error
	Get(uint64) (*models.Publication, error)
	GetAllByDate(date time.Time) ([]models.Publication, error)
	GetAllByManager(mng uint64) ([]models.Publication, error)
	GetAllByArtistSinceDate(date time.Time, artistID uint64) ([]models.Publication, error)
	Update(*models.Publication) error
}
