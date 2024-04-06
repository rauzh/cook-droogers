package publish_criteria

import (
	publicationService "cookdroogers/internal/publication/service"
	"cookdroogers/internal/requests/criteria_controller"
	"cookdroogers/internal/requests/publish"
	cdtime "cookdroogers/pkg/time"
)

const (
	ArtistReleaseLimitPerSeason     criteria.CriteriaName = "No releases from artist more than limit"
	LimitPerSeason                                        = 2
	ExplanationArtistReleaseLimit                         = "More than limit releases per season"
	DiffArtistReleaseLimitPerSeason                       = -1
)

type ArtistReleaseLimitPerSeasonCriteria struct {
	req                *publish.PublishRequest
	publicationService publicationService.IPublicationService
}

func (oarpsc *ArtistReleaseLimitPerSeasonCriteria) Name() criteria.CriteriaName {
	return ArtistReleaseLimitPerSeason
}

func (oarpsc *ArtistReleaseLimitPerSeasonCriteria) Apply() (result criteria.CriteriaDiff) {

	pubsFromArtistLastSeason, err := oarpsc.publicationService.GetAllByArtistSinceDate(
		cdtime.RelevantPeriod(), oarpsc.req.ApplierID)

	if err != nil {
		result.Explanation = criteria.ExplanationCantApply
		return
	}

	if len(pubsFromArtistLastSeason) > LimitPerSeason {
		result.Diff = DiffArtistReleaseLimitPerSeason
		result.Explanation = ExplanationArtistReleaseLimit
		return
	}

	result.Explanation = criteria.ExplanationOK

	return
}

type ArtistReleaseLimitPerSeasonCriteriaFabric struct {
	req                *publish.PublishRequest
	publicationService publicationService.IPublicationService
}

func (fabric *ArtistReleaseLimitPerSeasonCriteriaFabric) Create() criteria.Criteria {
	return &ArtistReleaseLimitPerSeasonCriteria{req: fabric.req, publicationService: fabric.publicationService}
}
