package data_builders

import (
	"cookdroogers/models"
	"time"
)

type PublicationBuilder struct {
	Publication *models.Publication
}

func NewPublicationBuilder() *PublicationBuilder {
	return &PublicationBuilder{
		Publication: &models.Publication{
			PublicationID: 88,
			Date:          time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC),
			ReleaseID:     888,
			ManagerID:     8,
		},
	}
}

func (b *PublicationBuilder) WithPublicationID(id uint64) *PublicationBuilder {
	b.Publication.PublicationID = id
	return b
}

func (b *PublicationBuilder) WithDate(date time.Time) *PublicationBuilder {
	b.Publication.Date = date
	return b
}

func (b *PublicationBuilder) WithReleaseID(id uint64) *PublicationBuilder {
	b.Publication.ReleaseID = id
	return b
}

func (b *PublicationBuilder) WithManagerID(id uint64) *PublicationBuilder {
	b.Publication.ManagerID = id
	return b
}

func (b *PublicationBuilder) Build() *models.Publication {
	return b.Publication
}