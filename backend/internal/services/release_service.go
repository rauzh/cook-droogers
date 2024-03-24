package service

import "cookdroogers/internal/models"

type IReleaseService interface {
	Create(release *models.Release, tracks []models.Track) error
	Get(releaseID uint64) (*models.Release, error)
	GetMainGenre(releaseID uint64) (string, error)
	Update(*models.Release) error
}
