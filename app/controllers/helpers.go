package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func responseComposer(data any, responseWriter http.ResponseWriter) {
	responseData, err := json.Marshal(data)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	} else {
		_, err = io.Writer.Write(responseWriter, responseData)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
		}
	}
}
