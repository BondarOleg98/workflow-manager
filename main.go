package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"workflowmanager/app/db"
	route "workflowmanager/app/routing"
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
	route.InitBaseController()
	startServer()
	defer db.CloseDatabaseConnection()
}

func startDatabaseInstance() {
	go db.InitDatabaseInstance(db.CreatePgPool())
}

func startServer() {
	var err error
	address := fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))
	log.Printf("The app has started on the address %s", address)
	err = http.ListenAndServe(address, nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server is closed")
	} else if err != nil {
		log.Printf("Error during starting the server: %s", err)
		os.Exit(1)
	}
}

func checkIsProfileTypeDev() bool {
	const defaultProfile string = "dev"
	profile := flag.String("profile", defaultProfile, "application profile")
	flag.Parse()
	log.Printf("The profile name is: %s", *profile)
	return defaultProfile == *profile
}
