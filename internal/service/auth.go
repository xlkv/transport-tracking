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

	"tracking.xlkv.com/internal/jwt"
)

type Cache interface {
	SetRefreshToken(ctx context.Context, token string, userID int64) error
	GetRefreshToken(ctx context.Context, token string) (int64, error)
	DeleteRefreshToken(ctx context.Context, token string) error
}

type AuthRepositroy struct {
	Cache     Cache
	SecretKey string
}

func NewAuthRepository(cache Cache, secretKey string) *AuthRepositroy {
	return &AuthRepositroy{
		Cache:     cache,
		SecretKey: secretKey,
	}
}

func (r *AuthRepositroy) GenereateTokens(ctx context.Context, driverID int64) (string, string, error) {

	refreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	acessToken, err := jwt.GenerateAccessToken(driverID, r.SecretKey)
	if err != nil {
		return "", "", err
	}

	err = r.Cache.SetRefreshToken(ctx, refreshToken, driverID)

	if err != nil {
		return "", "", err
	}

	return acessToken, refreshToken, nil
}

