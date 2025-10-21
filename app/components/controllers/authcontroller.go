package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"workflowmanager/app/components/services"
	"workflowmanager/app/models"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (authController AuthController) AddAuthHandlers() {
	log.Println("Add the auth controller")
	baseAuthRoute := "/api/auth"
	http.HandleFunc(fmt.Sprintf("POST %s/register", baseAuthRoute), authController.registerUser)
	http.HandleFunc(fmt.Sprintf("POST %s/login", baseAuthRoute), authController.loginUser)
	http.HandleFunc(fmt.Sprintf("POST %s/refresh", baseAuthRoute), authController.refreshToken)
}

func (authController AuthController) registerUser(
	responseWriter http.ResponseWriter, request *http.Request) {
	var registerRequest models.RegisterRequest
	var err error
	if err = json.NewDecoder(request.Body).Decode(&registerRequest); err != nil {
		http.Error(responseWriter, "the invalid request payload", http.StatusBadRequest)
		return
	}
	if err = validateAuthRequestBody(registerRequest); err != nil {
		http.Error(responseWriter, "email, username, and password are required", http.StatusBadRequest)
		return
	}
	createdUser, err := authController.authService.RegisterUsingCredentials(registerRequest)
	if err != nil {
		if errors.Is(err, models.ErrEmailInUse) {
			http.Error(responseWriter, "the email already in use", http.StatusConflict)
			return
		}
		http.Error(responseWriter, "the error during creating a user", http.StatusInternalServerError)
		return
	}
	responseWriter.WriteHeader(http.StatusCreated)
	buildResponseBody(models.RegisterResponse{
		Id:       createdUser.Id.String(),
		Email:    createdUser.Email,
		Username: createdUser.Username,
	}, responseWriter)
}

func (authController AuthController) loginUser(
	responseWriter http.ResponseWriter, request *http.Request) {
	var loginRequest models.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(responseWriter, "the invalid request payload", http.StatusBadRequest)
		return
	}
	accessToken, refreshToken, err := authController.authService.LoginWithRefreshToken(loginRequest)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			http.Error(responseWriter, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(responseWriter, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	buildResponseBody(models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, responseWriter)
}

func (authController AuthController) refreshToken(
	responseWriter http.ResponseWriter, request *http.Request) {
	var req models.RefreshRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(responseWriter, "the invalid request payload", http.StatusBadRequest)
		return
	}
	accessToken, err := authController.authService.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		if errors.Is(err, models.ErrInvalidToken) ||
			errors.Is(err, models.ErrExpiredToken) {
			http.Error(responseWriter, "the invalid or expired refresh token", http.StatusUnauthorized)
		} else {
			http.Error(responseWriter, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	buildResponseBody(models.RefreshResponse{AccessToken: accessToken}, responseWriter)
}
