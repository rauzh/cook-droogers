package containers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"os"

	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	USER     = "rauzh"
	PASSWORD = "1337"
	DBNAME   = "cook_droogers_test"
)

func SetupTestDatabase() (testcontainers.Container, *sql.DB, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:13.3",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       DBNAME,
			"POSTGRES_PASSWORD": PASSWORD,
			"POSTGRES_USER":     USER,
		},
	}

	dbContainer, _ := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	dsnPGConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port.Int(), USER, PASSWORD, DBNAME)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		return dbContainer, nil, err
	}

	err = db.Ping()
	if err != nil {
		return dbContainer, nil, err
	}
	db.SetMaxOpenConns(10)

	text, err := os.ReadFile("/Users/rauzh/Desktop/PPO_DB/cook-droogers/backend/integration_tests/init-tests.sql")
	if err != nil {
		return dbContainer, nil, err
	}

	if _, err := db.Exec(string(text)); err != nil {
		fmt.Println(err)
		return dbContainer, nil, err
	}

	return dbContainer, db, nil
}
