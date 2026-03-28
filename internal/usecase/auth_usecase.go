package usecase

import (
	"context"
	"errors"
	"time"

	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/domain/services"
	"golang-clean-architechture/internal/dto"
	"golang-clean-architechture/internal/infrastructure/persistence"

	"gorm.io/gorm"
)

type AuthUsecase struct {
	db               *gorm.DB
	userRepo         repositories.UserRepository
	refreshTokenRepo repositories.RefreshTokenRepository
	passwordSvc      services.PasswordService
	tokenSvc         services.TokenService
}

func NewAuthUsecase(
	db *gorm.DB,
	userRepo repositories.UserRepository,
	refreshTokenRepo repositories.RefreshTokenRepository,
	passwordSvc services.PasswordService,
	tokenSvc services.TokenService,
) *AuthUsecase {
	return &AuthUsecase{
		db:               db,
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		passwordSvc:      passwordSvc,
		tokenSvc:         tokenSvc,
	}
}

func (u *AuthUsecase) Register(ctx context.Context, req dto.RegisterRequest) error {
	_, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := u.passwordSvc.Hash(req.Password)
	if err != nil {
		return err
	}

	user := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	return u.userRepo.Create(ctx, user)
}

func (u *AuthUsecase) Login(ctx context.Context, req dto.LoginRequest, userAgent, ipAddress string) (*dto.AuthResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := u.passwordSvc.Compare(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := u.tokenSvc.GenerateAccessToken(ctx, services.AccessTokenPayload{
		UserID: user.ID.String(),
		Email:  user.Email,
		Role:   user.Role,
	})
	if err != nil {
		return nil, err
	}

	tokenID := u.tokenSvc.NewTokenID()
	familyID := u.tokenSvc.NewFamilyID()

	refreshPlain, refreshHash, err := u.tokenSvc.GenerateRefreshToken(ctx, services.RefreshTokenPayload{
		TokenID:  tokenID.String(),
		UserID:   user.ID.String(),
		FamilyID: familyID.String(),
	})
	if err != nil {
		return nil, err
	}

	now := time.Now()
	refreshToken := &entities.RefreshToken{
		ID:        tokenID,
		UserID:    user.ID,
		TokenHash: refreshHash,
		FamilyID:  familyID,
		ExpiresAt: now.Add(7 * 24 * time.Hour),
		UserAgent: userAgent,
		IPAddress: ipAddress,
		CreatedAt: now,
	}

	if err := u.refreshTokenRepo.Create(ctx, refreshToken); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshPlain,
	}, nil
}

func (u *AuthUsecase) RefreshToken(ctx context.Context, req dto.RefreshRequest, userAgent, ipAddress string) (*dto.AuthResponse, error) {
	tokenHash := u.tokenSvc.HashToken(req.RefreshToken)
	currentToken, err := u.refreshTokenRepo.FindByTokenHash(ctx, tokenHash)
	if err != nil || currentToken == nil {
		return nil, errors.New("invalid refresh token")
	}

	now := time.Now()

	if currentToken.RevokedAt != nil {
		_ = u.refreshTokenRepo.RevokeFamily(ctx, currentToken.FamilyID, now)
		return nil, errors.New("refresh token reuse detected")
	}

	if now.After(currentToken.ExpiresAt) {
		return nil, errors.New("refresh token expired")
	}

	user, err := u.userRepo.FindByID(ctx, currentToken.UserID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	newTokenID := u.tokenSvc.NewTokenID()
	refreshPlain, refreshHash, err := u.tokenSvc.GenerateRefreshToken(ctx, services.RefreshTokenPayload{
		TokenID:  newTokenID.String(),
		UserID:   user.ID.String(),
		FamilyID: currentToken.FamilyID.String(),
	})
	if err != nil {
		return nil, err
	}

	newRefresh := &entities.RefreshToken{
		ID:            newTokenID,
		UserID:        user.ID,
		TokenHash:     refreshHash,
		FamilyID:      currentToken.FamilyID,
		ParentTokenID: &currentToken.ID,
		ExpiresAt:     now.Add(7 * 24 * time.Hour),
		UserAgent:     userAgent,
		IPAddress:     ipAddress,
		CreatedAt:     now,
	}

	if err := u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		refreshRepoTx := persistence.NewRefreshTokenGormRepository(tx)
		if err := refreshRepoTx.Revoke(ctx, currentToken.ID, now); err != nil {
			return err
		}
		if err := refreshRepoTx.Create(ctx, newRefresh); err != nil {
			return err
		}
		if err := refreshRepoTx.UpdateReplacement(ctx, currentToken.ID, newTokenID); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	accessToken, err := u.tokenSvc.GenerateAccessToken(ctx, services.AccessTokenPayload{
		UserID: user.ID.String(),
		Email:  user.Email,
		Role:   user.Role,
	})
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshPlain,
	}, nil
}

func (u *AuthUsecase) Logout(ctx context.Context, refreshToken string) error {
	tokenHash := u.tokenSvc.HashToken(refreshToken)
	rt, err := u.refreshTokenRepo.FindByTokenHash(ctx, tokenHash)
	if err != nil || rt == nil {
		return nil
	}
	return u.refreshTokenRepo.Revoke(ctx, rt.ID, time.Now())
}
