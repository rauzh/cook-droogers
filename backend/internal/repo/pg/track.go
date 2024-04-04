package postgres

import "cookdroogers/models"

//go:generate mockery --name TrackRepo --with-expecter
type TrackRepo interface {
	Create(*models.Track) (uint64, error)
	Get(uint64) (*models.Track, error)
}
