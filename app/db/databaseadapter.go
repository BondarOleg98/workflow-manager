package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
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
				log.Println("The DB's instance was initialized")
			})
	}
	return database
}

func GetDatabaseInstance() *sql.DB {
	log.Println("Getting the db instance")
	return database
}

func CloseDatabaseConnection() {
	if database != nil {
		if err := database.Close(); err != nil {
			log.Printf("The error during closing DB's instance: %v", err)
		} else {
			database = nil
		}
	}
}
