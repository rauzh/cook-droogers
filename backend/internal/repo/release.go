package repo

import (
	"context"
	"cookdroogers/models"
)

//go:generate mockery --name ReleaseRepo --with-expecter
type ReleaseRepo interface {
	Create(context.Context, *models.Release) error
	Get(context.Context, uint64) (*models.Release, error)
	GetAllByArtist(ctx context.Context, artistID uint64) ([]models.Release, error)
	GetAllTracks(ctx context.Context, release *models.Release) ([]models.Track, error)
	Update(context.Context, *models.Release) error
	UpdateStatus(ctx context.Context, id uint64, stat models.ReleaseStatus) error
}
