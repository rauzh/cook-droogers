package postgres

import (
	"context"
	"cookdroogers/internal/requests/sign_contract"
	"cookdroogers/internal/requests/sign_contract/repo"
	"database/sql"
	"encoding/json"
	"errors"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	sqlx "github.com/jmoiron/sqlx"
)

var (
	ErrorInvalidRequestID    error = errors.New("invalid request id")
	ErrorNoNicknameInMeta    error = errors.New("no nickname in meta")
	ErrorNoDescriptionInMeta error = errors.New("no description in meta")
)

type SignContractRequestPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

type SignContractReqMetaPgDTO struct {
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
}

func NewSignContractRequestPgRepo(db *sql.DB) repo.SignContractRequestRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &SignContractRequestPgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter}
}

func (sctRepo *SignContractRequestPgRepo) SetMeta(ctx context.Context, signReq *sign_contract.SignContractRequest) error {

	if signReq.RequestID == 0 {
		return ErrorInvalidRequestID
	}

	meta := SignContractReqMetaPgDTO{
		Nickname:    signReq.Nickname,
		Description: signReq.Description,
	}

	metaJson, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	q := "UPDATE requests SET meta=$1 WHERE request_id=$2"

	err = sctRepo.txResolver.DefaultTrOrDB(ctx, sctRepo.db).QueryRowxContext(ctx, q,
		metaJson, signReq.RequestID).Scan()

	return err
}

func (sctRepo *SignContractRequestPgRepo) Get(ctx context.Context, id uint64) (*sign_contract.SignContractRequest, error) {

	q := "SELECT * FROM requests WHERE request_id=$1"

	var mngID sql.NullInt64
	signRequest := sign_contract.SignContractRequest{}
	var metaJson []byte
	err := sctRepo.txResolver.DefaultTrOrDB(ctx, sctRepo.db).QueryRowxContext(ctx, q, id).Scan(
		&signRequest.RequestID, &signRequest.Status, &signRequest.Type,
		&signRequest.Date, &metaJson, &mngID, &signRequest.ApplierID)

	if err != nil {
		return nil, err
	}

	if mngID.Valid {
		signRequest.ManagerID = uint64(mngID.Int64)
	}

	meta := SignContractReqMetaPgDTO{}
	err = json.Unmarshal(metaJson, &meta)
	if err != nil {
		return nil, err
	}

	signRequest.Nickname = meta.Nickname
	signRequest.Description = meta.Description

	return &signRequest, nil
}

func (sctRepo *SignContractRequestPgRepo) Update(ctx context.Context, signReq *sign_contract.SignContractRequest) error {

	if signReq.RequestID == 0 {
		return ErrorInvalidRequestID
	}

	meta := SignContractReqMetaPgDTO{
		Nickname:    signReq.Nickname,
		Description: signReq.Description,
	}

	metaJson, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	q := "UPDATE requests SET status=$2, type=$3, meta=$4, manager_id=$5, user_id=$6 WHERE request_id=$1"

	err = sctRepo.txResolver.DefaultTrOrDB(ctx, sctRepo.db).QueryRowxContext(ctx, q,
		signReq.RequestID, signReq.Status, signReq.Type, metaJson, signReq.ManagerID, signReq.ApplierID).Scan()
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}

	return err
}

func (sctRepo *SignContractRequestPgRepo) Create(ctx context.Context, signReq *sign_contract.SignContractRequest) error {

	meta := SignContractReqMetaPgDTO{
		Nickname:    signReq.Nickname,
		Description: signReq.Description,
	}

	metaJson, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	q := "INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING request_id"

	var mngID sql.NullInt64
	if signReq.ManagerID == 0 {
		mngID.Valid = false
	} else {
		mngID.Valid = true
		mngID.Int64 = int64(signReq.ManagerID)
	}

	err = sctRepo.txResolver.DefaultTrOrDB(ctx, sctRepo.db).QueryRowxContext(ctx, q,
		signReq.Status, signReq.Type, signReq.Date, metaJson, mngID, signReq.ApplierID).Scan(&signReq.RequestID)

	return err
}
