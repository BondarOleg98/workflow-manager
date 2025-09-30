package services

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"workflowmanager/app/db"
	"workflowmanager/app/models"
	"workflowmanager/app/repository"
	"workflowmanager/app/util"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("token has expired")
	ErrEmailInUse         = errors.New("email already in use")
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtSecret      []byte
	accessTokenTTL time.Duration
}

func InitAuthService(
	jwtSecret []byte,
	accessTokenTTL time.Duration) AuthService {
	return AuthService{
		userRepository: repository.InitUserRepository(db.GetDatabaseInstance()),
		jwtSecret:      jwtSecret,
		accessTokenTTL: accessTokenTTL,
	}
}

func (authService *AuthService) Register(user models.User) (*models.User, error) {
	_, err := authService.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, ErrEmailInUse
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	createdUser, err := authService.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (authService *AuthService) Login(user models.User) (string, error) {
	retrievedUser, err := authService.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	if err := util.VerifyPassword(retrievedUser.Password, user.Password); err != nil {
		return "", ErrInvalidCredentials
	}
	jwtToken, err := authService.generateAccessToken(retrievedUser)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (authService *AuthService) generateAccessToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(authService.accessTokenTTL)
	claims := jwt.MapClaims{
		"sub":      user.ID.String(),
		"username": user.Username,
		"email":    user.Email,
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(authService.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (authService *AuthService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, ErrInvalidToken
		}
		return authService.jwtSecret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}
