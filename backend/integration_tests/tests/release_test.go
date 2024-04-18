package tests

import (
	"context"
	"cookdroogers/integration_tests/containers"
	postgres "cookdroogers/internal/repo/pg"
	"cookdroogers/internal/transactor"
	transactor2 "cookdroogers/internal/transactor/trm"
	"cookdroogers/models"
	cdtime "cookdroogers/pkg/time"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReleaseTracks_CreateGet(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	var trmng transactor.Transactor
	trmng = transactor2.NewATtrm(manager.Must(trmsqlx.NewDefaultFactory(sqlx.NewDb(db, "pgx"))))

	ctx := context.Background()

	artistRepo := postgres.NewArtistPgRepo(db)
	trackRepo := postgres.NewTrackPgRepo(db)
	releaseRepo := postgres.NewReleasePgRepo(db, trmng)

	err = artistRepo.Create(ctx, &models.Artist{
		ArtistID:     1,
		UserID:       2,
		Nickname:     "zeliboba",
		ContractTerm: cdtime.GetEndOfContract(),
		Activity:     true,
		ManagerID:    1,
	})

	id1, err := trackRepo.Create(ctx, &models.Track{
		Title:    "aboba",
		Duration: 120,
		Genre:    "rockkk",
		Type:     "song",
		Artists:  []uint64{1},
	})

	assert.Equal(t, nil, err)

	id2, err := trackRepo.Create(ctx, &models.Track{
		Title:    "kekue",
		Duration: 110,
		Genre:    "rappp",
		Type:     "song",
		Artists:  []uint64{1},
	})

	assert.Equal(t, nil, err)

	release := &models.Release{
		Title:        "album-1",
		Status:       models.UnpublishedRelease,
		DateCreation: cdtime.GetToday(),
		Tracks:       []uint64{id1, id2},
		ArtistID:     1,
	}

	err = releaseRepo.Create(ctx, release)
	assert.Equal(t, nil, err)

	releaseCopy, err := releaseRepo.Get(ctx, release.ReleaseID)
	assert.Equal(t, nil, err)

	assert.Equal(t, release.ReleaseID, releaseCopy.ReleaseID)
	assert.Equal(t, len(release.Tracks), len(releaseCopy.Tracks))
}
