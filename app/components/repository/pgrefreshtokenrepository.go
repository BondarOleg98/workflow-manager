package repository

import (
	"github.com/oklog/ulid/v2"
	"time"
	"workflowmanager/app/models"
)

type RefreshTokenRepository interface {
	CreateRefreshToken(userId ulid.ULID, ttl time.Duration) (*models.RefreshToken, error)
	GetRefreshToken(token string) (*models.RefreshToken, error)
	RevokeRefreshToken(refreshToken string) error
}
