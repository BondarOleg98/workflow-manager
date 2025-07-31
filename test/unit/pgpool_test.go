package unit

import (
	"github.com/stretchr/testify/assert"
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
	assert.Equal(test, expectedPgPool, actualPgPool,
		"the expected pool and the actual pool are not equal")
}
