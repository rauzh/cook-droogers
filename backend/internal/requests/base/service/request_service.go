package service

import (
	"context"
	"cookdroogers/internal/requests/base"
	"cookdroogers/internal/requests/base/repo"
	"github.com/pkg/errors"
	"log/slog"
)

var DBerr error = errors.New("can't get reqs with err")

type IRequestService interface {
	GetAllByManagerID(uint64) ([]base.Request, error)
	GetAllByUserID(uint64) ([]base.Request, error)
	GetByID(uint64) (*base.Request, error)
}

type RequestService struct {
	repo repo.RequestRepo

	logger *slog.Logger
}

func NewRequestService(r repo.RequestRepo, logger *slog.Logger) IRequestService {
	return &RequestService{repo: r, logger: logger}
}

func (reqSvc *RequestService) GetAllByManagerID(id uint64) ([]base.Request, error) {

	reqs, err := reqSvc.repo.GetAllByManagerID(context.Background(), id)

	if err != nil {
		return nil, errors.Wrap(DBerr, err.Error())
	}

	return reqs, nil
}

func (reqSvc *RequestService) GetAllByUserID(id uint64) ([]base.Request, error) {

	reqs, err := reqSvc.repo.GetAllByUserID(context.Background(), id)

	if err != nil {
		return nil, errors.Wrap(DBerr, err.Error())
	}

	return reqs, nil
}

func (reqSvc *RequestService) GetByID(id uint64) (*base.Request, error) {

	req, err := reqSvc.repo.GetByID(context.Background(), id)

	if err != nil {
		reqSvc.logger.Error("REQ SVC: GetByID", "error", err.Error())
		return nil, errors.Wrap(DBerr, err.Error())
	}

	reqSvc.logger.Info("REQ SVC: GetByID SUCCESS")

	return req, nil
}
