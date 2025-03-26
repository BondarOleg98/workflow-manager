package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Pool struct {
	Username     string
	DatabaseName string
}

// OpenDatabaseConnection TODO: add handling of the error
func OpenDatabaseConnection(pool Pool) *sql.DB {
	const baseConnectionUrl = "user=%s dbname=%s sslmode=verify-full"
	dbConnectionUrl := fmt.Sprintf(baseConnectionUrl, pool.Username, pool.DatabaseName)
	database, _ := sql.Open("postgres", dbConnectionUrl)
// 	if err != nil {
// 	}
	return database
}

// CloseDatabaseConnection TODO: remove an empty return
func CloseDatabaseConnection(database *sql.DB) {
	err := database.Close()
	if err != nil {
		return
	}
}
