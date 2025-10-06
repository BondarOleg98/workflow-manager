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

func InitUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (userRepository *UserRepository) CreateUser(newUser models.User) (*models.User, error) {
	newUser.CreatedAt = time.Now()
	newUser.Id = ulid.Make()

	_, err := userRepository.database.Exec(
		queries.InsertUserQuery, newUser.Id.String(), newUser.Email, newUser.Username, newUser.Password, newUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (userRepository *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := userRepository.database.QueryRow(queries.SelectUserByEmail, email).
		Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.LastLogin,
		)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
