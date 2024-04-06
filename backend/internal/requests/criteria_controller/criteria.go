package criteria

const ExplanationCantApply = "Can't apply criteria"
const ExplanationOK = "OK"

type CriteriaDiff struct {
	Diff        int
	Explanation string
}

type CriteriaName string

type Criteria interface {
	Apply() CriteriaDiff
	Name() CriteriaName
}

type AbstractCriteriaFabric interface {
	Create() (Criteria, error)
}
