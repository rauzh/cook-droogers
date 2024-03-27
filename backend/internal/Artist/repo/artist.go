package repo

import "cookdroogers/models"

//go:generate mockery --name ArtistRepo --with-expecter
type ArtistRepo interface {
	Create(*models.Artist) error
	Get(uint64) (*models.Artist, error)
	GetAll() ([]models.Artist, error)
	Update(*models.Artist) error
}
