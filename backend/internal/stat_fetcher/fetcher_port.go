package fetcher

import "cookdroogers/internal/models"

type StatFetcher interface {
	Fetch(tracks []uint64) ([]models.Statistics, error)
}
