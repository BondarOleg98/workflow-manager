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

func (postgresRefreshTokenRepository *PostgresRefreshTokenRepository) CreateRefreshToken(refreshToken models.RefreshToken) error {
	_, err := postgresRefreshTokenRepository.database.Exec(queries.InsertRefreshTokenQuery,
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

func (postgresRefreshTokenRepository *PostgresRefreshTokenRepository) GetRefreshTokenByValue(token string) (*models.RefreshToken, error) {
	var retrievedToken models.RefreshToken
	err := postgresRefreshTokenRepository.database.QueryRow(queries.GetRefreshTokenByValueQuery, token).Scan(
		&retrievedToken.Id, &retrievedToken.UserId,
		&retrievedToken.Token, &retrievedToken.ExpiredAt,
		&retrievedToken.CreatedAt, &retrievedToken.Revoked,
	)
	if err != nil {
		return nil, err
	}
	return &retrievedToken, nil
}

func (postgresRefreshTokenRepository *PostgresRefreshTokenRepository) GetRefreshTokenByUserId(userId string) (*models.RefreshToken, error) {
	var retrievedToken models.RefreshToken
	err := postgresRefreshTokenRepository.database.QueryRow(queries.GetRefreshTokenByUserIdQuery, userId).Scan(
		&retrievedToken.Id, &retrievedToken.UserId,
		&retrievedToken.Token, &retrievedToken.ExpiredAt,
		&retrievedToken.CreatedAt, &retrievedToken.Revoked,
	)
	if err != nil {
		return nil, err
	}
	return &retrievedToken, nil
}

func (postgresRefreshTokenRepository *PostgresRefreshTokenRepository) RevokeRefreshToken(refreshToken string) error {
	_, err := postgresRefreshTokenRepository.database.Exec(queries.RevokedRefreshTokenQuery, refreshToken)
	return err
}

func (postgresRefreshTokenRepository *PostgresRefreshTokenRepository) RemoveRefreshToken(refreshToken string) error {
	_, err := postgresRefreshTokenRepository.database.Exec(queries.RemoveRefreshTokenQuery, refreshToken)
	return err
}
