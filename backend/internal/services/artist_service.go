package service

import (
	"cookdroogers/internal/models"
)

type IArtistService interface {
	Create(*models.Artist) error
	Get(uint64) (*models.Artist, error)
	GetAll() ([]models.Artist, error)
	Update(*models.Artist) error
	CreateSignApplication(uint64, string) error
	ApplySignApplication(uint64) error
	DeclineSignApplication(uint64) error
}
