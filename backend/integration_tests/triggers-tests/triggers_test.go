package triggers_test

import (
	"context"
	"cookdroogers/integration_tests/containers"
	cdtime "cookdroogers/pkg/time"
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestManagerTrigger_ERR(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	_, err = dbx.Query("INSERT INTO managers (user_id) VALUES ((SELECT user_id FROM users WHERE type <> 1 LIMIT 1));")

	assert.NotEqual(t, nil, err)

	assert.Equal(t, true, strings.Contains(err.Error(), "менеджером может стать только пользователь типа 1 (ManagerUser)"))
}

func TestManagerTrigger_OK(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var userId uint64
	err = dbx.QueryRowx(
		"UPDATE users SET type=1 WHERE user_id=(SELECT user_id FROM users WHERE type=0 LIMIT 1) RETURNING user_id").
		Scan(&userId)

	assert.Equal(t, nil, err)

	var managerId uint64
	err = dbx.QueryRowx("INSERT INTO managers (user_id) VALUES ($1) RETURNING manager_id;", userId).
		Scan(&managerId)

	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(5), managerId)
}

func TestArtistTrigger_ERR(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	_, err = dbx.Query(
		"INSERT INTO artists (user_id, manager_id) VALUES ((SELECT user_id FROM users WHERE type <> 2 LIMIT 1), 1);")

	assert.NotEqual(t, nil, err)

	assert.Equal(t, true, strings.Contains(err.Error(), "артистом может стать только пользователь типа 2 (ArtistUser)"))
}

func TestArtistTrigger_OK(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var userId uint64
	err = dbx.QueryRowx(
		"UPDATE users SET type=2 WHERE user_id=(SELECT user_id FROM users WHERE type=0 LIMIT 1) RETURNING user_id").
		Scan(&userId)

	assert.Equal(t, nil, err)

	var artistId uint64
	err = dbx.QueryRowx("INSERT INTO artists (user_id, manager_id, nickname) VALUES ($1, 1, 'huh') RETURNING artist_id;", userId).
		Scan(&artistId)

	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(3), artistId)
}

func TestPublicationTrigger_ERR(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var releaseId, artistId uint64
	err = dbx.QueryRowx("SELECT release_id, artist_id FROM releases WHERE title='old-test-album'").
		Scan(&releaseId, &artistId)
	assert.Equal(t, nil, err)

	var managerId uint64
	err = dbx.QueryRowx("SELECT manager_id FROM artists WHERE artist_id=$1", artistId).
		Scan(&managerId)
	assert.Equal(t, nil, err)

	_, err = dbx.Query(
		"INSERT INTO publications (release_id, manager_id) VALUES ($1, (SELECT manager_id FROM managers WHERE manager_id <> $2 LIMIT 1));",
		releaseId, managerId)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, true, strings.Contains(err.Error(),
		"ответственным менеджером за публикацию должен быть менеджер артиста-владельца релиза"))
}

func TestPublicationTrigger_OK(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var releaseId, artistId uint64
	err = dbx.QueryRowx("SELECT release_id, artist_id FROM releases WHERE title='old-test-album'").
		Scan(&releaseId, &artistId)
	assert.Equal(t, nil, err)

	var managerId uint64
	err = dbx.QueryRowx("SELECT manager_id FROM artists WHERE artist_id=$1", artistId).
		Scan(&managerId)
	assert.Equal(t, nil, err)

	var publicationId uint64
	err = dbx.QueryRowx(
		"INSERT INTO publications (release_id, manager_id) VALUES ($1, $2) returning publication_id;",
		releaseId, managerId).Scan(&publicationId)

	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), publicationId)
}

func TestTrackTrigger_OK(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	testTrackTitle := "testo"

	var trackId uint64
	err = dbx.QueryRowx("SELECT track_id FROM tracks WHERE title=$1", testTrackTitle).
		Scan(&trackId)
	assert.Equal(t, true, errors.Is(err, sql.ErrNoRows)) // доказали, что такого трека нет в tracks

	err = dbx.QueryRowx(
		"INSERT INTO tracks (title, release_id) VALUES ($1, $2) RETURNING track_id;",
		testTrackTitle, 1).Scan(&trackId)

	assert.Equal(t, nil, err)

	var linkTrackId uint64
	err = dbx.QueryRowx("SELECT track_id FROM track_artist WHERE track_id=$1", trackId).
		Scan(&linkTrackId) // доказали, что вставилось сразу и в track_artist

	assert.Equal(t, trackId, linkTrackId)
}

