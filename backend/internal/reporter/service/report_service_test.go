package service

import (
	artService "cookdroogers/internal/artist/service"
	mngService "cookdroogers/internal/manager/service"
	pbcService "cookdroogers/internal/publication/service"
	rlsService "cookdroogers/internal/release/service"
	mocks "cookdroogers/internal/repo/mocks"
	statFetcher "cookdroogers/internal/statistics/fetcher/mocks"
	statService "cookdroogers/internal/statistics/service"
	trackService "cookdroogers/internal/track/service"
	transacMock "cookdroogers/internal/transactor/mocks"
	"cookdroogers/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReportServiceJSON_GetReportForArtist(t *testing.T) {

	transactionMock := transacMock.NewTransactor(t)
	mockMngRepo := mocks.NewManagerRepo(t)
	mockArtRepo := mocks.NewArtistRepo(t)
	pbcMockRepo := mocks.NewPublicationRepo(t)
	rlsMockRepo := mocks.NewReleaseRepo(t)
	trkMockRepo := mocks.NewTrackRepo(t)
	statMockRepo := mocks.NewStatisticsRepo(t)

	now := time.Now().UTC()

	release1 := models.Release{
		ReleaseID:    1,
		Title:        "album 1",
		Status:       models.PublishedRelease,
		DateCreation: now.AddDate(-1, 0, 0),
		Tracks:       []uint64{11, 12, 13},
		ArtistID:     uint64(777),
	}

	release2 := models.Release{
		ReleaseID:    2,
		Title:        "album 2",
		Status:       models.PublishedRelease,
		DateCreation: now.AddDate(-1, -5, 0),
		Tracks:       []uint64{21, 22},
		ArtistID:     uint64(777),
	}

	rlsMockRepo.EXPECT().GetAllByArtist(uint64(777)).Return(
		[]models.Release{
			release1,
			release2,
		},
		nil).Once()

	rlsMockRepo.EXPECT().GetAllTracks(&release1).Return(
		[]models.Track{
			{
				TrackID:  11,
				Title:    "track 11",
				Duration: 120,
				Genre:    "genre1",
				Type:     "song",
				Artists:  []uint64{777, 128},
			},
			{
				TrackID:  12,
				Title:    "track 12",
				Duration: 120,
				Genre:    "genre1",
				Type:     "song",
				Artists:  []uint64{777},
			},
			{
				TrackID:  13,
				Title:    "track 13",
				Duration: 140,
				Genre:    "genre1",
				Type:     "song",
				Artists:  []uint64{777},
			},
		}, nil).Once()

	rlsMockRepo.EXPECT().GetAllTracks(&release2).Return(
		[]models.Track{
			{
				TrackID:  21,
				Title:    "track 21",
				Duration: 120,
				Genre:    "genre3",
				Type:     "song",
				Artists:  []uint64{777},
			},
			{
				TrackID:  22,
				Title:    "track 22",
				Duration: 150,
				Genre:    "genre3",
				Type:     "song",
				Artists:  []uint64{777},
			},
		}, nil).Once()

	statMockRepo.EXPECT().GetForTrack(uint64(11)).Return(
		[]models.Statistics{
			{
				StatID:  111,
				Date:    now.AddDate(-1, 2, 0),
				Streams: 100,
				Likes:   5,
				TrackID: 11,
			},
			{
				StatID:  112,
				Date:    now.AddDate(-1, 5, 0),
				Streams: 230,
				Likes:   5,
				TrackID: 11,
			},
		}, nil).Once()
	statMockRepo.EXPECT().GetForTrack(uint64(12)).Return(
		[]models.Statistics{
			{
				StatID:  121,
				Date:    now.AddDate(-1, 2, 0),
				Streams: 1010,
				Likes:   5,
				TrackID: 12,
			},
			{
				StatID:  122,
				Date:    now.AddDate(-1, 6, 0),
				Streams: 1230,
				Likes:   5,
				TrackID: 12,
			},
		}, nil).Once()
	statMockRepo.EXPECT().GetForTrack(uint64(13)).Return(
		[]models.Statistics{
			{
				StatID:  131,
				Date:    now.AddDate(-1, 2, 0),
				Streams: 11010,
				Likes:   5,
				TrackID: 13,
			},
			{
				StatID:  132,
				Date:    now.AddDate(-1, 6, 0),
				Streams: 11230,
				Likes:   5,
				TrackID: 13,
			},
		}, nil).Once()
	statMockRepo.EXPECT().GetForTrack(uint64(21)).Return(
		[]models.Statistics{
			{
				StatID:  211,
				Date:    now.AddDate(-1, 2, 0),
				Streams: 0,
				Likes:   5,
				TrackID: 21,
			},
			{
				StatID:  212,
				Date:    now.AddDate(0, -3, 0),
				Streams: 30,
				Likes:   5,
				TrackID: 21,
			},
		}, nil).Once()
	statMockRepo.EXPECT().GetForTrack(uint64(22)).Return(
		[]models.Statistics{
			{
				StatID:  221,
				Date:    now.AddDate(-1, 2, 0),
				Streams: 20,
				Likes:   5,
				TrackID: 22,
			},
			{
				StatID:  222,
				Date:    now.AddDate(0, -3, 0),
				Streams: 300,
				Likes:   5,
				TrackID: 22,
			},
		}, nil).Once()

	statMockFetcher := statFetcher.NewStatFetcher(t)

	trkSvc := trackService.NewTrackService(trkMockRepo)
	mngSvc := mngService.NewManagerService(mockMngRepo)
	artSvc := artService.NewArtistService(mockArtRepo)
	pbcSvc := pbcService.NewPublicationService(pbcMockRepo)
	rlsSvc := rlsService.NewReleaseService(trkSvc, transactionMock, rlsMockRepo)
	statSvc := statService.NewStatisticsService(trkSvc, statMockFetcher, statMockRepo, rlsSvc)

	rptSvc := NewReportService(mngSvc, statSvc, artSvc, pbcSvc, rlsSvc)

	report, err := rptSvc.GetReportForArtist(777)

	reportMap := make(map[string]map[string]uint64)
	for releaseName, releaseStatsJSON := range report {
		releaseStats := make(map[string]uint64)
		json.Unmarshal(releaseStatsJSON, &releaseStats)
		reportMap[releaseName] = releaseStats
	}

	t.Log(reportMap)
	assert.Nil(t, err)
	assert.Equal(t, map[string]map[string]uint64{
		"album 1": map[string]uint64{
			"track 11": 230,
			"track 12": 1230,
			"track 13": 11230,
		},
		"album 2": map[string]uint64{
			"track 21": 30,
			"track 22": 300,
		},
	}, reportMap)
}
