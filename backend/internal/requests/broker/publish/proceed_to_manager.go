package publish

import (
	"context"
	"cookdroogers/internal/requests/base"
	criteria "cookdroogers/internal/requests/criteria_controller"
	"cookdroogers/internal/requests/publish"
	"fmt"
)

func (handler *PublishProceedToManagerConsumerHandler) proceedToManager(pubReq *publish.PublishRequest) error {

	ctx := context.Background()

	pubReq.Status = base.ProcessingRequest

	err := handler.publishRepo.Update(ctx, pubReq)
	if err != nil {
		return fmt.Errorf("cant proceed publish request to manager: update repo with err %w", err)
	}

	handler.computeDegree(pubReq)

	artist, err := handler.artistRepo.GetByUserID(ctx, pubReq.ApplierID)
	if err != nil {
		return fmt.Errorf("cant proceed publish request to manager: get artist with err %w", err)
	}

	pubReq.ManagerID = artist.ManagerID
	pubReq.Status = base.OnApprovalRequest

	return handler.publishRepo.Update(ctx, pubReq)
}

func (handler *PublishProceedToManagerConsumerHandler) computeDegree(pubReq *publish.PublishRequest) {

	summaryDiff := handler.criterias.Apply(pubReq)

	pubReq.Grade = summaryDiff.ResultDiff
	for criteriaName, criteriaDiff := range summaryDiff.ResultExplanation {
		pubReq.Description += criteria.DiffToString(criteriaName, criteriaDiff.Explanation, criteriaDiff.Diff)
	}
}
