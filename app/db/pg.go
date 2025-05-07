package db

import (
	"fmt"
	"os"
)

func CreatePgPool() Pool {
	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	databaseName := os.Getenv("PG_HOST")
	host := os.Getenv("PG_DATABASE_NAME")
	return Pool{
		Username:     username,
		Password:     password,
		DatabaseName: databaseName,
		Host:         host,
		ConnectionUrl: fmt.Sprintf("postgres://%s:%s@%s/%s",
			username,
			password,
			databaseName,
			host,
		),
	}
}
