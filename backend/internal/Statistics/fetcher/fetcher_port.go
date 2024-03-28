package fetcher

import "cookdroogers/models"

//go:generate mockery --name StatFetcher --with-expecter
type StatFetcher interface {
	Fetch(tracks []models.Track) ([]models.Statistics, error)
}
