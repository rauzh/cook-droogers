package repo

import "cookdroogers/internal/models"

type TrackRepo interface {
	Create(*models.Track) (uint64, error)
	Get(uint64) (*models.Track, error)
}
