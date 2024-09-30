package data_builders

import (
	"cookdroogers/models"
	"time"
)

var date = time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC)

type ReleaseBuilder struct {
	Release *models.Release
}

func NewReleaseBuilder() *ReleaseBuilder {
	return &ReleaseBuilder{
		Release: &models.Release{
			ReleaseID:    888,
			Title:        "title",
			Status:       models.PublishedRelease,
			DateCreation: date,
			Tracks:       []uint64{1111},
			ArtistID:     7,
		},
	}
}

func (b *ReleaseBuilder) WithReleaseID(id uint64) *ReleaseBuilder {
	b.Release.ReleaseID = id
	return b
}

func (b *ReleaseBuilder) WithTitle(title string) *ReleaseBuilder {
	b.Release.Title = title
	return b
}

func (b *ReleaseBuilder) WithStatus(status models.ReleaseStatus) *ReleaseBuilder {
	b.Release.Status = status
	return b
}

func (b *ReleaseBuilder) WithDate(date time.Time) *ReleaseBuilder {
	b.Release.DateCreation = date
	return b
}

func (b *ReleaseBuilder) WithTracks(tracks []uint64) *ReleaseBuilder {
	b.Release.Tracks = tracks
	return b
}

func (b *ReleaseBuilder) WithArtistID(id uint64) *ReleaseBuilder {
	b.Release.ArtistID = id
	return b
}

func (b *ReleaseBuilder) Build() *models.Release {
	return b.Release
}
