package repository

import (
	"database/sql"
	"github.com/oklog/ulid/v2"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type PostgresUserRepository struct {
	database *sql.DB
}

func NewPostgresUserRepository(database *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{database: database}
}

func (postgresUserRepository *PostgresUserRepository) CreateUser(newUser models.User) (*models.User, error) {
	_, err := postgresUserRepository.database.Exec(
		queries.SaveUserQuery, newUser.Id.String(), newUser.Email, newUser.Username, newUser.Password, newUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (postgresUserRepository *PostgresUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	var lastLogin sql.NullTime
	err := postgresUserRepository.database.QueryRow(queries.GetUserByEmailQuery, email).
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

func (postgresUserRepository *PostgresUserRepository) GetUserById(id ulid.ULID) (*models.User, error) {
	var user models.User
	var lastLogin sql.NullTime
	err := postgresUserRepository.database.QueryRow(queries.GetUserByIdQuery, id.String()).Scan(
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
