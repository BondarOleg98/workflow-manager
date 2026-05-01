package integration

import (
	"database/sql"
	"testing"
	"workflowmanager/app/db"
)

func TestCorrectDatabaseInstance(test *testing.T) {
	pool := db.Pool{
		DriverName:    "postgres",
		ConnectionUrl: "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable",
	}
	db.InitDatabaseInstance(pool)
	defer db.CloseDatabaseConnection()

	databaseActualName, _ := getActualDatabaseNameThroughAdapter()
	databaseExpectedName, _ := getExpectedDatabaseName()
	if databaseExpectedName != databaseActualName {
		test.Errorf(
			"the DBs are not equals: actual - %s, expected - %s",
			databaseActualName, databaseExpectedName)
	}
}

func getActualDatabaseNameThroughAdapter() (string, error) {
	var databaseActualName string
	databaseActual := db.GetDatabaseInstance()
	err := databaseActual.QueryRow("SELECT current_database()").Scan(&databaseActualName)
	if err != nil {
		return "", err
	}
	return databaseActualName, nil
}

func getExpectedDatabaseName() (string, error) {
	var databaseExpectedName string
	driverName := "postgres"
	connectionUrl := "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable"
	databaseExpected, _ := sql.Open(driverName, connectionUrl)
	defer func(databaseExpected *sql.DB) {
		err := databaseExpected.Close()
		if err != nil {
			return
		}
	}(databaseExpected)
	err := databaseExpected.QueryRow("SELECT current_database()").Scan(&databaseExpectedName)
	if err != nil {
		return "", err
	}
	return databaseExpectedName, nil
}
