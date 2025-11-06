package unit

import (
	"database/sql"
	"fmt"
	"github.com/oklog/ulid/v2"
	"workflowmanager/app/models"
)

type MockUserRepository struct {
	users []models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (mockUserRepository *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	for _, retrievedUser := range mockUserRepository.users {
		if retrievedUser.Email == email {
			return &retrievedUser, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (mockUserRepository *MockUserRepository) GetUserById(id ulid.ULID) (*models.User, error) {
	for _, retrievedUser := range mockUserRepository.users {
		if retrievedUser.Id == id {
			return &retrievedUser, nil
		}
	}
	return nil, fmt.Errorf("the user with id: %s haven't found", id.String())
}

func (mockUserRepository *MockUserRepository) CreateUser(newUser models.User) (*models.User, error) {
	mockUserRepository.users = append(mockUserRepository.users, newUser)
	return &newUser, nil
}
