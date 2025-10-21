package controllers

import (
	"context"
	"github.com/oklog/ulid/v2"
	"net/http"
	"workflowmanager/app/components/services"
	"workflowmanager/app/models"
)

type AuthChecker struct {
	authService *services.AuthService
}

func NewAuthChecker(authService *services.AuthService) *AuthChecker {
	return &AuthChecker{authService: authService}
}

func (authChecker *AuthChecker) checkAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		token, err := validateAuthorizationHeader(request)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, err := authChecker.authService.ValidateToken(token)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
			return
		}
		userIdStr, err := validateSubClaims(claims)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
			return
		}
		userId, err := ulid.Parse(userIdStr)
		if err != nil {
			http.Error(responseWriter, "Invalid userId in token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(request.Context(), models.UserIdKey, userId)
		handler.ServeHTTP(responseWriter, request.WithContext(ctx))
	})
}
