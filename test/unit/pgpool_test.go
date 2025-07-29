package unit

import (
	"testing"
	"workflowmanager/app/db"
	"workflowmanager/app/util"
)

func TestCreateNonSslPgPool(test *testing.T) {
	const configFilePath string = "../../app/resources/dev_env.yaml"
	err := util.LoadConfigs(configFilePath)
	if err != nil {
		test.FailNow()
	}
	actualPgPool := db.CreatePgPool()
	expectedPgPool := db.Pool{
		DriverName:    "postgres",
		ConnectionUrl: "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable",
	}
	if actualPgPool != expectedPgPool {
		test.Errorf("the expected pool and the actual pool are not equal")
	}
}
