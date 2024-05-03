package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"database/sql"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
)

type TrackPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

func NewTrackPgRepo(db *sql.DB) repo.TrackRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &TrackPgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter}
}

func (trk *TrackPgRepo) Create(ctx context.Context, track *models.Track) (uint64, error) {

	q := "INSERT INTO tracks (title, genre, duration, type) VALUES ($1, $2, $3, $4) RETURNING track_id"

	var trackID uint64
	err := trk.txResolver.DefaultTrOrDB(ctx, trk.db).QueryRowxContext(ctx, q,
		track.Title, track.Genre, track.Duration, track.Type).Scan(&trackID)
	if err != nil {
		return 0, err
	}

	track.TrackID = trackID

	q = "INSERT INTO track_artist (artist_id, track_id) VALUES ($1, $2) RETURNING track_artist_id"

	for _, artistID := range track.Artists {
		var connId uint64
		err = trk.txResolver.DefaultTrOrDB(ctx, trk.db).QueryRowxContext(ctx, q,
			artistID, track.TrackID).Scan(&connId)
		if err != nil {
			return 0, err
		}
	}

	return trackID, nil
}

func (trk *TrackPgRepo) Get(ctx context.Context, trackID uint64) (*models.Track, error) {

	q := "SELECT track_id, title, genre, duration, type FROM tracks WHERE track_id=$1"

	track := models.Track{}

	err := trk.txResolver.DefaultTrOrDB(ctx, trk.db).QueryRowxContext(ctx, q, trackID).Scan(
		&track.TrackID, &track.Title, &track.Genre, &track.Duration, &track.Type)

	if err != nil {
		return nil, err
	}

	return &track, nil
}
