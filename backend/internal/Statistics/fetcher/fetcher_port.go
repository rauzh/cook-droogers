package fetcher

import "cookdroogers/models"

type StatFetcher interface {
	Fetch(tracks []uint64) ([]models.Statistics, error)
}
