package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/internal/transactor"
	"cookdroogers/models"
	"database/sql"
	"errors"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
)

type ReleasePgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
	transactor transactor.Transactor
}

func NewReleasePgRepo(db *sql.DB, transactor transactor.Transactor) repo.ReleaseRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &ReleasePgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter, transactor: transactor}
}

func (rel *ReleasePgRepo) Create(ctx context.Context, release *models.Release) error {

	return rel.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		q := "INSERT INTO releases (title, status, creation_date, artist_id) VALUES ($1, $2, $3, $4) RETURNING release_id"

		var releaseID uint64

		err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).QueryRowxContext(ctx, q,
			release.Title, release.Status, release.DateCreation, release.ArtistID).Scan(&releaseID)

		if err != nil {
			return err
		}

		release.ReleaseID = releaseID

		for _, trackID := range release.Tracks {
			q = "UPDATE tracks SET release_id=$1 WHERE track_id=$2"

			_, err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).ExecContext(ctx, q, releaseID, trackID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (rel *ReleasePgRepo) Get(ctx context.Context, releaseID uint64) (*models.Release, error) {

	q := "SELECT title, status, creation_date, artist_id FROM releases WHERE release_id=$1"

	release := models.Release{}

	err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).QueryRowxContext(ctx, q, releaseID).Scan(
		&release.Title, &release.Status, &release.DateCreation, &release.ArtistID)

	if err != nil {
		return nil, err
	}

	release.ReleaseID = releaseID

	q = "SELECT track_id FROM tracks WHERE release_id=$1"

	rows, err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).QueryxContext(ctx, q, releaseID)
	if errors.Is(err, sql.ErrNoRows) {
		return &release, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trackID uint64
		err = rows.Scan(&trackID)
		if err != nil {
			return nil, err
		}

		release.Tracks = append(release.Tracks, trackID)
	}

	return &release, nil
}

func (rel *ReleasePgRepo) GetAllByArtist(ctx context.Context, artistID uint64) ([]models.Release, error) {

	q := "SELECT title, status, creation_date, release_id FROM releases WHERE artist_id=$1"

	releases := make([]models.Release, 0)

	rows, err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).QueryxContext(ctx, q, artistID)

	if errors.Is(err, sql.ErrNoRows) {
		return releases, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		release := models.Release{}

		err = rows.Scan(&release.Title, &release.Status, &release.DateCreation, &release.ReleaseID)
		if err != nil {
			return nil, err
		}

		release.ArtistID = artistID

		qTRK := "SELECT track_id FROM tracks WHERE release_id=$1"

		rowsTRK, err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).QueryxContext(ctx, qTRK, release.ReleaseID)
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		if err != nil {
			return nil, err
		}

		for rowsTRK.Next() {
			var trackID uint64
			err = rowsTRK.Scan(&trackID)
			if err != nil {
				return nil, err
			}

			release.Tracks = append(release.Tracks, trackID)
		}

		_ = rowsTRK.Close()

		releases = append(releases, release)
	}

	return releases, nil
}

func (rel *ReleasePgRepo) GetAllTracks(ctx context.Context, release *models.Release) ([]models.Track, error) {

	q1 := "SELECT track_id, title, genre, duration, type FROM tracks WHERE track_id=$1"
	q2 := "SELECT artist_id FROM track_artist WHERE track_id=$1"

	tracks := make([]models.Track, 1)

	for _, trackID := range release.Tracks {

		track := models.Track{}

		err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).QueryRowxContext(ctx, q1, trackID).Scan(
			&track.TrackID, &track.Title, &track.Genre, &track.Duration, &track.Type)

		if err != nil {
			return nil, err
		}

		rows, err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).QueryxContext(ctx, q2, trackID)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var artistID uint64
			err := rows.Scan(&artistID)
			if err != nil {
				return nil, err
			}

			track.Artists = append(track.Artists, artistID)
		}
		_ = rows.Close()
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (rel *ReleasePgRepo) Update(ctx context.Context, release *models.Release) error {

	q := "UPDATE releases SET title=$1, status=$2, creation_date=$3, artist_id=$4 WHERE release_id=$5"

	_, err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).ExecContext(ctx, q,
		release.Title, release.Status, release.DateCreation, release.ArtistID, release.ReleaseID)

	return err
}

func (rel *ReleasePgRepo) UpdateStatus(ctx context.Context, id uint64, stat models.ReleaseStatus) error {

	q := "UPDATE releases SET status=$1 WHERE release_id=$2"

	_, err := rel.txResolver.DefaultTrOrDB(ctx, rel.db).ExecContext(ctx, q, stat, id)

	return err
}
