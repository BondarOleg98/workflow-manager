package routing

import (
	"encoding/json"
	"errors"
	"net/http"
	"workflowmanager/app/auth/models"
	"workflowmanager/app/auth/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func (authHandler *AuthHandler) RegisterController(responseWriter http.ResponseWriter, request *http.Request) {
	var registerRequest models.RegisterRequest
	if err := json.NewDecoder(request.Body).Decode(&registerRequest); err != nil {
		http.Error(responseWriter, "The invalid request payload", http.StatusBadRequest)
		return
	}
	if registerRequest.Email == "" ||
		registerRequest.Username == "" ||
		registerRequest.Password == "" {
		http.Error(responseWriter, "email, username, and password are required", http.StatusBadRequest)
		return
	}
	newUser := models.User{
		Email:    registerRequest.Email,
		Username: registerRequest.Username,
		Password: registerRequest.Password,
	}
	createdUser, err := authHandler.authService.Register(newUser)
	if err != nil {
		if errors.Is(err, services.ErrEmailInUse) {
			http.Error(responseWriter, "the email already in use", http.StatusConflict)
			return
		}
		http.Error(responseWriter, "the error during creating a user", http.StatusInternalServerError)
		return
	}

	response := models.RegisterResponse{
		ID:       createdUser.ID.String(),
		Email:    createdUser.Email,
		Username: createdUser.Username,
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(responseWriter).Encode(response)
	if err != nil {
		return
	}
}
