package publish_criteria

import (
	releaseService "cookdroogers/internal/release/service"
	criteria "cookdroogers/internal/requests/criteria_controller"
	"cookdroogers/internal/requests/publish"
	statService "cookdroogers/internal/statistics/service"
	"strings"
)

const (
	RelevantGenre            criteria.CriteriaName = "Genre should be relevant"
	ExplanationRelevantGenre                       = "Genre is irrelevant"
	DiffRelevantGenre                              = -1
)

type RelevantGenreCriteria struct {
	req            *publish.PublishRequest
	releaseService releaseService.IReleaseService
	statService    statService.IStatisticsService
}

func (rgc *RelevantGenreCriteria) Name() criteria.CriteriaName {
	return RelevantGenre
}

func (rgc *RelevantGenreCriteria) Apply() (result criteria.CriteriaDiff) {

	releaseGenre, err := rgc.releaseService.GetMainGenre(rgc.req.ReleaseID)
	if err != nil {
		result.Explanation = criteria.ExplanationCantApply
		return
	}

	relevantGenre, err := rgc.statService.GetRelevantGenre()
	if err != nil {
		result.Explanation = criteria.ExplanationCantApply
		return
	}

	if strings.ToLower(releaseGenre) != strings.ToLower(relevantGenre) {
		result.Diff = DiffRelevantGenre
		result.Explanation = ExplanationRelevantGenre
		return
	}

	result.Explanation = criteria.ExplanationOK

	return
}

type RelevantGenreCriteriaFabric struct {
	req            *publish.PublishRequest
	releaseService releaseService.IReleaseService
	statService    statService.IStatisticsService
}

func (fabric *RelevantGenreCriteriaFabric) Create() criteria.Criteria {
	return &RelevantGenreCriteria{req: fabric.req, releaseService: fabric.releaseService, statService: fabric.statService}
}
