package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

var onceInitDatabase sync.Once
var database *sql.DB

func InitDatabaseInstance(pool Pool) *sql.DB {
	if database == nil {
		onceInitDatabase.Do(
			func() {
				database, _ = sql.Open(pool.Username, pool.ConnectionUrl)
			})
	}
	return database
}

func GetDatabaseInstance() *sql.DB {
	return database
}

func CloseDatabaseConnection() error {
	err := database.Close()
	if err != nil {
		return err
	}
	return nil
}
