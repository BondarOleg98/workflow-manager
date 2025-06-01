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

	var secure string
	if os.Getenv("POSTGRES_HOST_AUTH_METHOD") != "trust" {
		secure = "disable"
	}
	return Pool{
	    DriverName: "postgres",
	    ConnectionUrl: fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
			username, password, databaseName, host, secure,
		),
	}
}
