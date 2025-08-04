package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var onceInitDatabase sync.Once
var database *sql.DB

func InitDatabaseInstance(pool Pool) *sql.DB {
	if database == nil {
		onceInitDatabase.Do(
			func() {
				var err error
				database, err = sql.Open(pool.DriverName, pool.ConnectionUrl)
				if err != nil {
					log.Fatalf("The error during initialization DB's instance: %s", err)
				}
			})
	}
	log.Println("The DB's instance was initialized")
	return database
}

func GetDatabaseInstance() *sql.DB {
	log.Println("Getting the db instance")
	return database
}

func CloseDatabaseConnection() error {
	err := database.Close()
	if err != nil {
		return err
	}
	return nil
}
