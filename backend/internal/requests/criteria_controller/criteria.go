package criteria

import "cookdroogers/internal/requests/base"

const ExplanationCantApply = "Can't apply criteria"
const ExplanationOK = "OK"

type CriteriaDiff struct {
	Diff        int
	Explanation string
}

type CriteriaName string

type Criteria interface {
	Apply(base.IRequest) CriteriaDiff
	Name() CriteriaName
}

type AbstractCriteriaFabric interface {
	Create() (Criteria, error)
}
