package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"workflowmanager/app/models"
	"workflowmanager/app/services"
	"workflowmanager/app/util"
)

type AuthController struct {
	authService services.AuthService
}

func InitAuthController() AuthController {
	return AuthController{
		authService: services.InitAuthService(
			[]byte(os.Getenv("JWT_SECRET")),
			util.ParseTimeConfigVariable(os.Getenv("JWT_ACCESS_TTL"))),
	}
}

func (authController AuthController) AddAuthHandlers() {
	log.Println("Add the auth controller")
	baseAuthRoute := "/api/auth"
	http.HandleFunc(fmt.Sprintf("POST %s/register", baseAuthRoute), authController.registerUser)
	http.HandleFunc(fmt.Sprintf("POST %s/login", baseAuthRoute), authController.loginUser)
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
	createdUser, err := authController.authService.Register(registerRequest)
	if err != nil {
		if errors.Is(err, services.ErrEmailInUse) {
			http.Error(responseWriter, "the email already in use", http.StatusConflict)
			return
		}
		http.Error(responseWriter, "the error during creating a user", http.StatusInternalServerError)
		return
	}
	buildResponseBody(models.RegisterResponse{
		Id:       createdUser.Id.String(),
		Email:    createdUser.Email,
		Username: createdUser.Username,
	}, responseWriter)
	responseWriter.WriteHeader(http.StatusCreated)
}

func (authController AuthController) loginUser(
	responseWriter http.ResponseWriter, request *http.Request) {
	var loginRequest models.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(responseWriter, "the invalid request payload", http.StatusBadRequest)
		return
	}
	token, err := authController.authService.Login(loginRequest)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			http.Error(responseWriter, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(responseWriter, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	buildResponseBody(models.LoginResponse{Token: token}, responseWriter)
}
