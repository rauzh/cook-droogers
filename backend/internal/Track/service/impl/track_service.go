package service

import (
	trackErrors "cookdroogers/internal/Track/errors"
	"cookdroogers/internal/Track/repo"
	trackService "cookdroogers/internal/Track/service"
	"cookdroogers/models"
	"fmt"
)

type TrackService struct {
	repo repo.TrackRepo
}

func NewTrackService(r repo.TrackRepo) trackService.ITrackService {
	return &TrackService{repo: r}
}

func (trkSvc *TrackService) validate(track *models.Track) error {
	if track.Genre == "" {
		return trackErrors.ErrNoGenre
	}

	if len(track.Artists) < 1 {
		return trackErrors.ErrNoArtist
	}

	if track.Type == "" {
		return trackErrors.ErrNoType
	}

	if track.Title == "" {
		return trackErrors.ErrNoTitle
	}

	return nil
}

func (trkSvc *TrackService) Create(track *models.Track) (uint64, error) {

	if err := trkSvc.validate(track); err != nil {
		return 0, err
	}

	trackID, err := trkSvc.repo.Create(track)
	if err != nil {
		return 0, fmt.Errorf("can't create track with err %w", err)
	}
	return trackID, nil
}

func (trkSvc *TrackService) Get(trackID uint64) (*models.Track, error) {
	track, err := trkSvc.repo.Get(trackID)
	if err != nil {
		return nil, fmt.Errorf("can't get track with err %w", err)
	}
	return track, nil
}
