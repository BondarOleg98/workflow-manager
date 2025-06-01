package db

import (
	"fmt"
	"os"
)

func CreatePgPool() Pool {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_HOST")
	host := os.Getenv("POSTGRES_DB")
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
