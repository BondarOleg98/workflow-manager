package repository

import (
	"github.com/google/uuid"
	"workflowmanager/app/models"
)

type UserRepository interface {
	CreateUser(newUser models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uuid.UUID) (*models.User, error)
}
