package utils

import "os"

func IsUnitTestsFailed() bool {
	return os.Getenv("UNIT_SUCCESS") != "1"
}

func IsIntegrationTestsFailed() bool {
	return os.Getenv("INTEGRATION_SUCCESS") != "1"
}
