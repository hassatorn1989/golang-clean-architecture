package persistence

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/infrastructure/persistence/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenGormRepository struct {
	db *gorm.DB
}

func NewRefreshTokenGormRepository(db *gorm.DB) repositories.RefreshTokenRepository {
	return &RefreshTokenGormRepository{db: db}
}

func (r *RefreshTokenGormRepository) Create(ctx context.Context, token *entities.RefreshToken) error {
	model := &models.RefreshTokenModel{
		ID:            token.ID,
		UserID:        token.UserID,
		TokenHash:     token.TokenHash,
		FamilyID:      token.FamilyID,
		ParentTokenID: token.ParentTokenID,
		ExpiresAt:     token.ExpiresAt,
		RevokedAt:     token.RevokedAt,
		ReplacedByID:  token.ReplacedByID,
		UserAgent:     token.UserAgent,
		IPAddress:     token.IPAddress,
		CreatedAt:     token.CreatedAt,
	}
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *RefreshTokenGormRepository) FindByTokenHash(ctx context.Context, tokenHash string) (*entities.RefreshToken, error) {
	var model models.RefreshTokenModel
	if err := r.db.WithContext(ctx).Where("token_hash = ?", tokenHash).First(&model).Error; err != nil {
		return nil, err
	}

	return &entities.RefreshToken{
		ID:            model.ID,
		UserID:        model.UserID,
		TokenHash:     model.TokenHash,
		FamilyID:      model.FamilyID,
		ParentTokenID: model.ParentTokenID,
		ExpiresAt:     model.ExpiresAt,
		RevokedAt:     model.RevokedAt,
		ReplacedByID:  model.ReplacedByID,
		UserAgent:     model.UserAgent,
		IPAddress:     model.IPAddress,
		CreatedAt:     model.CreatedAt,
	}, nil
}

func (r *RefreshTokenGormRepository) Revoke(ctx context.Context, tokenID uuid.UUID, revokedAt time.Time) error {
	return r.db.WithContext(ctx).
		Model(&models.RefreshTokenModel{}).
		Where("id = ?", tokenID).
		Update("revoked_at", revokedAt).Error
}

func (r *RefreshTokenGormRepository) RevokeFamily(ctx context.Context, familyID uuid.UUID, revokedAt time.Time) error {
	return r.db.WithContext(ctx).
		Model(&models.RefreshTokenModel{}).
		Where("family_id = ? AND revoked_at IS NULL", familyID).
		Update("revoked_at", revokedAt).Error
}

func (r *RefreshTokenGormRepository) UpdateReplacement(ctx context.Context, tokenID uuid.UUID, replacedByID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&models.RefreshTokenModel{}).
		Where("id = ?", tokenID).
		Update("replaced_by_id", replacedByID).Error
}
