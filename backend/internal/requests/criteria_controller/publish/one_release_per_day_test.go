package publish_criteria

import (
	"cookdroogers/internal/repo/mocks"
	criteria2 "cookdroogers/internal/requests/criteria_controller"
	"cookdroogers/internal/requests/publish"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"testing"
)

type _depFieldsORPD struct {
	publicationRepo *mocks.PublicationRepo
}

// TEST_HW: default test
type OneReleasePerDayCriteriaSuite struct {
	suite.Suite
}

func _newMockOneReleasePerDayDepFields(t provider.T) *_depFieldsORPD {
	return &_depFieldsORPD{
		publicationRepo: mocks.NewPublicationRepo(t),
	}
}

func (s *OneReleasePerDayCriteriaSuite) TestApply_LimitExceeded(t provider.T) {
	t.Title("Apply: More than one release on that day")
	t.Tags("OneReleasePerDayCriteria")
	t.Parallel()
	t.WithNewStep("Limit exceeded", func(sCtx provider.StepCtx) {

		df := _newMockOneReleasePerDayDepFields(t)

		// Мокируем получение публикаций за тот день
		df.publicationRepo.EXPECT().GetAllByDate(mock.AnythingOfType("context.backgroundCtx"), mock.Anything).
			Return([]models.Publication{
				{}, // Одна публикация уже есть
				{}, // Еще одна публикация (превышен лимит)
			}, nil).Once()

		pubReq := publish.NewPublishRequest(7, 9, cdtime.GetToday().AddDate(0, 0, 14))

		criteria := OneReleasePerDayCriteria{
			publicationRepo: df.publicationRepo,
		}

		result := criteria.Apply(pubReq)

		// Проверяем, что сработало ограничение по количеству публикаций за день
		sCtx.Assert().Equal(DiffOneRelease, result.Diff)
		sCtx.Assert().Equal(ExplanationOneRelease, result.Explanation)
	})
}

func (s *OneReleasePerDayCriteriaSuite) TestApply_NoLimitExceeded(t provider.T) {
	t.Title("Apply: One release or less on that day")
	t.Tags("OneReleasePerDayCriteria")
	t.Parallel()
	t.WithNewStep("Within limit", func(sCtx provider.StepCtx) {

		df := _newMockOneReleasePerDayDepFields(t)

		// Мокируем получение публикаций за тот день
		df.publicationRepo.EXPECT().GetAllByDate(mock.AnythingOfType("context.backgroundCtx"), mock.Anything).
			Return([]models.Publication{
				{}, // Одна публикация
			}, nil).Once()

		pubReq := publish.NewPublishRequest(7, 9, cdtime.GetToday().AddDate(0, 0, 14))

		criteria := OneReleasePerDayCriteria{
			publicationRepo: df.publicationRepo,
		}

		result := criteria.Apply(pubReq)

		// Проверяем, что ограничение не сработало (публикаций не более одной)
		sCtx.Assert().Equal(0, result.Diff)
		sCtx.Assert().Equal(criteria2.ExplanationOK, result.Explanation)
	})
}

func TestOneReleasePerDayCriteriaSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(OneReleasePerDayCriteriaSuite))
}
