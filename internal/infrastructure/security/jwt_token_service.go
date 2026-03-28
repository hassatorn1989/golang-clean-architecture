package security

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"golang-clean-architechture/internal/domain/services"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTTokenService struct {
	accessSecret        string
	refreshSecret       string
	issuer              string
	accessExpireMinutes int
	refreshExpireDays   int
}

func NewJWTTokenService(accessSecret, refreshSecret, issuer string, accessExpireMinutes, refreshExpireDays int) *JWTTokenService {
	return &JWTTokenService{
		accessSecret:        accessSecret,
		refreshSecret:       refreshSecret,
		issuer:              issuer,
		accessExpireMinutes: accessExpireMinutes,
		refreshExpireDays:   refreshExpireDays,
	}
}

type AccessClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	TokenID  string `json:"token_id"`
	UserID   string `json:"user_id"`
	FamilyID string `json:"family_id"`
	jwt.RegisteredClaims
}

func (j *JWTTokenService) GenerateAccessToken(ctx context.Context, payload services.AccessTokenPayload) (string, error) {
	now := time.Now()
	claims := AccessClaims{
		UserID: payload.UserID,
		Email:  payload.Email,
		Role:   payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			Subject:   payload.UserID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(j.accessExpireMinutes) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.accessSecret))
}

func (j *JWTTokenService) GenerateRefreshToken(ctx context.Context, payload services.RefreshTokenPayload) (string, string, error) {
	now := time.Now()
	claims := RefreshClaims{
		TokenID:  payload.TokenID,
		UserID:   payload.UserID,
		FamilyID: payload.FamilyID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			Subject:   payload.UserID,
			ID:        payload.TokenID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(j.refreshExpireDays) * 24 * time.Hour)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := jwtToken.SignedString([]byte(j.refreshSecret))
	if err != nil {
		return "", "", err
	}

	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", "", err
	}
	nonce := base64.RawURLEncoding.EncodeToString(randomBytes)
	plain := signed + "." + nonce

	return plain, j.HashToken(plain), nil
}

func (j *JWTTokenService) ParseAccessToken(ctx context.Context, tokenString string) (*services.AccessTokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.accessSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	return &services.AccessTokenPayload{
		UserID: claims.UserID,
		Email:  claims.Email,
		Role:   claims.Role,
	}, nil
}

func (j *JWTTokenService) HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (j *JWTTokenService) NewTokenID() uuid.UUID {
	return uuid.New()
}

func (j *JWTTokenService) NewFamilyID() uuid.UUID {
	return uuid.New()
}

func (j *JWTTokenService) ParseRefreshTokenClaims(tokenString string) (*RefreshClaims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) < 4 {
		return nil, errors.New("invalid refresh token format")
	}
	jwtPart := strings.Join(parts[:3], ".")

	token, err := jwt.ParseWithClaims(jwtPart, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.refreshSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}
