package repo

import "cookdroogers/internal/models"

type ArtistRepo interface {
	Create(*models.Artist) error
	Get(*models.Artist, uint64) error
	Update(*models.Artist) error
}
