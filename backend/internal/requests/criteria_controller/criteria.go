package criteria

import "cookdroogers/internal/requests/base"

const ExplanationCantApply = "Can't apply criteria"
const ExplanationOK = "OK"

type CriteriaDiff struct {
	Diff        int
	Explanation string
}

type CriteriaName string

//go:generate mockery --name Criteria --with-expecter
type Criteria interface {
	Apply(base.IRequest) CriteriaDiff
	Name() CriteriaName
}

//go:generate mockery --name AbstractCriteriaFabric --with-expecter
type AbstractCriteriaFabric interface {
	Create() (Criteria, error)
}
