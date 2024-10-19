//go:build integration
// +build integration

package integration_tests

import (
	"cookdroogers/integration_tests/buisness_logic"
	"cookdroogers/integration_tests/data_access"
	"cookdroogers/integration_tests/utils"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"os"
	"testing"
)

func BeforeAll() {
	pgInfo := utils.PostgresInfo{
		Host:     "postgres",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     "5432",
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	db, err := utils.InitDB(pgInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := os.ReadFile("/builds/rpp21u198/test-cd-73/backend/integration_tests/init.sql")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(string(text))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestIntegrationRunner(t *testing.T) {
	BeforeAll()

	suite.RunSuite(t, new(buisness_logic.ManagerIntegrationCoreSuite))
	suite.RunSuite(t, new(buisness_logic.ArtistIntegrationCoreSuite))
	suite.RunSuite(t, new(buisness_logic.UserIntegrationCoreSuite))
	suite.RunSuite(t, new(buisness_logic.TrackIntegrationCoreSuite))

	suite.RunSuite(t, new(data_access.ManagerIntegrationPgSuite))
	suite.RunSuite(t, new(data_access.ArtistIntegrationPgSuite))
	suite.RunSuite(t, new(data_access.UserIntegrationPgSuite))
	suite.RunSuite(t, new(data_access.TrackIntegrationPgSuite))
}