func TestPublishReqTrigger_ERRemptyMeta(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var releaseId, artistId uint64
	err = dbx.QueryRowx("SELECT release_id, artist_id FROM releases WHERE title='old-test-album'").
		Scan(&releaseId, &artistId)
	assert.Equal(t, nil, err)

	var managerId, userId uint64
	err = dbx.QueryRowx("SELECT manager_id, user_id FROM artists WHERE artist_id=$1", artistId).
		Scan(&managerId, &userId)
	assert.Equal(t, nil, err)

	var publicationId uint64
	err = dbx.QueryRowx(
		"INSERT INTO publications (release_id, manager_id) VALUES ($1, $2) returning publication_id;",
		releaseId, managerId).Scan(&publicationId)

	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), publicationId) //  создали публикацию

	meta := make(map[string]string)
	metaJson, err := json.Marshal(meta)
	assert.Equal(t, nil, err)

	var requestId uint64
	err = dbx.QueryRowx("INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)"+
		" VALUES ('New', 'Publish', '2024-9-9'::timestamp, $1, $2, $3) RETURNING request_id",
		metaJson, managerId, userId).Scan(&requestId)

	assert.Equal(t, true, strings.Contains(err.Error(),
		"Отсутствует один или несколько обязательных ключей в метаданных JSON"))
}

func TestPublishReqTrigger_ERRinvalidDate(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var releaseId, artistId uint64
	err = dbx.QueryRowx("SELECT release_id, artist_id FROM releases WHERE title='old-test-album'").
		Scan(&releaseId, &artistId)
	assert.Equal(t, nil, err)

	var managerId, userId uint64
	err = dbx.QueryRowx("SELECT manager_id, user_id FROM artists WHERE artist_id=$1", artistId).
		Scan(&managerId, &userId)
	assert.Equal(t, nil, err)

	var publicationId uint64
	err = dbx.QueryRowx(
		"INSERT INTO publications (release_id, manager_id) VALUES ($1, $2) returning publication_id;",
		releaseId, managerId).Scan(&publicationId)

	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), publicationId) //  создали публикацию

	type publishRequestMetaPgDTO struct {
		ReleaseID    uint64 `json:"release_id"`
		Grade        int    `json:"grade"`
		ExpectedDate string `json:"expected_date"`
		Description  string `json:"description"`
	}

	meta := publishRequestMetaPgDTO{
		ReleaseID:    releaseId,
		ExpectedDate: "hehehe",
	}
	metaJson, err := json.Marshal(meta)
	assert.Equal(t, nil, err)

	var requestId uint64
	err = dbx.QueryRowx("INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)"+
		" VALUES ('New', 'Publish', '2024-9-9'::timestamp, $1, $2, $3) RETURNING request_id",
		metaJson, managerId, userId).Scan(&requestId)

	assert.Equal(t, true, strings.Contains(err.Error(),
		"invalid input syntax for type timestamp"))
}

func TestPublishReqTrigger_ERRinvalidReleaseId(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var releaseId, artistId uint64
	err = dbx.QueryRowx("SELECT release_id, artist_id FROM releases WHERE title='old-test-album'").
		Scan(&releaseId, &artistId)
	assert.Equal(t, nil, err)

	var managerId, userId uint64
	err = dbx.QueryRowx("SELECT manager_id, user_id FROM artists WHERE artist_id=$1", artistId).
		Scan(&managerId, &userId)
	assert.Equal(t, nil, err)

	var publicationId uint64
	err = dbx.QueryRowx(
		"INSERT INTO publications (release_id, manager_id) VALUES ($1, $2) returning publication_id;",
		releaseId, managerId).Scan(&publicationId)

	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), publicationId) //  создали публикацию

	type publishRequestMetaPgDTO struct {
		ReleaseID    string    `json:"release_id"`
		Grade        int       `json:"grade"`
		ExpectedDate time.Time `json:"expected_date"`
		Description  string    `json:"description"`
	}

	meta := publishRequestMetaPgDTO{
		ReleaseID:    "popo",
		ExpectedDate: cdtime.GetToday().AddDate(0, 3, 0),
	}
	metaJson, err := json.Marshal(meta)
	assert.Equal(t, nil, err)

	var requestId uint64
	err = dbx.QueryRowx("INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)"+
		" VALUES ('New', 'Publish', '2024-9-9'::timestamp, $1, $2, $3) RETURNING request_id",
		metaJson, managerId, userId).Scan(&requestId)

	assert.Equal(t, true, strings.Contains(err.Error(),
		"\"release_id\" в метаданных JSON должен быть натуральным числом"))
}

