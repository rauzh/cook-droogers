package postgres

import (
	"context"
	"cookdroogers/internal/repo"
	"cookdroogers/models"
	"database/sql"
	"errors"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
	"time"
)

type PublicationPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

func NewPublicationPgRepo(db *sql.DB) repo.PublicationRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &PublicationPgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter}
}

func (pub *PublicationPgRepo) Create(ctx context.Context, publication *models.Publication) error {

	q := "INSERT INTO publications(manager_id, release_id, creation_date) VALUES($1, $2, $3) RETURNING publication_id"

	var publicationID uint64
	err := pub.txResolver.DefaultTrOrDB(ctx, pub.db).QueryRowxContext(ctx, q,
		publication.ManagerID, publication.ReleaseID, publication.Date).Scan(&publicationID)

	if err != nil {
		return err
	}

	publication.PublicationID = publicationID
	return nil

}

func (pub *PublicationPgRepo) Get(ctx context.Context, publicationID uint64) (*models.Publication, error) {

	q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE publication_id=$1"

	publication := models.Publication{}
	err := pub.txResolver.DefaultTrOrDB(ctx, pub.db).QueryRowxContext(ctx, q,
		publicationID).Scan(&publication.PublicationID, &publication.Date, &publication.ManagerID, &publication.ReleaseID)

	if err != nil {
		return nil, err
	}

	return &publication, err
}

func (pub *PublicationPgRepo) GetAllByDate(ctx context.Context, date time.Time) ([]models.Publication, error) {

	q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE creation_date=$1"

	publications := make([]models.Publication, 0)

	rows, err := pub.txResolver.DefaultTrOrDB(ctx, pub.db).QueryxContext(ctx, q, date)

	if errors.Is(err, sql.ErrNoRows) {
		return publications, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		publication := models.Publication{}
		err = rows.Scan(&publication.PublicationID, &publication.Date, &publication.ManagerID, &publication.ReleaseID)

		if err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (pub *PublicationPgRepo) GetAllByManager(ctx context.Context, mng uint64) ([]models.Publication, error) {
	q := "SELECT publication_id, creation_date, manager_id, release_id FROM publications WHERE manager_id=$1"

	publications := make([]models.Publication, 0)

	rows, err := pub.txResolver.DefaultTrOrDB(ctx, pub.db).QueryxContext(ctx, q, mng)

	if errors.Is(err, sql.ErrNoRows) {
		return publications, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		publication := models.Publication{}
		err = rows.Scan(&publication.PublicationID, &publication.Date, &publication.ManagerID, &publication.ReleaseID)

		if err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (pub *PublicationPgRepo) GetAllByArtistSinceDate(ctx context.Context, date time.Time, artistID uint64) ([]models.Publication, error) {

	q := "SELECT p.publication_id, p.creation_date, p.manager_id, p.release_id " +
		"FROM publications p JOIN releases r ON p.release_id = r.release_id" +
		"WHERE r.artist_id=$1 AND p.creation_date>=$2;"

	publications := make([]models.Publication, 0)

	rows, err := pub.txResolver.DefaultTrOrDB(ctx, pub.db).QueryxContext(ctx, q, artistID, date)

	if errors.Is(err, sql.ErrNoRows) {
		return publications, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		publication := models.Publication{}
		err = rows.Scan(&publication.PublicationID, &publication.Date, &publication.ManagerID, &publication.ReleaseID)

		if err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil

}

func (pub *PublicationPgRepo) Update(ctx context.Context, publication *models.Publication) error {

	q := "UPDATE publications SET creation_date=$1, manager_id=$2, release_id=$3 WHERE publication_id=$4"

	_, err := pub.txResolver.DefaultTrOrDB(ctx, pub.db).ExecContext(ctx, q,
		publication.Date, publication.ManagerID, publication.ReleaseID, publication.PublicationID)

	return err
}
