package data_builders

import "cookdroogers/models"

type TrackBuilder struct {
	Track *models.Track
}

func NewTrackBuilder() *TrackBuilder {
	return &TrackBuilder{
		Track: &models.Track{
			TrackID:  1111,
			Title:    "title",
			Duration: 60,
			Genre:    "trap",
			Type:     "intro",
			Artists:  nil,
		},
	}
}

func (b *TrackBuilder) WithID(id uint64) *TrackBuilder {
	b.Track.TrackID = id
	return b
}

func (b *TrackBuilder) WithDuration(d uint64) *TrackBuilder {
	b.Track.Duration = d
	return b
}

func (b *TrackBuilder) WithTitle(t string) *TrackBuilder {
	b.Track.Title = t
	return b
}

func (b *TrackBuilder) WithGenre(g string) *TrackBuilder {
	b.Track.Genre = g
	return b
}

func (b *TrackBuilder) WithType(t string) *TrackBuilder {
	b.Track.Type = t
	return b
}

func (b *TrackBuilder) WithArtists(artists []uint64) *TrackBuilder {
	b.Track.Artists = artists
	return b
}

func (b *TrackBuilder) Build() *models.Track {
	return b.Track
}
