package services

import (
	"context"

	"github.com/google/uuid"
)

type AccessTokenPayload struct {
	UserID string
	Email  string
	Role   string
}

type RefreshTokenPayload struct {
	TokenID  string
	UserID   string
	FamilyID string
}

type TokenService interface {
	GenerateAccessToken(ctx context.Context, payload AccessTokenPayload) (string, error)
	GenerateRefreshToken(ctx context.Context, payload RefreshTokenPayload) (plain string, hash string, err error)
	ParseAccessToken(ctx context.Context, tokenString string) (*AccessTokenPayload, error)
	HashToken(token string) string
	NewTokenID() uuid.UUID
	NewFamilyID() uuid.UUID
}
