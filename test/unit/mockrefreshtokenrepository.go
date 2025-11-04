package unit

import (
	"fmt"
	"workflowmanager/app/models"
)

type MockRefreshTokenRepository struct {
	refreshTokens []models.RefreshToken
}

func NewRefreshTokenRepository() *MockRefreshTokenRepository {
	return &MockRefreshTokenRepository{}
}

func (mockRefreshTokenRepository *MockRefreshTokenRepository) CreateRefreshToken(refreshToken models.RefreshToken) error {
	mockRefreshTokenRepository.refreshTokens = append(mockRefreshTokenRepository.refreshTokens, refreshToken)
	return nil
}

func (mockRefreshTokenRepository *MockRefreshTokenRepository) GetRefreshToken(token string) (*models.RefreshToken, error) {
	for _, retrievedRefreshToken := range mockRefreshTokenRepository.refreshTokens {
		if retrievedRefreshToken.Token == token {
			return &retrievedRefreshToken, nil
		}
	}
	return nil, fmt.Errorf("the refresh token with value: %s haven't found", token)
}

func (mockRefreshTokenRepository *MockRefreshTokenRepository) RevokeRefreshToken(refreshToken string) error {
	for _, retrievedRefreshToken := range mockRefreshTokenRepository.refreshTokens {
		if retrievedRefreshToken.Token == refreshToken {
			retrievedRefreshToken.Revoked = true
			return nil
		}
	}
	return fmt.Errorf("the issue during revoke the refresh token")
}
