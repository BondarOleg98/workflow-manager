package repository

import (
	"database/sql"
	"github.com/oklog/ulid/v2"
	"time"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type TokenRepository struct {
	database *sql.DB
}

func NewTokenRepository(database *sql.DB) *TokenRepository {
	return &TokenRepository{database: database}
}

func (tokenRepository *TokenRepository) CreateRefreshToken(
	userId ulid.ULID, ttl time.Duration) (*models.RefreshToken, error) {
	refreshTokenId := ulid.Make()
	expiresAt := time.Now().Add(ttl)

	token := &models.RefreshToken{
		Id:        refreshTokenId,
		UserId:    userId,
		Token:     refreshTokenId.String(),
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
		Revoked:   false,
	}

	_, err := tokenRepository.database.Exec(queries.InsertRefreshToken,
		token.Id,
		token.UserId,
		token.Token,
		token.ExpiresAt,
		token.CreatedAt,
		token.Revoked)
	if err != nil {
		return nil, err
	}

	return token, nil
}
