package repository

import (
	"database/sql"
	"github.com/oklog/ulid/v2"
	"time"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type PostgresRefreshTokenRepository struct {
	database *sql.DB
}

func NewPostgresRefreshTokenRepository(database *sql.DB) *PostgresRefreshTokenRepository {
	return &PostgresRefreshTokenRepository{database: database}
}

func (PostgresRefreshTokenRepository *PostgresRefreshTokenRepository) CreateRefreshToken(
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
	_, err := PostgresRefreshTokenRepository.database.Exec(queries.InsertRefreshTokenQuery,
		token.Id.String(), token.UserId.String(), token.Token, token.ExpiredAt, token.CreatedAt, token.Revoked)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (PostgresRefreshTokenRepository *PostgresRefreshTokenRepository) GetRefreshToken(token string) (*models.RefreshToken, error) {
	var retrievedToken models.RefreshToken
	err := PostgresRefreshTokenRepository.database.QueryRow(queries.GetRefreshTokenQuery, token).Scan(
		&retrievedToken.Id, &retrievedToken.UserId,
		&retrievedToken.Token, &retrievedToken.ExpiredAt,
		&retrievedToken.CreatedAt, &retrievedToken.Revoked,
	)
	if err != nil {
		return nil, err
	}
	return &retrievedToken, nil
}

func (PostgresRefreshTokenRepository *PostgresRefreshTokenRepository) RevokeRefreshToken(refreshToken string) error {
	_, err := PostgresRefreshTokenRepository.database.Exec(queries.RevokedRefreshTokenQuery, refreshToken)
	return err
}
