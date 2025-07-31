package integration

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"workflowmanager/app/db"
)

func TestInitExistedDatabaseInstance(test *testing.T) {
	const driverName string = "postgres"
	const connectionExistDBUrl string = "postgres://postgres:postgres@localhost/workflow_manager?sslmode=disable"
	databaseActualName, _ := getActualDatabaseNameThroughAdapter(driverName, connectionExistDBUrl)
	databaseExpectedName, _ := getExpectedDatabaseName(driverName, connectionExistDBUrl)
	assert.EqualValues(test, databaseActualName, databaseExpectedName,
		"the DBs are not equals")
}

func TestInitNonExistedDatabaseInstance(test *testing.T) {
	const driverName string = "postgres"
	const connectionNonExistDBUrl string = "postgres://postgres:postgres@localhost/noexist?sslmode=disable"
	_, err := getActualDatabaseNameThroughAdapter(driverName, connectionNonExistDBUrl)
	if err == nil {
		test.Errorf("the DB is exist even the connection url is incorrect: %s", err)
	}
}

func getActualDatabaseNameThroughAdapter(driverName string, connectionUrl string) (string, error) {
	var databaseActualName string
	pool := db.Pool{
		DriverName:    driverName,
		ConnectionUrl: connectionUrl,
	}
	databaseActual := db.InitDatabaseInstance(pool)
	err := databaseActual.QueryRow("SELECT current_database()").Scan(&databaseActualName)
	if err != nil {
		return "", err
	}
	defer func(databaseActual *sql.DB) {
		err := databaseActual.Close()
		if err != nil {
			return
		}
	}(databaseActual)
	return databaseActualName, nil
}

func getExpectedDatabaseName(driverName string, connectionUrl string) (string, error) {
	var databaseExpectedName string
	databaseExpected, _ := sql.Open(driverName, connectionUrl)
	err := databaseExpected.QueryRow("SELECT current_database()").Scan(&databaseExpectedName)
	if err != nil {
		return "", err
	}
	defer func(databaseExpected *sql.DB) {
		err := databaseExpected.Close()
		if err != nil {
			return
		}
	}(databaseExpected)
	return databaseExpectedName, nil
}
