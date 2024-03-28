package service

import (
	"cookdroogers/internal/Statistics/fetcher"
	"cookdroogers/internal/Statistics/repo"
	s "cookdroogers/internal/Statistics/service"
	ts "cookdroogers/internal/Track/service"
	"cookdroogers/models"
	"errors"
	"fmt"
	"time"
)

type StatisticsService struct {
	trackService ts.ITrackService
	fetcher      fetcher.StatFetcher
	repo         repo.StatisticsRepo
}

func NewStatisticsService(
	ts ts.ITrackService,
	f fetcher.StatFetcher,
	r repo.StatisticsRepo) s.IStatisticsService {
	return &StatisticsService{
		trackService: ts,
		fetcher:      f,
		repo:         r,
	}
}

func (statSvc *StatisticsService) Create(stat *models.Statistics) error {

	stat.Date = time.Now()

	if err := statSvc.repo.Create(stat); err != nil {
		return fmt.Errorf("can't create stats with err %w", err)
	}
	return nil
}

func (statSvc *StatisticsService) GetByID(statID uint64) (*models.Statistics, error) {
	stat, err := statSvc.repo.GetByID(statID)
	if err != nil {
		return nil, fmt.Errorf("can't get stats with err %w", err)
	}
	return stat, nil
}

func (statSvc *StatisticsService) GetForTrack(trackID uint64) ([]models.Statistics, error) {
	stats, err := statSvc.repo.GetForTrack(trackID)
	if err != nil {
		return nil, fmt.Errorf("can't get stats for track %d with err %w", trackID, err)
	}
	return stats, nil
}

func (statSvc *StatisticsService) Fetch(tracks []uint64) error {

	stats, err := statSvc.fetcher.Fetch(tracks)
	if err != nil {
		return fmt.Errorf("can't fetch stats with err %w", err)
	}

	if len(stats) < 1 {
		return errors.New("no stats to fetch")
	}

	for _, stat := range stats {
		stat.Date = time.Now()
	}

	if err = statSvc.repo.CreateMany(stats); err != nil {
		return fmt.Errorf("can't create stats with err %w", err)
	}

	return nil
}

func (statSvc *StatisticsService) GetRelevantGenre() (string, error) {
	/*  По-хорошему надо конечно выгружать из БД только пачками по 100,
	и распараллелить по данным, но мне влом, а если прям надо, то сделаю */

	stats, err := statSvc.repo.GetAllGroupByTracksSince(time.Now().AddDate(0, -3, 0))
	if err != nil {
		return "", fmt.Errorf("can't get stats with err %w", err)
	}

	genres := make(map[string]uint64)
	for trackID, statsPerTrack := range *stats {
		track, err := statSvc.trackService.Get(trackID)
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

func (statSvc *StatisticsService) GetLatestStatForTrack(trackID uint64) (*models.Statistics, error) {

	stats, err := statSvc.GetForTrack(trackID)
	if err != nil {
		return nil, err
	}

	latestStatDate := stats[0].Date
	latestStat := stats[0]
	for _, stat := range stats {
		if stat.Date.After(latestStatDate) {
			latestStatDate = stat.Date
			latestStat = stat
		}
	}

	return &latestStat, nil
}
