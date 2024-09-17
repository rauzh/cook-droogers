package pg

import (
	"context"
	"cookdroogers/internal/requests/base"
	repo "cookdroogers/internal/requests/base/repo"
	"database/sql"
	"errors"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
)

type RequestPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

func NewRequestPgRepo(db *sql.DB) repo.RequestRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &RequestPgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter}
}

func (reqrepo *RequestPgRepo) GetAllByManagerID(ctx context.Context, mngID uint64) ([]base.Request, error) {

	q := "SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE manager_id=$1"

	reqs := make([]base.Request, 0)

	rows, err := reqrepo.txResolver.DefaultTrOrDB(ctx, reqrepo.db).QueryxContext(ctx, q, mngID)

	if errors.Is(err, sql.ErrNoRows) {
		return reqs, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		req := base.Request{}
		var mngID sql.NullInt64

		err := rows.Scan(&req.RequestID, &req.Status, &req.Type, &req.Date, &mngID, &req.ApplierID)
		if err != nil {
			return nil, err
		}

		if mngID.Valid {
			req.ManagerID = uint64(mngID.Int64)
		}
		reqs = append(reqs, req)
	}

	return reqs, nil
}

func (reqrepo *RequestPgRepo) GetByID(ctx context.Context, reqID uint64) (*base.Request, error) {
	q := "SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE request_id=$1"

	req := base.Request{}

	var mngID sql.NullInt64

	err := reqrepo.txResolver.DefaultTrOrDB(ctx, reqrepo.db).QueryRowxContext(ctx, q, reqID).
		Scan(&req.RequestID, &req.Status, &req.Type, &req.Date, &mngID, &req.ApplierID)
	if err != nil {
		return nil, err
	}

	if mngID.Valid {
		req.ManagerID = uint64(mngID.Int64)
	}

	return &req, nil
}

func (reqrepo *RequestPgRepo) GetAllByUserID(ctx context.Context, userID uint64) ([]base.Request, error) {

	q := "SELECT request_id, status, type, creation_date, manager_id, user_id FROM requests WHERE user_id=$1"

	reqs := make([]base.Request, 0)

	rows, err := reqrepo.txResolver.DefaultTrOrDB(ctx, reqrepo.db).QueryxContext(ctx, q, userID)

	if errors.Is(err, sql.ErrNoRows) {
		return reqs, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		req := base.Request{}

		var mngID sql.NullInt64

		err := rows.Scan(&req.RequestID, &req.Status, &req.Type, &req.Date, &mngID, &req.ApplierID)
		if err != nil {
			return nil, err
		}

		if mngID.Valid {
			req.ManagerID = uint64(mngID.Int64)
		}

		reqs = append(reqs, req)
	}

	return reqs, nil
}
