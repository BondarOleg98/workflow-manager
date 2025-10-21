package services

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"workflowmanager/app/components/repository"
	"workflowmanager/app/models"
	"workflowmanager/app/util"
)

type AuthService struct {
	authRepository  *repository.AuthRepository
	refreshTokenTTL time.Duration
	accessTokenTTL  time.Duration
}

func NewAuthService(authRepository *repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository:  authRepository,
		refreshTokenTTL: util.ParseTimeConfigVariable(os.Getenv("REFRESH_TOKEN_TTL")),
		accessTokenTTL:  util.ParseTimeConfigVariable(os.Getenv("ACCESS_TOKEN_TTL")),
	}
}

func (authService *AuthService) RegisterUsingCredentials(registerRequest models.RegisterRequest) (*models.User, error) {
	retrievedUser, err := authService.authRepository.GetUserByEmail(registerRequest.Email)
	if retrievedUser != nil {
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
	createdUser, err := authService.authRepository.CreateUser(models.User{
		Email:    registerRequest.Email,
		Username: registerRequest.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (authService *AuthService) Login(loginRequest models.LoginRequest) (string, error) {
	retrievedUser, err := authService.authRepository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return "", models.ErrInvalidCredentials
	}
	if err := util.VerifyPassword(retrievedUser.Password, loginRequest.Password); err != nil {
		return "", models.ErrInvalidCredentials
	}
	accessToken, err := authService.generateAccessToken(retrievedUser)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (authService *AuthService) LoginWithRefreshToken(loginRequest models.LoginRequest) (string, string, error) {
	retrievedUser, err := authService.authRepository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return "", "", models.ErrInvalidCredentials
	}
	if err := util.VerifyPassword(retrievedUser.Password, loginRequest.Password); err != nil {
		return "", "", models.ErrInvalidCredentials
	}
	accessToken, err := authService.generateAccessToken(retrievedUser)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := authService.authRepository.
		CreateRefreshToken(retrievedUser.Id, authService.refreshTokenTTL)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken.Token, nil
}

func (authService *AuthService) generateAccessToken(user *models.User) (string, error) {
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
	retrievedRefreshToken, err := authService.authRepository.GetRefreshToken(refreshToken)
	if err != nil {
		return "", models.ErrInvalidToken
	}
	if retrievedRefreshToken.Revoked {
		return "", models.ErrInvalidToken
	}
	if time.Now().After(retrievedRefreshToken.ExpiredAt) {
		return "", models.ErrExpiredToken
	}
	retrievedUser, err := authService.authRepository.GetUserById(retrievedRefreshToken.UserId)
	if err != nil {
		return "", err
	}
	err = authService.authRepository.RevokeRefreshToken(retrievedRefreshToken.Token)
	if err != nil {
		return "", err
	}
	accessToken, err := authService.generateAccessToken(retrievedUser)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
