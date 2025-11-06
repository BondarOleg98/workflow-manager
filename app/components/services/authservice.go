package services

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"log/slog"
	"os"
	"time"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/models"
	"workflowmanager/app/util"
)

type AuthService struct {
	userRepository         repository.UserRepository
	refreshTokenRepository repository.RefreshTokenRepository
	refreshTokenTTL        time.Duration
	accessTokenTTL         time.Duration
}

func NewAuthService(userRepository repository.UserRepository,
	refreshTokenRepository repository.RefreshTokenRepository) *AuthService {
	return &AuthService{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
		refreshTokenTTL:        util.ParseTimeConfigVariable(os.Getenv("REFRESH_TOKEN_TTL")),
		accessTokenTTL:         util.ParseTimeConfigVariable(os.Getenv("ACCESS_TOKEN_TTL")),
	}
}

func (authService *AuthService) RegisterUsingCredentials(registerRequest models.RegisterRequest) (*models.User, error) {
	retrievedUser, err := authService.userRepository.GetUserByEmail(registerRequest.Email)
	if retrievedUser != nil {
		slog.Info("retrieved the user by", "email", retrievedUser.Email)
		return nil, models.ErrEmailInUse
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	hashedPassword, err := util.HashPassword(registerRequest.Password)
	if err != nil {
		return nil, err
	}
	registerRequest.Password = hashedPassword
	createdUser, err := authService.userRepository.CreateUser(models.User{
		Email:     registerRequest.Email,
		Username:  registerRequest.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		Id:        ulid.Make(),
	})
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (authService *AuthService) Login(loginRequest models.LoginRequest) (string, string, error) {
	retrievedUser, err := authService.userRepository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return "", "", models.ErrInvalidCredentials
	}
	slog.Info("retrieved the user by", "email", retrievedUser.Email)
	if err := util.VerifyPassword(retrievedUser.Password, loginRequest.Password); err != nil {
		return "", "", models.ErrInvalidCredentials
	}
	accessToken, err := authService.generateAccessToken(retrievedUser)
	if err != nil {
		return "", "", err
	}
	refreshTokenId := ulid.Make()
	expiresAt := time.Now().Add(authService.refreshTokenTTL)
	refreshToken := models.RefreshToken{
		Id:        refreshTokenId,
		UserId:    retrievedUser.Id,
		Token:     refreshTokenId.String(),
		ExpiredAt: expiresAt,
		CreatedAt: time.Now(),
		Revoked:   false,
	}
	err = authService.refreshTokenRepository.CreateRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken.Token, nil
}

func (authService *AuthService) generateAccessToken(user *models.User) (string, error) {
	slog.Info("generating the access token")
	expirationTime := time.Now().Add(authService.accessTokenTTL)
	claims := jwt.MapClaims{
		"sub":      user.Id.String(),
		"username": user.Username,
		"email":    user.Email,
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (authService *AuthService) ValidateToken(token string) (jwt.MapClaims, error) {
	slog.Info("validating the access token")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, models.ErrInvalidToken
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, models.ErrExpiredToken
		}
		return nil, models.ErrInvalidToken
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, models.ErrInvalidToken
}

func (authService *AuthService) RefreshAccessToken(refreshToken string) (string, error) {
	retrievedRefreshToken, err := authService.refreshTokenRepository.GetRefreshToken(refreshToken)
	if err != nil {
		return "", models.ErrInvalidToken
	}
	if retrievedRefreshToken.Revoked {
		return "", models.ErrInvalidToken
	}
	if time.Now().After(retrievedRefreshToken.ExpiredAt) {
		return "", models.ErrExpiredToken
	}
	retrievedUser, err := authService.userRepository.GetUserById(retrievedRefreshToken.UserId)
	if err != nil {
		return "", err
	}
	slog.Info("retrieved the user by", "id", retrievedUser.Id)
	err = authService.refreshTokenRepository.RevokeRefreshToken(retrievedRefreshToken.Token)
	if err != nil {
		return "", err
	}
	accessToken, err := authService.generateAccessToken(retrievedUser)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
