package criteria

type CriteriaCollectionDiff struct {
	ResultDiff        int
	ResultExplanation map[CriteriaName]CriteriaDiff
}

type ICriteriaCollection interface {
	Apply() CriteriaCollectionDiff
}

type CriteriaCollection struct {
	criterias []Criteria
}

func (cc *CriteriaCollection) Apply() (result CriteriaCollectionDiff) {

	result.ResultExplanation = make(map[CriteriaName]CriteriaDiff)

	for _, crit := range cc.criterias {

		critRes := crit.Apply()

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
