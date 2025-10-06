package controllers

import (
	"context"
	"github.com/oklog/ulid/v2"
	"net/http"
	"workflowmanager/app/services"
)

type contextKey string

const (
	UserIdKey contextKey = "userId"
)

func AuthChain(authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			token, err := validateAuthorizationHeader(request)
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
			}
			claims, err := authService.ValidateToken(token)
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
			ctx := context.WithValue(request.Context(), UserIdKey, userId)
			next.ServeHTTP(responseWriter, request.WithContext(ctx))
		})
	}
}

func getUserId(request *http.Request) (ulid.ULID, bool) {
	userId, isContextKeyExist := request.Context().Value(UserIdKey).(ulid.ULID)
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
