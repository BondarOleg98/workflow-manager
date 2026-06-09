package unit

import (
	"testing"
	"workflowmanager/app/db"
	"workflowmanager/app/util"
)

func TestCreateNonSslPgPool(test *testing.T) {
	const configFilePath string = "../../app/resources/test_env.yaml"
	util.LoadConfigs(configFilePath)
	actualPgPool := db.CreatePgPool()
	expectedPgPool := db.Pool{
		DriverName:    "postgres",
		ConnectionUrl: "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable",
	}
	if expectedPgPool != actualPgPool {
		test.Errorf("the expected pool and the actual pool are not equal")
	}
}
