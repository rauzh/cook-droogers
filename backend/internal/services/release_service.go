package service

import "cookdroogers/internal/models"

type IReleaseService interface {
	Create(*models.Release) error
	Get(releaseID uint64) (*models.Release, error)
	GetMainGenre(releaseID uint64) (string, error)
	Update(*models.Release) error
}
