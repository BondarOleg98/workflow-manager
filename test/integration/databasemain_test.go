package integration

import (
	"testing"
	"workflowmanager/app/db"
)

func TestDatabaseMain(test *testing.T) {
	pool := db.Pool{
		DriverName:    "postgres",
		ConnectionUrl: "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable",
	}
	db.InitDatabaseInstance(pool)
	defer db.CloseDatabaseConnection()

	test.Run("TestCrdOperations", func(test *testing.T) {
		testCrdOperations(test)
	})
	test.Run("TestCorrectDatabaseInstance", func(t *testing.T) {
		testCorrectDatabaseInstance(test)
	})
}
