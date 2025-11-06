package repository

import (
	"database/sql"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type PostgresRefreshTokenRepository struct {
	database *sql.DB
}

func NewPostgresRefreshTokenRepository(database *sql.DB) *PostgresRefreshTokenRepository {
	return &PostgresRefreshTokenRepository{database: database}
}

func (PostgresRefreshTokenRepository *PostgresRefreshTokenRepository) CreateRefreshToken(refreshToken models.RefreshToken) error {
	_, err := PostgresRefreshTokenRepository.database.Exec(queries.InsertRefreshTokenQuery,
		refreshToken.Id.String(),
		refreshToken.UserId.String(),
		refreshToken.Token,
		refreshToken.ExpiredAt,
		refreshToken.CreatedAt,
		refreshToken.Revoked)
	if err != nil {
		return err
	}
	return nil
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
