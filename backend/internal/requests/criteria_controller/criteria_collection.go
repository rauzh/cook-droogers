package criteria

import (
	"cookdroogers/internal/requests/base"
	"fmt"
)

type CriteriaCollectionDiff struct {
	ResultDiff        int
	ResultExplanation map[CriteriaName]CriteriaDiff
}

type ICriteriaCollection interface {
	Apply(request base.IRequest) CriteriaCollectionDiff
}

type CriteriaCollection struct {
	criterias []Criteria
}

func (cc *CriteriaCollection) Apply(request base.IRequest) (result CriteriaCollectionDiff) {

	result.ResultExplanation = make(map[CriteriaName]CriteriaDiff)

	for _, crit := range cc.criterias {

		critRes := crit.Apply(request)

		result.ResultDiff += critRes.Diff
		result.ResultExplanation[crit.Name()] = critRes
	}

	return
}

func BuildCollection(fabrics ...AbstractCriteriaFabric) (ICriteriaCollection, error) {

	crits := make([]Criteria, len(fabrics))
	for _, fabric := range fabrics {
		crit, err := fabric.Create()
		if err != nil {
			return nil, err
		}
		crits = append(crits, crit)
	}

	return &CriteriaCollection{criterias: crits}, nil
}

func DiffToString(name CriteriaName, explanation string, diff int) string {
	return fmt.Sprintf("**%s** diff: %d\n**%s** reason: %s", name, diff, name, explanation)
}
