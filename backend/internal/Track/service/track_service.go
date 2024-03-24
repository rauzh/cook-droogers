package service

import "cookdroogers/models"

type ITrackService interface {
	Create(*models.Track) (uint64, error)
	Get(uint64) (*models.Track, error)
}
