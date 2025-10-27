package repository

import (
	"github.com/oklog/ulid/v2"
	"workflowmanager/app/models"
)

type UserRepository interface {
	CreateUser(newUser models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id ulid.ULID) (*models.User, error)
}
