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
			Artists:  []uint64{7},
		},
	}
}

func (b *TrackBuilder) Build() *models.Track {
	return b.Track
}
