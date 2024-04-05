package repo

import (
	"context"
	"cookdroogers/models"
	"time"
)

//go:generate mockery --name PublicationRepo --with-expecter
type PublicationRepo interface {
	Create(context.Context, *models.Publication) error
	Get(context.Context, uint64) (*models.Publication, error)
	GetAllByDate(context.Context, time.Time) ([]models.Publication, error)
	GetAllByManager(ctx context.Context, mng uint64) ([]models.Publication, error)
	GetAllByArtistSinceDate(ctx context.Context, date time.Time, artistID uint64) ([]models.Publication, error)
	Update(context.Context, *models.Publication) error
}
