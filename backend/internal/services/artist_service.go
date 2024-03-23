package service

import (
	"cookdroogers/internal/models"
)

type IArtistService interface {
	Create(*models.Artist) error
	CreateSignApplication(uint64, string) error
	ApplySignApplication(*models.Application) error
	DeclineSignApplication(*models.Application) error
}
