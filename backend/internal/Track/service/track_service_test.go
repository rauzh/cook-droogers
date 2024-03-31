package service

import (
	mocks "cookdroogers/internal/repo/mocks"
	"cookdroogers/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackService_Get(t *testing.T) {

	mockTrackRepo := mocks.NewTrackRepo(t)
	mockTrackRepo.EXPECT().Get(uint64(1234)).Return(&models.Track{
		TrackID:  1234,
		Title:    "aa",
		Duration: 120,
		Genre:    "rock",
		Artists:  []uint64{82, 4},
	}, nil).Once()

	ts := NewTrackService(mockTrackRepo)

	track, err := ts.Get(1234)
	assert.Nil(t, err)
	assert.Equal(t, "rock", track.Genre)
}
