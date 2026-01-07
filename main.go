package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
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
	slog.Info("The app has started on the", "address", address)
	err = http.ListenAndServe(address, nil)

	if errors.Is(err, http.ErrServerClosed) {
		slog.Error("Server is closed")
	} else if err != nil {
		slog.Error("Error during starting the server: %s", err)
		os.Exit(1)
	}
}

func checkIsProfileTypeDev() bool {
	const defaultProfile string = "dev"
	profile := flag.String("profile", defaultProfile, "application profile")
	flag.Parse()
	slog.Info("The profile name is: %s", *profile)
	return defaultProfile == *profile
}
