package controllers

import (
	"encoding/json"
	"github.com/oklog/ulid/v2"
	"net/http"
	"workflowmanager/app/models"
)

func buildResponseBody(data any, responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(responseWriter).Encode(data)
}

func getUserId(request *http.Request) (ulid.ULID, bool) {
	userId, isContextKeyExist := request.Context().Value(models.UserIdKey).(ulid.ULID)
	return userId, isContextKeyExist
}

func isAuthorised(
	responseWriter http.ResponseWriter, request *http.Request) {
	_, isContextKeyExist := getUserId(request)
	if !isContextKeyExist {
		http.Error(responseWriter, "Unauthorized", http.StatusUnauthorized)
		return
	}
}
