package service

import (
	"cookdroogers/internal/Manager/repo"
	s "cookdroogers/internal/Manager/service"
	ss "cookdroogers/internal/Statistics/service"
	"cookdroogers/models"
	"fmt"
)

type ManagerService struct {
	ss   ss.IStatisticsService
	repo repo.ManagerRepo
}

func NewManagerService(
	r repo.ManagerRepo,
	ss ss.IStatisticsService) s.IManagerService {
	return &ManagerService{ss: ss, repo: r}
}

func (ms *ManagerService) Create(artist *models.Manager) error {
	if err := ms.repo.Create(artist); err != nil {
		return fmt.Errorf("can't create manager with err %w", err)
	}
	return nil
}

func (ms *ManagerService) Get(id uint64) (*models.Manager, error) {
	manager, err := ms.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("can't get manager with err %w", err)
	}
	return manager, nil
}

func (ms *ManagerService) GetRandomManagerID() (uint64, error) {
	id, err := ms.repo.GetRandManagerID()
	if err != nil {
		return 0, fmt.Errorf("can't get manager with err %w", err)
	}
	return id, nil
}

func (ms *ManagerService) GetReport(managerID uint64) (map[string]string, error) {

	// manager, err := ms.repo.Get(managerID)
	// if err != nil {
	// 	return nil, err
	// }

	// report := make(map[string]string)

	// relevantGenre, err := ms.ss.GetRelevantGenre()
	// if err != nil {
	// 	return nil, err
	// }
	// report["relevantGenre"] = relevantGenre

	// for _, artistID := range manager.Artists {

	// }

	return nil, nil
}
