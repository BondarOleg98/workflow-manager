package security

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"log/slog"
	"net/http"
	"strings"
	"workflowmanager/app/components/services"
)

type contextKey string

const UserIdKey contextKey = "userId"

type PreAuthorize struct {
	authService *services.AuthService
}

func NewPreAuthorize(authService *services.AuthService) *PreAuthorize {
	return &PreAuthorize{authService: authService}
}

func (preAuthorize *PreAuthorize) SecurityFilterChain(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		token, err := preAuthorize.validateAuthorizationHeader(request)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, err := preAuthorize.authService.ValidateToken(token)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
			return
		}
		userIdClaim, err := preAuthorize.validateSubClaims(claims)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusUnauthorized)
			return
		}
		userId, err := ulid.Parse(userIdClaim)
		if err != nil {
			http.Error(responseWriter, "Invalid userId in token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(request.Context(), UserIdKey, userId)
		handler.ServeHTTP(responseWriter, request.WithContext(ctx))
	})
}

func (preAuthorize *PreAuthorize) IsAuthorised(
	responseWriter http.ResponseWriter, request *http.Request) {
	_, isContextKeyExist := preAuthorize.getUserId(request)
	if !isContextKeyExist {
		http.Error(responseWriter, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (preAuthorize *PreAuthorize) validateAuthorizationHeader(request *http.Request) (string, error) {
	const tokenCountParts int = 2
	const tokenType string = "Bearer"
	const separatedSymbol string = " "
	slog.Info("validating an authorization header")
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

func (preAuthorize *PreAuthorize) validateSubClaims(claims jwt.MapClaims) (string, error) {
	userId, isContextKeyExist := claims["sub"].(string)
	slog.Info("validating claims")
	if !isContextKeyExist {
		return "", errors.New("invalid token claims")
	}
	return userId, nil
}

func (preAuthorize *PreAuthorize) getUserId(request *http.Request) (ulid.ULID, bool) {
	userId, isContextKeyExist := request.Context().Value(UserIdKey).(ulid.ULID)
	return userId, isContextKeyExist
}
