package service

import (
	"cookdroogers/internal/Release/repo"
	releaseService "cookdroogers/internal/Release/service"
	trackService "cookdroogers/internal/Track/service"
	"cookdroogers/models"
	"fmt"
	"runtime"
	"sync"
)

type ReleaseService struct {
	trkSvc trackService.ITrackService
	repo   repo.ReleaseRepo
}

func NewReleaseService(
	trkSvc trackService.ITrackService,
	r repo.ReleaseRepo) releaseService.IReleaseService {
	return &ReleaseService{trkSvc: trkSvc, repo: r}
}

// Create new release and its tracks in DB
func (rlsSvc *ReleaseService) Create(release *models.Release, tracks []models.Track) error {

	release.Status = models.UnpublishedRelease

	// Divide tracks on parts by CPUnum and upload then concurrently
	rlsSvc.uploadTracksParallel(release, tracks)

	if err := rlsSvc.repo.Create(release); err != nil {
		return fmt.Errorf("can't create release with err %w", err)
	}

	return nil
}

func (rlsSvc *ReleaseService) uploadTracksParallel(release *models.Release, tracks []models.Track) {

	wg := new(sync.WaitGroup)

	workersNum := runtime.NumCPU()
	tracksLen := len(tracks)
	mu := new(sync.Mutex)

	for worker := 0; worker < workersNum; worker++ {
		start, end := tracksLen/workersNum*worker, tracksLen/workersNum*(worker+1)
		if worker == workersNum-1 {
			end = tracksLen - 1
		}

		wg.Add(1)
		go rlsSvc.uploadBunchOfTracks(release, start, end, tracks, wg, mu)
	}
	wg.Wait()
}

func (rlsSvc *ReleaseService) uploadBunchOfTracks(release *models.Release,
	start, end int, tracks []models.Track,
	wg *sync.WaitGroup, mu *sync.Mutex) {

	defer wg.Done()

	for i := start; start < end; i++ {
		trackID, err := rlsSvc.trkSvc.Create(&tracks[i])
		if err != nil {
			continue
		}

		mu.Lock()
		release.Tracks = append(release.Tracks, trackID)
		mu.Unlock()

	}
}

func (rlsSvc *ReleaseService) Get(releaseID uint64) (*models.Release, error) {
	release, err := rlsSvc.repo.Get(releaseID)
	if err != nil {
		return nil, fmt.Errorf("can't get release with err %w", err)
	}
	return release, nil
}

func (rlsSvc *ReleaseService) GetAllByArtist(artistID uint64) ([]models.Release, error) {
	releases, err := rlsSvc.repo.GetAllByArtist(artistID)
	if err != nil {
		return nil, fmt.Errorf("can't get release with err %w", err)
	}
	return releases, nil
}

func (rlsSvc *ReleaseService) GetAllTracks(release *models.Release) ([]models.Track, error) {
	tracks, err := rlsSvc.repo.GetAllTracks(release)
	if err != nil {
		return nil, fmt.Errorf("can't get release with err %w", err)
	}
	return tracks, nil
}

func (rlsSvc *ReleaseService) Update(release *models.Release) error {
	if err := rlsSvc.repo.Update(release); err != nil {
		return fmt.Errorf("can't update release with err %w", err)
	}
	return nil
}

func (rlsSvc *ReleaseService) UpdateStatus(id uint64, stat models.ReleaseStatus) error {
	if err := rlsSvc.repo.UpdateStatus(id, stat); err != nil {
		return fmt.Errorf("can't update release with err %w", err)
	}
	return nil
}

func (rlsSvc *ReleaseService) GetMainGenre(releaseID uint64) (string, error) {
	release, err := rlsSvc.repo.Get(releaseID)
	if err != nil {
		return "", fmt.Errorf("can't get release with err %w", err)
	}

	genres := make(map[string]int)
	for _, trackID := range release.Tracks {
		track, err := rlsSvc.trkSvc.Get(trackID)
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
