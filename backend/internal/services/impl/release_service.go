package service

import (
	"cookdroogers/internal/models"
	"cookdroogers/internal/repo"
	service "cookdroogers/internal/services"
	"fmt"
	"runtime"
	"sync"
)

type ReleaseService struct {
	trackService service.ITrackService
	repo         repo.ReleaseRepo
}

func (rs *ReleaseService) Create(release *models.Release, tracks []models.Track) error {

	release.Status = models.UnpublishedRelease

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
		go func(start, end int, mu *sync.Mutex) {
			defer wg.Done()

			for i := start; start < end; i++ {
				trackID, err := rs.trackService.Create(&tracks[i])
				if err != nil {
					continue
				}

				mu.Lock()
				release.Tracks = append(release.Tracks, trackID)
				mu.Unlock()
			}
		}(start, end, mu)
	}
	wg.Wait()

	if err := rs.repo.Create(release); err != nil {
		return fmt.Errorf("can't create release with err %w", err)
	}
	return nil
}

func (rs *ReleaseService) Get(releaseID uint64) (*models.Release, error) {
	release, err := rs.repo.Get(releaseID)
	if err != nil {
		return nil, fmt.Errorf("can't get release with err %w", err)
	}
	return release, nil
}

func (rs *ReleaseService) Update(release *models.Release) error {
	if err := rs.repo.Update(release); err != nil {
		return fmt.Errorf("can't update release with err %w", err)
	}
	return nil
}

func (rs *ReleaseService) GetMainGenre(releaseID uint64) (string, error) {
	release, err := rs.repo.Get(releaseID)
	if err != nil {
		return "", fmt.Errorf("can't get release with err %w", err)
	}

	genres := make(map[string]int)
	for _, trackID := range release.Tracks {
		track, err := rs.trackService.Get(trackID)
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
