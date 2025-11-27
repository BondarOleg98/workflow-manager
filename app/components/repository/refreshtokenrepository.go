package repository

import (
	"workflowmanager/app/models"
)

type RefreshTokenRepository interface {
	CreateRefreshToken(token models.RefreshToken) error
	GetRefreshToken(token string) (*models.RefreshToken, error)
	RevokeRefreshToken(refreshToken string) error
}
