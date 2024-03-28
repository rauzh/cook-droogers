package adapters

import (
	"bytes"
	artistRepo "cookdroogers/internal/Artist/repo"
	releaseRepo "cookdroogers/internal/Release/repo"
	"cookdroogers/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type StatFetcherAdapter struct {
	url         string
	artistRepo  artistRepo.ArtistRepo
	releaseRepo releaseRepo.ReleaseRepo
}

type SendData struct {
	Artist string `json:"artist"`
	Track  string `json:"track"`
}

type StatJSON struct {
	Date    time.Time `json:"date"`
	Streams uint64    `json:"streams"`
	Likes   uint64    `json:"likes"`
	Track   uint64    `json:"track"`
	Artist  uint64    `json:"artist"`
}

func (fetcher *StatFetcherAdapter) Fetch(tracks []models.Track) ([]models.Statistics, error) {

	tracksInfo := make([]SendData, len(tracks))

	for _, track := range tracks {

		artist, err := fetcher.artistRepo.Get(track.Artists[0])
		if err != nil {
			return nil, err
		}

		tracksInfo = append(tracksInfo, SendData{
			Artist: artist.Nickname,
			Track:  track.Title,
		})
	}

	tracksInfoJSON, err := json.Marshal(tracksInfo)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fetcher.url, "application/json", bytes.NewBuffer(tracksInfoJSON))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("can't fetch stats: non-200 status code")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	statsExternal := []StatJSON{}
	err = json.Unmarshal(body, &statsExternal)
	if err != nil {
		return nil, err
	}

	statsInternal := make([]models.Statistics, len(statsExternal))
	for idx, stat := range statsExternal {
		statsInternal = append(statsInternal, models.Statistics{
			Date:    stat.Date,
			Streams: stat.Streams,
			Likes:   stat.Likes,
			TrackID: tracks[idx].TrackID,
		})
	}

	return statsInternal, nil
}
