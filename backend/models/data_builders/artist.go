package data_builders

import (
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"time"
)

type ArtistBuilder struct {
	artist *models.Artist
}

func NewArtistBuilder() *ArtistBuilder {
	return &ArtistBuilder{
		artist: &models.Artist{
			ArtistID:     1,
			UserID:       7,
			Nickname:     "uzi",
			ContractTerm: cdtime.GetEndOfContract(),
			Activity:     true,
			ManagerID:    9,
		},
	}
}

func (b *ArtistBuilder) WithID(id uint64) *ArtistBuilder {
	b.artist.ArtistID = id
	return b
}

func (b *ArtistBuilder) WithNickname(nickname string) *ArtistBuilder {
	b.artist.Nickname = nickname
	return b
}

func (b *ArtistBuilder) WithUserID(userid uint64) *ArtistBuilder {
	b.artist.UserID = userid
	return b
}

func (b *ArtistBuilder) WithContractTerm(contractTerm time.Time) *ArtistBuilder {
	b.artist.ContractTerm = contractTerm
	return b
}

func (b *ArtistBuilder) WithActivity(activity bool) *ArtistBuilder {
	b.artist.Activity = activity
	return b
}

func (b *ArtistBuilder) WithManagerID(managerid uint64) *ArtistBuilder {
	b.artist.ManagerID = managerid
	return b
}

func (b *ArtistBuilder) Build() *models.Artist {
	return b.artist
}
