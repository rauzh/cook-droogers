package service

import (
	"cookdroogers/internal/Track/repo"
	s "cookdroogers/internal/Track/service"
	"cookdroogers/models"
	"fmt"
)

type TrackService struct {
	repo repo.TrackRepo
}

func NewTrackService(r repo.TrackRepo) s.ITrackService {
	return &TrackService{repo: r}
}

func (ts *TrackService) Create(track *models.Track) (uint64, error) {
	trackID, err := ts.repo.Create(track)
	if err != nil {
		return 0, fmt.Errorf("can't create track with err %w", err)
	}
	return trackID, nil
}

func (ts *TrackService) Get(trackID uint64) (*models.Track, error) {
	track, err := ts.repo.Get(trackID)
	if err != nil {
		return nil, fmt.Errorf("can't get track with err %w", err)
	}
	return track, nil
}
