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

type _depFields struct {
	publicationRepo *mocks.PublicationRepo
	artistRepo      *mocks.ArtistRepo
}

// TEST_HW: default test
type ArtistReleaseLimitPerSeasonCriteriaSuite struct {
	suite.Suite
}

func _newMockArtistReleaseLimitDepFields(t provider.T) *_depFields {
	return &_depFields{
		publicationRepo: mocks.NewPublicationRepo(t),
		artistRepo:      mocks.NewArtistRepo(t),
	}
}

func (s *ArtistReleaseLimitPerSeasonCriteriaSuite) TestApply_LimitExceeded(t provider.T) {
	t.Title("Apply: Artist exceeded release limit for the season")
	t.Tags("ArtistReleaseLimitPerSeasonCriteria")
	t.Parallel()
	t.WithNewStep("Limit exceeded", func(sCtx provider.StepCtx) {

		df := _newMockArtistReleaseLimitDepFields(t)

		df.artistRepo.EXPECT().GetByUserID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
			Return(&models.Artist{ArtistID: 10}, nil).Once()

		df.publicationRepo.EXPECT().GetAllByArtistSinceDate(mock.AnythingOfType("context.backgroundCtx"),
			cdtime.RelevantPeriod(), uint64(10)).
			Return([]models.Publication{
				{}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			}, nil).Once()

		pubReq := publish.NewPublishRequest(7, 9, cdtime.GetToday().AddDate(0, 0, 14))

		criteria := ArtistReleaseLimitPerSeasonCriteria{
			publicationRepo: df.publicationRepo,
			artistRepo:      df.artistRepo,
		}

		result := criteria.Apply(pubReq)

		sCtx.Assert().Equal(DiffArtistReleaseLimitPerSeason, result.Diff)
		sCtx.Assert().Equal(ExplanationArtistReleaseLimit, result.Explanation)
	})
}

func (s *ArtistReleaseLimitPerSeasonCriteriaSuite) TestApply_NoLimitExceeded(t provider.T) {
	t.Title("Apply: Artist within release limit for the season")
	t.Tags("ArtistReleaseLimitPerSeasonCriteria")
	t.Parallel()
	t.WithNewStep("Within limit", func(sCtx provider.StepCtx) {

		df := _newMockArtistReleaseLimitDepFields(t)

		df.artistRepo.EXPECT().GetByUserID(mock.AnythingOfType("context.backgroundCtx"), uint64(7)).
			Return(&models.Artist{ArtistID: 10}, nil).Once()

		df.publicationRepo.EXPECT().GetAllByArtistSinceDate(mock.AnythingOfType("context.backgroundCtx"),
			cdtime.RelevantPeriod(), uint64(10)).
			Return([]models.Publication{
				{}, {},
			}, nil).Once()

		pubReq := publish.NewPublishRequest(7, 9, cdtime.GetToday().AddDate(0, 0, 14))

		criteria := ArtistReleaseLimitPerSeasonCriteria{
			publicationRepo: df.publicationRepo,
			artistRepo:      df.artistRepo,
		}

		result := criteria.Apply(pubReq)

		// Проверяем, что ограничение не сработало
		sCtx.Assert().Equal(0, result.Diff)
		sCtx.Assert().Equal(criteria2.ExplanationOK, result.Explanation)
	})
}

func TestArtistReleaseLimitPerSeasonCriteriaSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(ArtistReleaseLimitPerSeasonCriteriaSuite))
}
