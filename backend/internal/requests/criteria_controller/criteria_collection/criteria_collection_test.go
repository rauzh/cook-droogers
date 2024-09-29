package criteria_collection

import (
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/criteria_controller"
	"cookdroogers/internal/requests/criteria_controller/mocks"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type _depFields struct {
	criterias   []*mocks.Criteria
	critFabrics []*mocks.AbstractCriteriaFabric
}

type CriteriaCollectionSuite struct {
	suite.Suite
}

func _newMockCriteriaDepFields(t provider.T) *_depFields {

	mockCrit1 := mocks.NewCriteria(t)
	mockCrit2 := mocks.NewCriteria(t)

	critFabr1 := mocks.NewAbstractCriteriaFabric(t)
	critFabr2 := mocks.NewAbstractCriteriaFabric(t)

	return &_depFields{
		criterias:   []*mocks.Criteria{mockCrit1, mockCrit2},
		critFabrics: []*mocks.AbstractCriteriaFabric{critFabr1, critFabr2},
	}
}

func (s *CriteriaCollectionSuite) TestCriteriaCollection_Apply(t provider.T) {
	t.Title("Apply: All Criterias OK")
	t.Tags("CriteriaCollection")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {

		df := _newMockCriteriaDepFields(t)

		df.critFabrics[0].EXPECT().Create().Return(df.criterias[0], nil).Once()
		df.critFabrics[1].EXPECT().Create().Return(df.criterias[1], nil).Once()

		df.criterias[0].EXPECT().Apply(base.GetBaseRequestObject()).
			Return(criteria.CriteriaDiff{
				Diff:        2,
				Explanation: "Explanation 1",
			}).Once()
		df.criterias[0].EXPECT().Name().Return(criteria.CriteriaName("Criteria1")).Once()

		df.criterias[1].EXPECT().Apply(base.GetBaseRequestObject()).
			Return(criteria.CriteriaDiff{
				Diff:        3,
				Explanation: "Explanation 2",
			}).Once()
		df.criterias[1].EXPECT().Name().Return(criteria.CriteriaName("Criteria2")).Once()

		request := base.GetBaseRequestObject()
		cc, _ := BuildCollection(df.critFabrics[0], df.critFabrics[1])

		result := cc.Apply(request)

		sCtx.Assert().Equal(5, result.ResultDiff)

		sCtx.Assert().Equal(map[criteria.CriteriaName]criteria.CriteriaDiff{
			"Criteria1": {Diff: 2, Explanation: "Explanation 1"},
			"Criteria2": {Diff: 3, Explanation: "Explanation 2"},
		}, result.ResultExplanation)
	})
}

func TestCriteriaCollectionSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(CriteriaCollectionSuite))
}
