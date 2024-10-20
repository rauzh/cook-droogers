package service

import (
	"context"
	releaseErrors "cookdroogers/internal/release/errors"
	"cookdroogers/internal/repo"
	trackService "cookdroogers/internal/track/service"
	"cookdroogers/internal/transactor"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

type IReleaseService interface {
	Create(ctx context.Context, release *models.Release, tracks []*models.Track) error
	Get(ctx context.Context, releaseID uint64) (*models.Release, error)
	GetMainGenre(ctx context.Context, releaseID uint64) (string, error)
	UpdateStatus(context.Context, uint64, models.ReleaseStatus) error
	GetAllByArtist(context.Context, uint64) ([]models.Release, error)
	GetAllTracks(ctx context.Context, release *models.Release) ([]models.Track, error)
}

type ReleaseService struct {
	trkSvc     trackService.ITrackService
	repo       repo.ReleaseRepo
	transactor transactor.Transactor
	logger     *slog.Logger
}

func NewReleaseService(
	trkSvc trackService.ITrackService,
	transactor transactor.Transactor,
	r repo.ReleaseRepo,
	logger *slog.Logger) IReleaseService {
	return &ReleaseService{trkSvc: trkSvc, repo: r, transactor: transactor, logger: logger}
}

func (rlsSvc *ReleaseService) validate(release *models.Release) error {

	if release.Title == "" {
		return releaseErrors.ErrNoTitle
	}

	if release.DateCreation.IsZero() {
		return releaseErrors.ErrNoDate
	}

	return nil
}

func (rlsSvc *ReleaseService) Create(ctx context.Context, release *models.Release, tracks []*models.Track) error {

	if err := rlsSvc.validate(release); err != nil {
		return err
	}

	release.Status = models.UnpublishedRelease

	return rlsSvc.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		if err := rlsSvc.uploadTracks(ctx, release, tracks); err != nil {
			return fmt.Errorf("can't create release with err %w", err)
		}

		if err := rlsSvc.repo.Create(ctx, release); err != nil {
			return fmt.Errorf("can't create release with err %w", err)
		}

		return nil
	})
}

func (rlsSvc *ReleaseService) uploadTracks(ctx context.Context, release *models.Release, tracks []*models.Track) error {

	for _, track := range tracks {
		trackID, err := rlsSvc.trkSvc.Create(ctx, track)
		if err != nil {
			return err
		}

		release.Tracks = append(release.Tracks, trackID)
	}

	return nil
}

func (rlsSvc *ReleaseService) Get(ctx context.Context, releaseID uint64) (*models.Release, error) {
	release, err := rlsSvc.repo.Get(ctx, releaseID)
	if err != nil && strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
		return nil, userErrors.ErrNoUser
	}
	if err != nil {
		return nil, fmt.Errorf("can't get release with err %w", err)
	}
	return release, nil
}

func (rlsSvc *ReleaseService) GetAllByArtist(ctx context.Context, artistID uint64) ([]models.Release, error) {
	releases, err := rlsSvc.repo.GetAllByArtist(ctx, artistID)

	if err != nil {
		return nil, fmt.Errorf("can't get release with err %w", err)
	}
	return releases, nil
}

func (rlsSvc *ReleaseService) GetAllTracks(ctx context.Context, release *models.Release) ([]models.Track, error) {
	tracks, err := rlsSvc.repo.GetAllTracks(ctx, release)

	if err != nil {
		return nil, fmt.Errorf("can't get release with err %w", err)
	}
	return tracks, nil
}

func (rlsSvc *ReleaseService) Update(ctx context.Context, release *models.Release) error {
	if err := rlsSvc.repo.Update(ctx, release); err != nil {
		return fmt.Errorf("can't update release with err %w", err)
	}
	return nil
}

func (rlsSvc *ReleaseService) UpdateStatus(ctx context.Context, id uint64, stat models.ReleaseStatus) error {
	if err := rlsSvc.repo.UpdateStatus(ctx, id, stat); err != nil {
		return fmt.Errorf("can't update release with err %w", err)
	}
	return nil
}

func (rlsSvc *ReleaseService) GetMainGenre(ctx context.Context, releaseID uint64) (string, error) {
	release, err := rlsSvc.repo.Get(ctx, releaseID)
	if err != nil {
		return "", fmt.Errorf("can't get release with err %w", err)
	}

	genres := make(map[string]int)
	for _, trackID := range release.Tracks {
		track, err := rlsSvc.trkSvc.Get(ctx, trackID)
		if err != nil {
			return "", fmt.Errorf("can't get track %d with err %w", trackID, err)
		}

		genres[track.Genre]++
	}

	var maxAmount int
	var mainGenre string
	for genre, amount := range genres {
		if amount > maxAmount {
			maxAmount = amount
			mainGenre = genre
		}
	}

	return mainGenre, nil
}

//
//func (rlsSvc *ReleaseService) uploadTracks(
//	release *models.Release, tracks []models.Track) {
//
//	wg := new(sync.WaitGroup)
//
//	workersNum := runtime.NumCPU()
//	tracksLen := len(tracks)
//	mu := new(sync.Mutex)
//
//	for worker := 0; worker < workersNum; worker++ {
//		start, end := tracksLen/workersNum*worker, tracksLen/workersNum*(worker+1)
//		if worker == workersNum-1 {
//			end = tracksLen - 1
//		}
//
//		wg.Add(1)
//		go rlsSvc.uploadBunchOfTracks(release, start, end, tracks, wg, mu)
//	}
//	wg.Wait()
//}
//
//func (rlsSvc *ReleaseService) uploadBunchOfTracks(release *models.Release,
//	start, end int, tracks []models.Track,
//	wg *sync.WaitGroup, mu *sync.Mutex) {
//
//	defer wg.Done()
//
//	for i := start; start < end; i++ {
//		trackID, err := rlsSvc.trkSvc.Create(&tracks[i])
//		if err != nil {
//			return
//		}
//
//		mu.Lock()
//		release.Tracks = append(release.Tracks, trackID)
//		mu.Unlock()
//
//	}
//}
