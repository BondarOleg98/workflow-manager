package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"workflowmanager/app/components/services"
	"workflowmanager/app/models"
	"workflowmanager/app/models/requestmodels"
	"workflowmanager/app/models/responsemodels"
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
	var registerRequest requestmodels.RegisterRequest
	var err error
	if err = json.NewDecoder(request.Body).Decode(&registerRequest); err != nil {
		http.Error(responseWriter, "the invalid request payload", http.StatusBadRequest)
		return
	}
	if registerRequest.Email == "" || registerRequest.Username == "" || registerRequest.Password == "" {
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
	buildResponseBody(responsemodels.RegisterResponse{
		Email:    createdUser.Email,
		Username: createdUser.Username,
	}, responseWriter)
}

func (authController AuthController) loginUser(
	responseWriter http.ResponseWriter, request *http.Request) {
	var loginRequest requestmodels.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(responseWriter, "the invalid request payload", http.StatusBadRequest)
		return
	}
	accessToken, refreshToken, err := authController.authService.Login(loginRequest)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			http.Error(responseWriter, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(responseWriter, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	buildResponseBody(responsemodels.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, responseWriter)
}

func (authController AuthController) refreshToken(
	responseWriter http.ResponseWriter, request *http.Request) {
	var req requestmodels.RefreshRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(responseWriter, "the invalid requestmodels payload", http.StatusBadRequest)
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
	buildResponseBody(responsemodels.RefreshResponse{AccessToken: accessToken}, responseWriter)
}
