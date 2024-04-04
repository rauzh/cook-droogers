package postgres

import "cookdroogers/models"

//go:generate mockery --name ReleaseRepo --with-expecter
type ReleaseRepo interface {
	Create(*models.Release) error
	Get(uint64) (*models.Release, error)
	GetAllByArtist(artistID uint64) ([]models.Release, error)
	GetAllTracks(release *models.Release) ([]models.Track, error)
	Update(*models.Release) error
	UpdateStatus(id uint64, stat models.ReleaseStatus) error
}
