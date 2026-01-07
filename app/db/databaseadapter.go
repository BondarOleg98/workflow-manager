package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"sync"
)

var (
	database         *sql.DB
	onceInitDatabase sync.Once
)

func InitDatabaseInstance(pool Pool) *sql.DB {
	if database == nil {
		onceInitDatabase.Do(
			func() {
				var err error
				database, err = sql.Open(pool.DriverName, pool.ConnectionUrl)
				if err != nil {
					log.Fatalf("The error during initialization DB's instance: %v", err)
				}
				slog.Info("The DB's instance was initialized")
			})
	}
	return database
}

func GetDatabaseInstance() *sql.DB {
	slog.Info("Getting the db instance")
	return database
}

func CloseDatabaseConnection() {
	if database != nil {
		if err := database.Close(); err != nil {
			slog.Error("The error during closing DB's instance:", "err", err)
		} else {
			database = nil
		}
	}
}
