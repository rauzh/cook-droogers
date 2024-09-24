package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"database/sql"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	sqlx "github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ArtistPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

func NewArtistPgRepo(db *sql.DB) repo.ArtistRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &ArtistPgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter}
}

func (art *ArtistPgRepo) Get(ctx context.Context, id uint64) (*models.Artist, error) {

	q := "SELECT artist_id, nickname, contract_due, activity, user_id, manager_id FROM artists WHERE artist_id=$1"

	artist := models.Artist{}
	err := art.txResolver.DefaultTrOrDB(ctx, art.db).QueryRowxContext(ctx, q, id).Scan(
		&artist.ArtistID,
		&artist.Nickname, &artist.ContractTerm, &artist.Activity,
		&artist.UserID, &artist.ManagerID)

	if err != nil {
		return nil, errors.Wrap(PgDbErr, err.Error())
	}

	return &artist, nil
}

func (art *ArtistPgRepo) GetByUserID(ctx context.Context, id uint64) (*models.Artist, error) {

	q := "SELECT artist_id, nickname, contract_due, activity, user_id, manager_id FROM artists WHERE user_id=$1"

	artist := models.Artist{}
	err := art.txResolver.DefaultTrOrDB(ctx, art.db).QueryRowxContext(ctx, q, id).Scan(
		&artist.ArtistID,
		&artist.Nickname, &artist.ContractTerm, &artist.Activity,
		&artist.UserID, &artist.ManagerID)

	if err != nil {
		return nil, errors.Wrap(PgDbErr, err.Error())
	}

	return &artist, nil
}

func (art *ArtistPgRepo) Update(ctx context.Context, artist *models.Artist) error {
	q := "UPDATE artists SET user_id=$1, nickname=$2, contract_due=$3, activity=$4, manager_id=$5 WHERE artist_id=$6 RETURNING *"

	err := art.txResolver.DefaultTrOrDB(ctx, art.db).QueryRowxContext(ctx, q,
		artist.UserID, artist.Nickname, artist.ContractTerm, artist.Activity, artist.ManagerID, artist.ArtistID).Scan(
		&artist.ArtistID,
		&artist.Nickname, &artist.ContractTerm, &artist.Activity,
		&artist.UserID, &artist.ManagerID)

	if err != nil {
		return errors.Wrap(PgDbErr, err.Error())
	}

	return nil
}

func (art *ArtistPgRepo) Create(ctx context.Context, artist *models.Artist) error {
	q := "INSERT INTO artists(user_id, nickname, contract_due, activity, manager_id)" +
		"VALUES($1, $2, $3, $4, $5) RETURNING artist_id"

	var artistID uint64
	err := art.txResolver.DefaultTrOrDB(ctx, art.db).QueryRowxContext(ctx, q,
		artist.UserID, artist.Nickname, artist.ContractTerm, artist.Activity, artist.ManagerID).Scan(&artistID)

	if err != nil {
		return errors.Wrap(PgDbErr, err.Error())
	}

	artist.ArtistID = artistID
	return nil
}
