package service

import (
	"context"
	releaseService "cookdroogers/internal/release/service"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/statistics/fetcher"
	ts "cookdroogers/internal/track/service"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"errors"
	"fmt"
)

type IStatisticsService interface {
	Create(*models.Statistics) error
	FetchByRelease(release *models.Release) error
	GetForTrack(uint64) ([]models.Statistics, error)
	GetByID(uint64) (*models.Statistics, error)
	GetRelevantGenre() (string, error)
	GetLatestStatForTrack(trackID uint64) (*models.Statistics, error)
}

type StatisticsService struct {
	trackService   ts.ITrackService
	releaseService releaseService.IReleaseService
	fetcher        fetcher.StatFetcher
	repo           repo.StatisticsRepo
}

func NewStatisticsService(
	ts ts.ITrackService,
	f fetcher.StatFetcher,
	r repo.StatisticsRepo,
	rls releaseService.IReleaseService) IStatisticsService {
	return &StatisticsService{
		trackService:   ts,
		releaseService: rls,
		fetcher:        f,
		repo:           r,
	}
}

func (statSvc *StatisticsService) Create(stat *models.Statistics) error {

	stat.Date = cdtime.GetToday()

	if err := statSvc.repo.Create(context.Background(), stat); err != nil {
		return fmt.Errorf("can't create stats with err %w", err)
	}
	return nil
}

func (statSvc *StatisticsService) GetByID(statID uint64) (*models.Statistics, error) {
	stat, err := statSvc.repo.GetByID(context.Background(), statID)

	if err != nil {
		return nil, fmt.Errorf("can't get stats with err %w", err)
	}
	return stat, nil
}

func (statSvc *StatisticsService) GetForTrack(trackID uint64) ([]models.Statistics, error) {
	stats, err := statSvc.repo.GetForTrack(context.Background(), trackID)

	if err != nil {
		return nil, fmt.Errorf("can't get stats for track %d with err %w", trackID, err)
	}
	return stats, nil
}

func (statSvc *StatisticsService) FetchByRelease(release *models.Release) error {

	tracks, err := statSvc.releaseService.GetAllTracks(release)
	if err != nil {
		return fmt.Errorf("can't fetch stats with err %w", err)
	}

	stats, err := statSvc.fetcher.Fetch(tracks)
	if err != nil {
		return fmt.Errorf("can't fetch stats with err %w", err)
	}

	if len(stats) < 1 {
		return errors.New("no stats to fetch")
	}

	for _, stat := range stats {
		stat.Date = cdtime.GetToday()
	}

	if err = statSvc.repo.CreateMany(context.Background(), stats); err != nil {
		return fmt.Errorf("can't create stats with err %w", err)
	}

	return nil
}

func (statSvc *StatisticsService) GetRelevantGenre() (string, error) {

	stats, err := statSvc.repo.GetAllGroupByTracksSince(context.Background(), cdtime.RelevantPeriod())
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
