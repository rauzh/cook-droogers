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

type StatisticsPgRepo struct {
	db         *sqlx.DB
	txResolver *trmsqlx.CtxGetter
}

func NewStatisticsPgRepo(db *sql.DB) repo.StatisticsRepo {
	dbx := sqlx.NewDb(db, "pgx")

	return &StatisticsPgRepo{db: dbx, txResolver: trmsqlx.DefaultCtxGetter}
}

func (stat *StatisticsPgRepo) Create(ctx context.Context, stats *models.Statistics) (err error) {

	q := "INSERT INTO stats (streams, likes, creation_date, track_id) VALUES ($1, $2, $3, $4) RETURNING stat_id"

	var statID uint64
	err = stat.txResolver.DefaultTrOrDB(ctx, stat.db).QueryRowxContext(ctx, q,
		stats.Streams, stats.Likes, stats.Date, stats.TrackID).Scan(&statID)

	if err != nil {
		return
	}

	stats.StatID = statID
	return
}

func (stat *StatisticsPgRepo) GetForTrack(ctx context.Context, trackID uint64) ([]models.Statistics, error) {

	q := "SELECT stat_id, streams, likes, creation_date, track_id FROM stats WHERE track_id=$1"

	stats := make([]models.Statistics, 0)

	rows, err := stat.txResolver.DefaultTrOrDB(ctx, stat.db).QueryxContext(ctx, q, trackID)

	if errors.Is(err, sql.ErrNoRows) {
		return stats, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		curStat := models.Statistics{}
		err = rows.Scan(&curStat.StatID, &curStat.Streams, &curStat.Likes, &curStat.Date, &curStat.TrackID)
		if err != nil {
			return nil, err
		}

		stats = append(stats, curStat)
	}

	return stats, nil
}

func (stat *StatisticsPgRepo) GetByID(ctx context.Context, statID uint64) (*models.Statistics, error) {

	q := "SELECT stat_id, streams, likes, creation_date, track_id FROM stats WHERE stat_id=$1"

	curStat := models.Statistics{}
	err := stat.txResolver.DefaultTrOrDB(ctx, stat.db).QueryRowxContext(ctx, q, statID).Scan(
		&curStat.StatID, &curStat.Streams, &curStat.Likes, &curStat.Date, &curStat.TrackID)

	if err != nil {
		return nil, err
	}

	return &curStat, nil
}

func (stat *StatisticsPgRepo) GetAllGroupByTracksSince(ctx context.Context, date time.Time) (*map[uint64][]models.Statistics, error) {

	query := "SELECT likes, streams, track_id, stat_id, creation_date FROM stats WHERE stats.creation_date >= $1"

	rows, err := stat.txResolver.DefaultTrOrDB(ctx, stat.db).QueryxContext(ctx, query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	statsMap := make(map[uint64][]models.Statistics)

	for rows.Next() {
		curStat := models.Statistics{}
		if err := rows.Scan(&curStat.Likes, &curStat.Streams, &curStat.TrackID, &curStat.StatID, &curStat.Date); err != nil {
			return nil, err
		}
		statsMap[curStat.TrackID] = append(statsMap[curStat.TrackID], curStat)
	}

	return &statsMap, nil
}

func (stat *StatisticsPgRepo) CreateMany(ctx context.Context, stats []models.Statistics) error {

	for _, curStat := range stats {
		err := stat.Create(ctx, &curStat)
		if err != nil {
			return err
		}
	}

	return nil
}
