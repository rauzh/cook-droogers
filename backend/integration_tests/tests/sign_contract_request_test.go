package tests

//
//import (
//	"context"
//	"cookdroogers/integration_tests/containers"
//	"cookdroogers/internal/requests/base"
//	"cookdroogers/internal/requests/sign_contract"
//	signReqPgRepo "cookdroogers/internal/requests/sign_contract/repo/pg"
//	cdtime "cookdroogers/pkg/time"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestSignContractRequest_CreateGet(t *testing.T) {
//	dbContainer, db, err := containers.SetupTestDatabase()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	defer func() {
//		_ = dbContainer.Terminate(context.Background())
//	}()
//
//	signreqrepo := signReqPgRepo.NewSignContractRequestPgRepo(db)
//
//	ctx := context.Background()
//
//	signReq := sign_contract.SignContractRequest{
//		Request: base.Request{
//			Type:      sign_contract.SignRequest,
//			Status:    base.OnApprovalRequest,
//			ApplierID: 2,
//			ManagerID: 1,
//			Date:      cdtime.GetToday(),
//		},
//		Nickname:    "zeliboba",
//		Description: "",
//	}
//
//	err = signreqrepo.Create(ctx, &signReq)
//
//	assert.Equal(t, nil, err)
//
//	signReqCopy, err := signreqrepo.Get(ctx, signReq.RequestID)
//
//	assert.Equal(t, signReqCopy.ManagerID, signReq.ManagerID)
//	assert.Equal(t, signReqCopy.RequestID, signReq.RequestID)
//	assert.Equal(t, signReqCopy.Type, signReq.Type)
//	assert.Equal(t, signReqCopy.ApplierID, signReq.ApplierID)
//	assert.Equal(t, signReqCopy.Date, signReq.Date)
//	assert.Equal(t, signReqCopy.Nickname, signReq.Nickname)
//	assert.Equal(t, signReqCopy.Description, signReq.Description)
//}
