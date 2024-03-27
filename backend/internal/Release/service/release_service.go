package service

import "cookdroogers/models"

type IReleaseService interface {
	Create(release *models.Release, tracks []models.Track) error
	Get(releaseID uint64) (*models.Release, error)
	GetMainGenre(releaseID uint64) (string, error)
	UpdateStatus(uint64, models.ReleaseStatus) error
	GetAllByArtist(uint64) ([]models.Release, error)
	GetAllTracks(release *models.Release) ([]models.Track, error)
}
