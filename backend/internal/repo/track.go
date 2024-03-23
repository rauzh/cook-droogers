package repo

import "cookdroogers/internal/models"

type TrackRepo interface {
	Create(*models.Track) error
	Get(uint64) (*models.Track, error)
}
