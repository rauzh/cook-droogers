package data_builders

import (
	"cookdroogers/models"
)

type ManagerBuilder struct {
	Manager *models.Manager
}

func NewManagerBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager: &models.Manager{
			ManagerID: 8,
			UserID:    88,
			Artists:   nil,
		},
	}
}

func (b *ManagerBuilder) WithManagerID(managerid uint64) *ManagerBuilder {
	b.Manager.ManagerID = managerid
	return b
}

func (b *ManagerBuilder) WithUserID(id uint64) *ManagerBuilder {
	b.Manager.UserID = id
	return b
}

func (b *ManagerBuilder) WithArtists(artists []uint64) *ManagerBuilder {
	b.Manager.Artists = artists
	return b
}

func (b *ManagerBuilder) Build() *models.Manager {
	return b.Manager
}
