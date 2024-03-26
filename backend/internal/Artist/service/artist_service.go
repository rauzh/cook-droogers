package service

import "cookdroogers/models"

type IArtistService interface {
	Create(*models.Artist) error
	Get(uint64) (*models.Artist, error)
	GetAll() ([]models.Artist, error)
	Update(*models.Artist) error
}
