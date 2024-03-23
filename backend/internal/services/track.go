package service

import "cookdroogers/internal/models"

type ITrackService interface {
	Create(*models.Track) error
	Get(uint64) (*models.Track, error)
}
