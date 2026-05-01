package cache

import (
	"context"
	"time"
)

func (c *RedisCache) SetRefreshToken(ctx context.Context, token string, userID int64) error {
	key := "refresh_token:" + token
	return c.Set(ctx, key, userID, 24*7*time.Hour)
}

func (c *RedisCache) GetRefreshToken(ctx context.Context, token string) (int64, error) {
	key := "refresh_token:" + token
	var id int64
	err := c.Get(ctx, key, &id)
	return id, err
}

func (c *RedisCache) DeleteRefreshToken(ctx context.Context, token string) error {
	key := "refresh_token:" + token
	return c.Delete(ctx, key)
}
