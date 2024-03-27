package fetcher

import "cookdroogers/models"

//go:generate mockery --name StatFetcher --with-expecter
type StatFetcher interface {
	Fetch(tracks []uint64) ([]models.Statistics, error)
}
