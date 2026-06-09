package unit

import (
	"fmt"
	"slices"
	"workflowmanager/app/models"
)

type MockRefreshTokenRepository struct {
	refreshTokens []models.RefreshToken
}

func NewMockRefreshTokenRepository() *MockRefreshTokenRepository {
	return &MockRefreshTokenRepository{}
}

func (mockRefreshTokenRepository *MockRefreshTokenRepository) CreateRefreshToken(refreshToken models.RefreshToken) error {
	mockRefreshTokenRepository.refreshTokens = append(mockRefreshTokenRepository.refreshTokens, refreshToken)
	return nil
}

func (mockRefreshTokenRepository *MockRefreshTokenRepository) GetRefreshTokenByValue(token string) (*models.RefreshToken, error) {
	for _, retrievedRefreshToken := range mockRefreshTokenRepository.refreshTokens {
		if retrievedRefreshToken.Token == token {
			return &retrievedRefreshToken, nil
		}
	}
	return nil, fmt.Errorf("the refresh token with value: %s haven't found", token)
}

func (mockRefreshTokenRepository *MockRefreshTokenRepository) GetRefreshTokenByUserId(userId string) (*models.RefreshToken, error) {
	for _, retrievedRefreshToken := range mockRefreshTokenRepository.refreshTokens {
		if retrievedRefreshToken.Token == userId {
			return &retrievedRefreshToken, nil
		}
	}
	return nil, fmt.Errorf("the refresh token with userId: %s haven't found", userId)
}

func (mockRefreshTokenRepository *MockRefreshTokenRepository) RevokeRefreshToken(refreshToken string) error {
	for indexRefreshToken, retrievedRefreshToken := range mockRefreshTokenRepository.refreshTokens {
		if retrievedRefreshToken.Token == refreshToken {
			mockRefreshTokenRepository.refreshTokens[indexRefreshToken].Revoked = true
			return nil
		}
	}
	return fmt.Errorf("the refresh token with value: %s haven't found", refreshToken)
}

func (mockRefreshTokenRepository *MockRefreshTokenRepository) RemoveRefreshToken(refreshToken string) error {
	tmpRefreshTokens := mockRefreshTokenRepository.refreshTokens
	for indexRefreshToken, retrievedRefreshToken := range tmpRefreshTokens {
		if retrievedRefreshToken.Token == refreshToken {
			mockRefreshTokenRepository.refreshTokens = slices.Delete(mockRefreshTokenRepository.refreshTokens, indexRefreshToken, indexRefreshToken)
			return nil
		}
	}
	return fmt.Errorf("the refresh token with value: %s haven't found", refreshToken)
}
