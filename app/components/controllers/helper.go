package controllers

import (
	"encoding/json"
	"net/http"
)

func buildResponseBody(data any, responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(responseWriter).Encode(data)
}