func TestPublishReqTrigger_OK(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var releaseId, artistId uint64
	err = dbx.QueryRowx("SELECT release_id, artist_id FROM releases WHERE title='old-test-album'").
		Scan(&releaseId, &artistId)
	assert.Equal(t, nil, err)

	var managerId, userId uint64
	err = dbx.QueryRowx("SELECT manager_id, user_id FROM artists WHERE artist_id=$1", artistId).
		Scan(&managerId, &userId)
	assert.Equal(t, nil, err)

	var publicationId uint64
	err = dbx.QueryRowx(
		"INSERT INTO publications (release_id, manager_id) VALUES ($1, $2) returning publication_id;",
		releaseId, managerId).Scan(&publicationId)

	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), publicationId) //  создали публикацию

	type publishRequestMetaPgDTO struct {
		ReleaseID    uint64    `json:"release_id"`
		Grade        int       `json:"grade"`
		ExpectedDate time.Time `json:"expected_date"`
		Description  string    `json:"description"`
	}

	meta := publishRequestMetaPgDTO{
		ReleaseID:    releaseId,
		ExpectedDate: cdtime.GetToday().AddDate(0, 3, 0),
	}
	metaJson, err := json.Marshal(meta)
	assert.Equal(t, nil, err)

	var requestId uint64
	err = dbx.QueryRowx("INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)"+
		" VALUES ('New', 'Publish', '2024-9-9'::timestamp, $1, $2, $3) RETURNING request_id",
		metaJson, managerId, userId).Scan(&requestId)

	assert.Equal(t, nil, err)
}

func TestSignReqTrigger_ERRnoNickname(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var userId uint64
	err = dbx.QueryRowx(
		"UPDATE users SET type=2 WHERE user_id=(SELECT user_id FROM users WHERE type=0 LIMIT 1) RETURNING user_id").
		Scan(&userId)

	assert.Equal(t, nil, err) // получили юзера типа "артист"

	type publishRequestMetaPgDTO struct {
		Description string `json:"description"`
	}

	meta := publishRequestMetaPgDTO{}
	metaJson, err := json.Marshal(meta)
	assert.Equal(t, nil, err)

	var requestId uint64
	err = dbx.QueryRowx("INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)"+
		" VALUES ('New', 'Sign', '2024-9-9'::timestamp, $1, $2, $3) RETURNING request_id",
		metaJson, uint64(1), userId).Scan(&requestId)

	assert.Equal(t, true, strings.Contains(err.Error(),
		"Отсутствует обязательный ключ в метаданных JSON: \"nickname\""))
}

func TestSignReqTrigger_ERRemptyNickname(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var userId uint64
	err = dbx.QueryRowx(
		"UPDATE users SET type=2 WHERE user_id=(SELECT user_id FROM users WHERE type=0 LIMIT 1) RETURNING user_id").
		Scan(&userId)

	assert.Equal(t, nil, err) // получили юзера типа "артист"

	type publishRequestMetaPgDTO struct {
		Nickname    string `json:"nickname"`
		Description string `json:"description"`
	}

	meta := publishRequestMetaPgDTO{}
	metaJson, err := json.Marshal(meta)
	assert.Equal(t, nil, err)

	var requestId uint64
	err = dbx.QueryRowx("INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)"+
		" VALUES ('New', 'Sign', '2024-9-9'::timestamp, $1, $2, $3) RETURNING request_id",
		metaJson, uint64(1), userId).Scan(&requestId)

	assert.Equal(t, true, strings.Contains(err.Error(),
		"nickname не должен быть пустой строкой"))
}

func TestSignReqTrigger_OK(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	dbx := sqlx.NewDb(db, "pgx")

	var userId uint64
	err = dbx.QueryRowx(
		"UPDATE users SET type=2 WHERE user_id=(SELECT user_id FROM users WHERE type=0 LIMIT 1) RETURNING user_id").
		Scan(&userId)

	assert.Equal(t, nil, err) // получили юзера типа "артист"

	type publishRequestMetaPgDTO struct {
		Nickname    string `json:"nickname"`
		Description string `json:"description"`
	}

	meta := publishRequestMetaPgDTO{
		Nickname: "hehehe",
	}
	metaJson, err := json.Marshal(meta)
	assert.Equal(t, nil, err)

	var requestId uint64
	err = dbx.QueryRowx("INSERT INTO requests (status, type, creation_date, meta, manager_id, user_id)"+
		" VALUES ('New', 'Sign', '2024-9-9'::timestamp, $1, $2, $3) RETURNING request_id",
		metaJson, uint64(1), userId).Scan(&requestId)

	assert.Equal(t, nil, err)
}
