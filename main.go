package main

import (
	route "./routing"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	route.WorkflowEndpoints{}.BaseController()
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
