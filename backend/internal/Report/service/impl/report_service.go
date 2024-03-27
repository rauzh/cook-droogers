package impl

import (
	artistService "cookdroogers/internal/Artist/service"
	managerService "cookdroogers/internal/Manager/service"
	publicationService "cookdroogers/internal/Publication/service"
	releaseService "cookdroogers/internal/Release/service"
	reportService "cookdroogers/internal/Report/service"
	statisticsServive "cookdroogers/internal/Statistics/service"
	"cookdroogers/models"
	"encoding/json"
	"time"
)

type ReportService struct {
	mngSvc  managerService.IManagerService
	statSvc statisticsServive.IStatisticsService
	artSvc  artistService.IArtistService
	pbcSvc  publicationService.IPublicationService
	rlsSvc  releaseService.IReleaseService
}

func NewReportService(
	mngSvc managerService.IManagerService,
	statSvc statisticsServive.IStatisticsService,
	artSvc artistService.IArtistService,
	pbcSvc publicationService.IPublicationService,
	rlsSvc releaseService.IReleaseService,
) reportService.IReportService {
	return &ReportService{
		mngSvc:  mngSvc,
		statSvc: statSvc,
		artSvc:  artSvc,
		pbcSvc:  pbcSvc,
		rlsSvc:  rlsSvc,
	}
}

func (rptSvc *ReportService) GetReportForManager(mngID uint64) (map[string][]byte, error) {

	report := make(map[string][]byte)

	relevantGenre, err := rptSvc.statSvc.GetRelevantGenre()
	if err != nil {
		return nil, nil
	}

	relGenreJson, err := json.Marshal(relevantGenre)
	if err != nil {
		return nil, err
	}

	report["relevant_genre"] = relGenreJson

	manager, err := rptSvc.mngSvc.Get(mngID)
	if err != nil {
		return nil, err
	}

	artists := make(map[uint64]*models.Artist)

	for _, artistID := range manager.Artists {
		artist, err := rptSvc.artSvc.Get(artistID)
		if err != nil {
			return nil, err
		}
		artists[artistID] = artist
	}

	pubs, err := rptSvc.pbcSvc.GetAllByManager(mngID)
	if err != nil {
		return nil, err
	}

	lastSeasonStatDate := time.Now().UTC().AddDate(0, -3, 0)
	currentDate := time.Now().UTC()

	artistStats := make(map[string]map[string]map[uint64][]models.Statistics)

	for _, pub := range pubs {

		if pub.Date.After(currentDate) {
			continue
		}

		releaseStats := make(map[string]map[uint64][]models.Statistics)

		release, err := rptSvc.rlsSvc.Get(pub.ReleaseID)
		if err != nil {
			return nil, err
		}

		tracks, err := rptSvc.rlsSvc.GetAllTracks(release)
		if err != nil {
			return nil, err
		}

		tracksStats := make(map[uint64][]models.Statistics)
		for _, track := range tracks {
			stats, err := rptSvc.statSvc.GetForTrack(track.TrackID)
			if err != nil {
				return nil, err
			}

			var closestToLastSeasonDate time.Time
			var lastSeasonStat models.Statistics
			latestStatDate := stats[0].Date
			latestStat := stats[0]
			for _, stat := range stats {
				if stat.Date.After(latestStatDate) {
					latestStatDate = stat.Date
					latestStat = stat
				}

				if stat.Date.Before(lastSeasonStatDate) && stat.Date.After(closestToLastSeasonDate) {
					closestToLastSeasonDate = stat.Date
					lastSeasonStat = stat
				}
			}

			tracksStats[track.TrackID] = []models.Statistics{lastSeasonStat, latestStat}
		}
		releaseStats[release.Title] = tracksStats
		artistStats[artists[release.ArtistID].Nickname] = releaseStats
	}

	artistStatsJson, err := json.Marshal(artistStats)
	if err != nil {
		return nil, err
	}
	report["artists_stats"] = artistStatsJson

	return report, nil
}

func (rptSvc *ReportService) GetReportForArtist(artistID uint64) (map[string][]byte, error) {

	report := make(map[string][]byte)

	releases, err := rptSvc.rlsSvc.GetAllByArtist(artistID)
	if err != nil {
		return nil, err
	}

	for _, release := range releases {
		tracks, err := rptSvc.rlsSvc.GetAllTracks(&release)
		if err != nil {
			return nil, err
		}

		tracksStats := make(map[string]uint64)
		for _, track := range tracks {
			stats, err := rptSvc.statSvc.GetForTrack(track.TrackID)
			if err != nil {
				return nil, err
			}

			var totalStreams uint64
			for _, stat := range stats {
				totalStreams += stat.Streams
			}

			tracksStats[track.Ttile] = totalStreams
		}
		releaseStatsJson, err := json.Marshal(tracksStats)
		if err != nil {
			return nil, err
		}
		report[release.Title] = releaseStatsJson
	}

	return report, nil
}
