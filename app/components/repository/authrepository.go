package repository

import (
	"database/sql"
	"github.com/oklog/ulid/v2"
	"time"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type AuthRepository struct {
	database *sql.DB
}

func NewAuthRepository(database *sql.DB) *AuthRepository {
	return &AuthRepository{database: database}
}

func (authRepository *AuthRepository) CreateUser(newUser models.User) (*models.User, error) {
	newUser.CreatedAt = time.Now()
	newUser.Id = ulid.Make()

	_, err := authRepository.database.Exec(
		queries.SaveUserQuery, newUser.Id.String(), newUser.Email, newUser.Username, newUser.Password, newUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (authRepository *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	var lastLogin sql.NullTime

	err := authRepository.database.QueryRow(queries.GetUserByEmailQuery, email).
		Scan(
			&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.LastLogin,
		)
	if err != nil {
		return nil, err
	}

	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}

func (authRepository *AuthRepository) GetUserById(id ulid.ULID) (*models.User, error) {
	var user models.User
	var lastLogin sql.NullTime
	err := authRepository.database.QueryRow(queries.GetUserByIdQuery, id.String()).Scan(
		&user.Id, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &lastLogin,
	)
	if err != nil {
		return nil, err
	}
	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}
	return &user, nil
}

func (authRepository *AuthRepository) CreateRefreshToken(
	userId ulid.ULID, ttl time.Duration) (*models.RefreshToken, error) {
	refreshTokenId := ulid.Make()
	expiresAt := time.Now().Add(ttl)

	token := &models.RefreshToken{
		Id:        refreshTokenId,
		UserId:    userId,
		Token:     refreshTokenId.String(),
		ExpiredAt: expiresAt,
		CreatedAt: time.Now(),
		Revoked:   false,
	}

	_, err := authRepository.database.Exec(queries.InsertRefreshTokenQuery,
		token.Id.String(), token.UserId.String(), token.Token, token.ExpiredAt, token.CreatedAt, token.Revoked)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (authRepository *AuthRepository) GetRefreshToken(token string) (*models.RefreshToken, error) {
	var retrievedToken models.RefreshToken
	err := authRepository.database.QueryRow(queries.GetRefreshTokenQuery, token).Scan(
		&retrievedToken.Id, &retrievedToken.UserId,
		&retrievedToken.Token, &retrievedToken.ExpiredAt,
		&retrievedToken.CreatedAt, &retrievedToken.Revoked,
	)
	if err != nil {
		return nil, err
	}
	return &retrievedToken, nil
}

func (authRepository *AuthRepository) RevokeRefreshToken(refreshToken string) error {
	_, err := authRepository.database.Exec(queries.RevokedRefreshTokenQuery, refreshToken)
	return err
}
