package repository

import (
	"database/sql"
	"github.com/oklog/ulid/v2"
	"time"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type RefreshTokenRepository struct {
	database *sql.DB
}

func NewRefreshTokenRepository(database *sql.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{database: database}
}

func (refreshTokenRepository *RefreshTokenRepository) CreateRefreshToken(
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
	_, err := refreshTokenRepository.database.Exec(queries.InsertRefreshTokenQuery,
		token.Id.String(), token.UserId.String(), token.Token, token.ExpiredAt, token.CreatedAt, token.Revoked)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (refreshTokenRepository *RefreshTokenRepository) GetRefreshToken(token string) (*models.RefreshToken, error) {
	var retrievedToken models.RefreshToken
	err := refreshTokenRepository.database.QueryRow(queries.GetRefreshTokenQuery, token).Scan(
		&retrievedToken.Id, &retrievedToken.UserId,
		&retrievedToken.Token, &retrievedToken.ExpiredAt,
		&retrievedToken.CreatedAt, &retrievedToken.Revoked,
	)
	if err != nil {
		return nil, err
	}
	return &retrievedToken, nil
}

func (refreshTokenRepository *RefreshTokenRepository) RevokeRefreshToken(refreshToken string) error {
	_, err := refreshTokenRepository.database.Exec(queries.RevokedRefreshTokenQuery, refreshToken)
	return err
}
