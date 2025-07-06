package routing

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseHandler(data any, responseWriter http.ResponseWriter) {
	responseData, err := json.Marshal(data)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	} else {
		responseWriter.WriteHeader(http.StatusOK)
		_, err = io.Writer.Write(responseWriter, responseData)
	}
}

func notFoundHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.WriteHeader(http.StatusNotFound)
	defaultNotFoundMessage := "The api request was not found"
	_, err := fmt.Fprint(responseWriter, defaultNotFoundMessage)
	if err != nil {
		log.Printf("%s", err.Error())
	}
}
