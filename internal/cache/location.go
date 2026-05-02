package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"tracking.xlkv.com/internal/domain"
)

func (c *RedisCache) SetLocation(ctx context.Context, vehicleID int64, location domain.Location) error {
	key := "location:vehicle:" + strconv.Itoa(int(vehicleID))
	return c.Set(ctx, key, location, 10*time.Second)
}

func (c *RedisCache) PublishLocation(ctx context.Context, vehicleID int64, location domain.Location) error {
	channel := fmt.Sprintf("location:vehicle:%d", vehicleID)
	data, err := json.Marshal(location)
	if err != nil {
		return err
	}
	return c.client.Publish(ctx, channel, data).Err()
}

func (c *RedisCache) SubscribeLocation(ctx context.Context, vehicleID int64) *redis.PubSub {
	channel := fmt.Sprintf("location:vehicle:%d", vehicleID)
	return c.client.Subscribe(ctx, channel)
}
