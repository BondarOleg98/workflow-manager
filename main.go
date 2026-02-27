package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"workflowmanager/app/components"
	"workflowmanager/app/db"
	"workflowmanager/app/util"
)

func main() {
	var err error
	if checkIsProfileTypeDev() {
		const configFilePath string = "app/resources/dev_env.yaml"
		err = util.LoadConfigs(configFilePath)
	}
	if err != nil {
		return
	}
	setLogging()
	startDatabaseInstance()
	components.InitAppComponents()
	startServer()
	defer db.CloseDatabaseConnection()
}

func startDatabaseInstance() {
	db.InitDatabaseInstance(db.CreatePgPool())
}

func startServer() {
	var err error
	address := fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))
	slog.Info("The app has started", "address", address)
	err = http.ListenAndServe(address, nil)

	if errors.Is(err, http.ErrServerClosed) {
		slog.Error("Server is closed")
	} else if err != nil {
		slog.Error("Error during starting the server", "err", err)
		os.Exit(1)
	}
}

func checkIsProfileTypeDev() bool {
	const defaultProfile string = "dev"
	profile := flag.String("profile", defaultProfile, "application profile")
	flag.Parse()
	slog.Info("The app", "profile", *profile)
	return defaultProfile == *profile
}

func setLogging() {
	var logLevel slog.LevelVar
	defaultEnvLevel := "info"
	envLevel := os.Getenv("LOG_LEVEL")
	if envLevel == "" {
		envLevel = defaultEnvLevel
	}
	if err := logLevel.UnmarshalText([]byte(strings.ToLower(envLevel))); err != nil {
		slog.Error("The invalid LOG_LEVEL environment variable", "err", err)
		logLevel.Set(slog.LevelInfo)
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: &logLevel,
	}))
	slog.SetDefault(logger)
}
