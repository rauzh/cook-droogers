package pg

import (
	"context"
	"cookdroogers/internal/requests/publish"
	"cookdroogers/internal/requests/publish/repo"
	"database/sql"
	"encoding/json"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
	"time"
)

type PublishRequestPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

type PublishRequestMetaPgDTO struct {
	ReleaseID    uint64    `json:"release_id"`
	Grade        int       `json:"grade"`
	ExpectedDate time.Time `json:"expected_date"`
	Description  string    `json:"description"`
}

func NewPublishRequestPgRepo(db *sql.DB) repo.PublishRequestRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &PublishRequestPgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter}
}

func (pubreqrepo *PublishRequestPgRepo) Create(ctx context.Context, req *publish.PublishRequest) error {

	q := "INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)" +
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING request_id"

	meta := PublishRequestMetaPgDTO{
		ReleaseID:    req.ReleaseID,
		Grade:        req.Grade,
		ExpectedDate: req.ExpectedDate,
		Description:  req.Description,
	}

	metaJson, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	err = pubreqrepo.txResolver.DefaultTrOrDB(ctx, pubreqrepo.db).QueryRowxContext(ctx, q,
		req.Status, req.Type, req.Date, metaJson, req.ManagerID, req.ApplierID).Scan(&req.RequestID)

	return err
}

func (pubreqrepo *PublishRequestPgRepo) Get(ctx context.Context, id uint64) (*publish.PublishRequest, error) {

	q := "SELECT request_id, status, type, creation_date, meta, manager_id, user_id FROM requests WHERE request_id=$1"

	metaJson := make([]byte, 0)
	pubreq := publish.PublishRequest{}
	err := pubreqrepo.txResolver.DefaultTrOrDB(ctx, pubreqrepo.db).QueryRowxContext(ctx, q, id).Scan(
		&pubreq.RequestID, &pubreq.Status, &pubreq.Type, &pubreq.Date, &metaJson, &pubreq.ManagerID, &pubreq.ApplierID)

	if err != nil {
		return nil, err
	}

	meta := PublishRequestMetaPgDTO{}
	err = json.Unmarshal(metaJson, &meta)

	pubreq.ReleaseID = meta.ReleaseID
	pubreq.Grade = meta.Grade
	pubreq.ExpectedDate = meta.ExpectedDate
	pubreq.Description = meta.Description

	return &pubreq, nil
}

func (pubreqrepo *PublishRequestPgRepo) Update(ctx context.Context, req *publish.PublishRequest) error {

	q := "UPDATE requests SET status=$1, type=$2, creation_date=$3, meta=$4, manager_id=$5, user_id=$6 WHERE request_id=$7"

	meta := PublishRequestMetaPgDTO{
		ReleaseID:    req.ReleaseID,
		Grade:        req.Grade,
		ExpectedDate: req.ExpectedDate,
		Description:  req.Description,
	}

	metaJson, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	_, err = pubreqrepo.txResolver.DefaultTrOrDB(ctx, pubreqrepo.db).ExecContext(ctx, q,
		req.Status, req.Type, req.Date, metaJson, req.ManagerID, req.ApplierID, req.RequestID)

	return err
}

func (pubreqrepo *PublishRequestPgRepo) SetMeta(ctx context.Context, req *publish.PublishRequest) error {

	q := "UPDATE requests SET meta=$1 WHERE request_id=$2"

	meta := PublishRequestMetaPgDTO{
		ReleaseID:    req.ReleaseID,
		Grade:        req.Grade,
		ExpectedDate: req.ExpectedDate,
		Description:  req.Description,
	}

	metaJson, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	_, err = pubreqrepo.txResolver.DefaultTrOrDB(ctx, pubreqrepo.db).ExecContext(ctx, q, metaJson, req.RequestID)

	return err
}
