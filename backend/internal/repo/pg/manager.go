package postgres

import (
	"context"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	"database/sql"
	"errors"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
)

type ManagerPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
	transactor transactor.Transactor
}

func (mng *ManagerPgRepo) Create(ctx context.Context, manager *models.Manager) error {

	return mng.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		q := "INSERT INTO managers(user_id) VALUES($1) RETURNING manager_id"

		var managerID uint64
		err := mng.txResolver.DefaultTrOrDB(ctx, mng.db).QueryRowxContext(ctx, q,
			manager.UserID).Scan(&managerID)

		if err != nil {
			return err
		}

		for _, artistID := range manager.Artists {
			q := "UPDATE artists SET manager_id=$1 WHERE artist_id=$2"

			_, err := mng.txResolver.DefaultTrOrDB(ctx, mng.db).ExecContext(ctx, q,
				managerID, artistID)

			if err != nil {
				return err
			}
		}

		manager.ManagerID = managerID
		return nil
	})

}

func (mng *ManagerPgRepo) Get(ctx context.Context, userID uint64) (*models.Manager, error) {
	q := "SELECT manager_id, user_id FROM managers WHERE manager_id=$1"

	manager := models.Manager{}
	err := mng.txResolver.DefaultTrOrDB(ctx, mng.db).QueryRowxContext(ctx, q,
		userID).Scan(&manager.ManagerID, &manager.UserID)

	if err != nil {
		return nil, err
	}

	q = "SELECT artist_id FROM artists WHERE manager_id=$1"

	rows, err := mng.txResolver.DefaultTrOrDB(ctx, mng.db).QueryxContext(ctx, q,
		manager.ManagerID)

	if errors.Is(err, sql.ErrNoRows) {
		return &manager, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var artistID uint64
		err = rows.Scan(&artistID)
		if err != nil {
			return nil, err
		}
		manager.Artists = append(manager.Artists, artistID)
	}

	return &manager, nil
}

func (mng *ManagerPgRepo) GetRandManagerID(ctx context.Context) (uint64, error) {

	q := "SELECT manager_id FROM managers ORDER BY random() LIMIT 1;"

	var managerID uint64
	err := mng.txResolver.DefaultTrOrDB(ctx, mng.db).QueryRowxContext(ctx, q).Scan(&managerID)

	return managerID, err
}
