package repo

import (
	"context"
	"cookdroogers/models"
)

//go:generate mockery --name ArtistRepo --with-expecter
type ArtistRepo interface {
	Create(context.Context, *models.Artist) error
	Get(context.Context, uint64) (*models.Artist, error)
	GetByUserID(context.Context, uint64) (*models.Artist, error)
	Update(context.Context, *models.Artist) error
}
