package service

import (
	"context"
	"cookdroogers/internal/repo"
	trackErrors "cookdroogers/internal/track/errors"
	"cookdroogers/models"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

type ITrackService interface {
	Create(context.Context, *models.Track) (uint64, error)
	Get(context.Context, uint64) (*models.Track, error)
}

type TrackService struct {
	repo repo.TrackRepo

	logger *slog.Logger
}

func NewTrackService(r repo.TrackRepo, logger *slog.Logger) ITrackService {
	return &TrackService{repo: r, logger: logger}
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

func (trkSvc *TrackService) Create(ctx context.Context, track *models.Track) (uint64, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if err := trkSvc.validate(track); err != nil {
		return 0, err
	}

	trackID, err := trkSvc.repo.Create(ctx, track)

	if err != nil {
		return 0, fmt.Errorf("can't create track with err %w", err)
	}
	return trackID, nil
}

func (trkSvc *TrackService) Get(ctx context.Context, trackID uint64) (*models.Track, error) {
	track, err := trkSvc.repo.Get(ctx, trackID)
	if err != nil && strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		return nil, trackErrors.ErrNoTrack
	}
	if err != nil {
		return nil, fmt.Errorf("can't get track with err %w", err)
	}
	return track, nil
}
