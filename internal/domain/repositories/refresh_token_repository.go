package repositories

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, token *entities.RefreshToken) error
	FindByTokenHash(ctx context.Context, tokenHash string) (*entities.RefreshToken, error)
	Revoke(ctx context.Context, tokenID uuid.UUID, revokedAt time.Time) error
	RevokeFamily(ctx context.Context, familyID uuid.UUID, revokedAt time.Time) error
	UpdateReplacement(ctx context.Context, tokenID uuid.UUID, replacedByID uuid.UUID) error
}
