package repository

import (
	"database/sql"
	"github.com/oklog/ulid/v2"
	"time"
	"workflowmanager/app/db/queries"
	"workflowmanager/app/models"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (userRepository *UserRepository) CreateUser(newUser models.User) (*models.User, error) {
	newUser.CreatedAt = time.Now()
	newUser.Id = ulid.Make()
	_, err := userRepository.database.Exec(
		queries.SaveUserQuery, newUser.Id.String(), newUser.Email, newUser.Username, newUser.Password, newUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (userRepository *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	var lastLogin sql.NullTime
	err := userRepository.database.QueryRow(queries.GetUserByEmailQuery, email).
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

func (userRepository *UserRepository) GetUserById(id ulid.ULID) (*models.User, error) {
	var user models.User
	var lastLogin sql.NullTime
	err := userRepository.database.QueryRow(queries.GetUserByIdQuery, id.String()).Scan(
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
