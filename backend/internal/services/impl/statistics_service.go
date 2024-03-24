package service

import (
	"cookdroogers/internal/models"
	"cookdroogers/internal/repo"
	service "cookdroogers/internal/services"
	fetcher "cookdroogers/internal/stat_fetcher"
	"errors"
	"fmt"
	"time"
)

type StatisticsService struct {
	trackService service.ITrackService
	fetcher      fetcher.StatFetcher
	repo         repo.StatisticsRepo
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

func (ss *StatisticsService) GetRelevantGenre() (string, error) {
	/*  По-хорошему надо конечно выгружать из БД только пачками по 100,
	и распараллелить по данным, но мне влом, а если прям надо, то сделаю */

	var err error
	var stats map[uint64][]models.Statistics

	stats, err = ss.repo.GetAllGroupByTracksSince(time.Now().AddDate(0, -3, 0))
	if err != nil {
		return "", fmt.Errorf("can't get stats with err %w", err)
	}

	genres := map[string]uint64{}
	for trackID, statsPerTrack := range stats {
		track, err := ss.trackService.Get(trackID)
		if err != nil {
			return "", fmt.Errorf("can't get track %d with err %w", trackID, err)
		}

		for _, stat := range statsPerTrack {
			genres[track.Genre] += stat.Streams
		}
	}

	var relevantGenre string
	var maxStreamsPerGenre uint64
	for genre, streams := range genres {
		if streams > maxStreamsPerGenre {
			maxStreamsPerGenre = streams
			relevantGenre = genre
		}
	}

	return relevantGenre, err
}
