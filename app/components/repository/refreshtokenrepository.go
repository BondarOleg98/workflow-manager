package repository

import (
	"workflowmanager/app/models"
)

type RefreshTokenRepository interface {
	CreateRefreshToken(token models.RefreshToken) error
	GetRefreshTokenByValue(token string) (*models.RefreshToken, error)
	GetRefreshTokenByUserId(userId string) (*models.RefreshToken, error)
	RevokeRefreshToken(refreshToken string) error
	RemoveRefreshToken(refreshToken string) error
}
