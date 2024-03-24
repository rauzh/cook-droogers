package service

import (
	"cookdroogers/internal/models"
	"cookdroogers/internal/repo"
	fetcher "cookdroogers/internal/stat_fetcher"
	"errors"
	"fmt"
	"time"
)

type StatisticsService struct {
	fetcher fetcher.StatFetcher
	repo    repo.StatisticsRepo
}

func (ss *StatisticsService) Create(stat *models.Statistics) error {

	stat.Date = time.Now()

	if err := ss.repo.Create(stat); err != nil {
		return fmt.Errorf("can't create stats with err %w", err)
	}
	return nil
}

func (ss *StatisticsService) GetByID(statID uint64) (*models.Statistics, error) {
	stat, err := ss.repo.GetByID(statID)
	if err != nil {
		return nil, fmt.Errorf("can't get stats with err %w", err)
	}
	return stat, nil
}

func (ss *StatisticsService) GetForTrack(trackID uint64) ([]models.Statistics, error) {
	stats, err := ss.repo.GetForTrack(trackID)
	if err != nil {
		return nil, fmt.Errorf("can't get stats for track %d with err %w", trackID, err)
	}
	return stats, nil
}

func (ss *StatisticsService) Fetch(tracks []uint64) error {

	stats, err := ss.fetcher.Fetch(tracks)
	if err != nil {
		return fmt.Errorf("can't fetch stats with err %w", err)
	}

	if len(stats) < 1 {
		return errors.New("no stats to fetch")
	}

	for _, stat := range stats {
		stat.Date = time.Now()
	}

	if err = ss.repo.CreateMany(stats); err != nil {
		return fmt.Errorf("can't create stats with err %w", err)
	}

	return nil
}
