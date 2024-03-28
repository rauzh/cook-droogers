package service

import (
	artMocks "cookdroogers/internal/Artist/repo/mocks"
	artService "cookdroogers/internal/Artist/service/impl"
	mngMocks "cookdroogers/internal/Manager/repo/mocks"
	mngService "cookdroogers/internal/Manager/service/impl"
	reqMocks "cookdroogers/internal/Request/repo/mocks"
	transacMock "cookdroogers/internal/TransactionManager/mocks"
	usrMocks "cookdroogers/internal/User/repo/mocks"
	usrService "cookdroogers/internal/User/service/impl"
	"cookdroogers/models"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSignContractService_Apply(t *testing.T) {

	mockTrans := transacMock.NewTransactionManager(t)
	mockArtRepo := artMocks.NewArtistRepo(t)

	mockUsrRepo := usrMocks.NewUserRepo(t)

	mockMngRepo := mngMocks.NewManagerRepo(t)
	mockMngRepo.EXPECT().GetRandManagerID().Return(uint64(1337), nil).Once()

	mockRequestRepo := reqMocks.NewRequestRepo(t)
	y, m, d := time.Now().UTC().Date()
	mockRequestRepo.EXPECT().Create(
		&models.Request{
			Type:   models.SignRequest,
			Status: models.NewRequest,
			Date:   time.Date(y, m, d, 0, 0, 0, 0, time.UTC),
			Meta: map[string]string{
				"nickname": "aboba",
				"descr":    "",
			},
			ApplierID: uint64(777),
		},
	).Return(nil).Once()
	mockRequestRepo.EXPECT().Update(
		&models.Request{
			Type:   models.SignRequest,
			Status: models.OnApprovalRequest,
			Date:   time.Date(y, m, d, 0, 0, 0, 0, time.UTC),
			Meta: map[string]string{
				"nickname": "aboba",
				"descr":    "",
			},
			ApplierID: uint64(777),
			ManagerID: uint64(1337),
		},
	).Return(nil).Once()

	reqSvc := NewRequestServiceImpl(mockRequestRepo)
	mngSvc := mngService.NewManagerService(mockMngRepo)
	usrSvc := usrService.NewUserService(mockUsrRepo)
	artSvc := artService.NewArtistService(mockArtRepo)

	sctSvc := NewSignContractService(reqSvc, mngSvc, usrSvc, artSvc, mockTrans)

	err := sctSvc.Apply(777, "aboba")

	time.Sleep(time.Second)

	assert.Nil(t, err)
}

func TestSignContractService_Apply_Fail(t *testing.T) {

	mockTrans := transacMock.NewTransactionManager(t)

	mockArtRepo := artMocks.NewArtistRepo(t)

	mockUsrRepo := usrMocks.NewUserRepo(t)

	mockMngRepo := mngMocks.NewManagerRepo(t)

	mockRequestRepo := reqMocks.NewRequestRepo(t)
	y, m, d := time.Now().UTC().Date()
	mockRequestRepo.EXPECT().Create(
		&models.Request{
			Type:   models.SignRequest,
			Status: models.NewRequest,
			Date:   time.Date(y, m, d, 0, 0, 0, 0, time.UTC),
			Meta: map[string]string{
				"nickname": "aboba",
				"descr":    "",
			},
			ApplierID: uint64(777),
		},
	).Return(errors.New("db error")).Once()

	reqSvc := NewRequestServiceImpl(mockRequestRepo)
	mngSvc := mngService.NewManagerService(mockMngRepo)
	usrSvc := usrService.NewUserService(mockUsrRepo)
	artSvc := artService.NewArtistService(mockArtRepo)

	sctSvc := NewSignContractService(reqSvc, mngSvc, usrSvc, artSvc, mockTrans)

	err := sctSvc.Apply(777, "aboba")

	time.Sleep(time.Second)

	assert.Equal(t, err.Error(),
		"can't apply sign contract request with err can't create request info with error db error")
}
