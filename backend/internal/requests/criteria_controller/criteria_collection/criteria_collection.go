package criteria_collection

import (
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/criteria_controller"
	"fmt"
)

type CriteriaCollectionDiff struct {
	ResultDiff        int
	ResultExplanation map[criteria.CriteriaName]criteria.CriteriaDiff
}

type ICriteriaCollection interface {
	Apply(request base.IRequest) CriteriaCollectionDiff
}

type CriteriaCollection struct {
	criterias []criteria.Criteria
}

func (cc *CriteriaCollection) Apply(request base.IRequest) (result CriteriaCollectionDiff) {

	result.ResultExplanation = make(map[criteria.CriteriaName]criteria.CriteriaDiff)

	for _, crit := range cc.criterias {

		critRes := crit.Apply(request)

		result.ResultDiff += critRes.Diff
		result.ResultExplanation[crit.Name()] = critRes
	}

	return
}

func BuildCollection(fabrics ...criteria.AbstractCriteriaFabric) (ICriteriaCollection, error) {

	crits := make([]criteria.Criteria, 0)
	for _, fabric := range fabrics {
		crit, err := fabric.Create()
		if err != nil {
			return nil, err
		}
		crits = append(crits, crit)
	}

	return &CriteriaCollection{criterias: crits}, nil
}

func DiffToString(name criteria.CriteriaName, explanation string, diff int) string {
	return fmt.Sprintf("**%s** diff: %d\n**%s** reason: %s", name, diff, name, explanation)
}
