package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"workflowmanager/db"
	route "workflowmanager/routing"
	"workflowmanager/util"
)

func main() {
	var err error
	if checkIsProfileTypeDev() {
		err = util.LoadConfigs()
	}
	if err != nil {
		return
	}
	pgPool := db.CreatePgPool()
	go db.InitDatabaseInstance(pgPool)

	route.WorkflowEndpoints{}.BaseController()

	address := fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))
	err = http.ListenAndServe(address, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func checkIsProfileTypeDev() bool {
	const defaultProfile string = "dev"
	profile := flag.String("profile", defaultProfile, "application profile")
	flag.Parse()
	return defaultProfile == *profile
}
