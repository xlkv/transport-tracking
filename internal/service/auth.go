// AuthService:
// → Cache interface (SetRefreshToken, GetRefreshToken, DeleteRefreshToken)
// → secretKey string

// 3 ta metod:
// → GenerateTokens(ctx, driverID) → access_token, refresh_token
// → Refresh(ctx, refresh_token)   → yangi access_token, refresh_token
// → Logout(ctx, refresh_token)    → Redis dan o'chir

package service

import (
	"context"

	"tracking.xlkv.com/internal/domain"
	"tracking.xlkv.com/internal/jwt"
)

type Cache interface {
	SetRefreshToken(ctx context.Context, token string, userID int64) error
	GetRefreshToken(ctx context.Context, token string) (int64, error)
	DeleteRefreshToken(ctx context.Context, token string) error
}

type AuthService struct {
	cache     Cache
	secretKey string
}

func NewAuthService(cache Cache, secretKey string) *AuthService {
	return &AuthService{
		cache:     cache,
		secretKey: secretKey,
	}
}

func (r *AuthService) GenereateTokens(ctx context.Context, driverID int64) (string, string, error) {

	refreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	acessToken, err := jwt.GenerateAccessToken(driverID, r.secretKey)
	if err != nil {
		return "", "", err
	}

	err = r.cache.SetRefreshToken(ctx, refreshToken, driverID)

	if err != nil {
		return "", "", err
	}

	return acessToken, refreshToken, nil
}

func (r *AuthService) Refresh(ctx context.Context, refresh_token string) (string, error) {
	userID, err := r.cache.GetRefreshToken(ctx, refresh_token)

	if err != nil {
		return "", domain.ErrUnauthorized
	}

	accessToken, err := jwt.GenerateAccessToken(userID, r.secretKey)
	return accessToken, err
}

func (r *AuthService) Logout(ctx context.Context, refresh_token string) error {
	return r.cache.DeleteRefreshToken(ctx, refresh_token)
}
