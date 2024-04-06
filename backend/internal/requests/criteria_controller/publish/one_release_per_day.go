package publish_criteria

import (
	publicationService "cookdroogers/internal/publication/service"
	"cookdroogers/internal/requests/criteria_controller"
	"cookdroogers/internal/requests/publish"
)

const (
	OneReleasePerDay      criteria.CriteriaName = "No releases that day"
	ExplanationOneRelease                       = "More than one release per day"
	DiffOneRelease                              = -1
)

type OneReleasePerDayCriteria struct {
	req                *publish.PublishRequest
	publicationService publicationService.IPublicationService
}

func (orpdc *OneReleasePerDayCriteria) Name() criteria.CriteriaName {
	return OneReleasePerDay
}

func (orpdc *OneReleasePerDayCriteria) Apply() (result criteria.CriteriaDiff) {

	pubsThatDay, err := orpdc.publicationService.GetAllByDate(orpdc.req.Date)
	if err != nil {
		result.Explanation = criteria.ExplanationCantApply
		return
	}

	if len(pubsThatDay) > 1 {
		result.Diff = DiffOneRelease
		result.Explanation = ExplanationOneRelease
		return
	}

	result.Explanation = criteria.ExplanationOK

	return
}

type OneReleasePerDayCriteriaFabric struct {
	req                *publish.PublishRequest
	publicationService publicationService.IPublicationService
}

func (fabric *OneReleasePerDayCriteriaFabric) Create() criteria.Criteria {
	return &OneReleasePerDayCriteria{req: fabric.req, publicationService: fabric.publicationService}
}
