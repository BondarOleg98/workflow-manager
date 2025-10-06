package controllers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"workflowmanager/app/models"
)

func validateAuthRequestBody(registerRequest models.RegisterRequest) error {
	if registerRequest.Email == "" ||
		registerRequest.Username == "" ||
		registerRequest.Password == "" {
		return errors.New("email, username, and password are required")
	}
	return nil
}

func validateAuthorizationHeader(request *http.Request) (string, error) {
	const tokenCountParts int = 2
	const tokenType string = "Bearer"
	const separatedSymbol string = " "
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header required")
	}
	tokenParts := strings.Split(authHeader, separatedSymbol)
	if len(tokenParts) != tokenCountParts || tokenParts[0] != tokenType {
		return "", errors.New("invalid authorization format")
	}
	token := tokenParts[1]
	return token, nil
}

func validateSubClaims(claims jwt.MapClaims) (string, error) {
	userId, isContextKeyExist := claims["sub"].(string)
	if !isContextKeyExist {
		return "", errors.New("invalid token claims")
	}
	return userId, nil
}
