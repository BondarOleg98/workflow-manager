package db

import (
	"fmt"
	"log"
	"os"
)

func CreatePgPool() Pool {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	databaseName := os.Getenv("POSTGRES_DB")

	secure := "disable"
	postgresAuthMethod := os.Getenv("POSTGRES_HOST_AUTH_METHOD")
	if postgresAuthMethod != "" {
		secure = postgresAuthMethod
	}
	log.Printf("The DB ssl is: %s", secure)

	return Pool{
		DriverName: "postgres",
		ConnectionUrl: fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
			username, password, host, databaseName, secure,
		),
	}
}
