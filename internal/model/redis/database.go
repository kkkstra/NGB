package redis

import (
	"NGB/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

func init() {
	// 连接到redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.C.Database.Redis.Host, config.C.Database.Redis.Port),
		Password: config.C.Database.Redis.Password,
		DB:       config.C.Database.Redis.Db,
	})
	// 测试连接
	// ctx := context.Background()
	// _, err := rdb.Ping(ctx).Result()
	// if err != nil {
	// 	logrus.Logger.Error(err)
	// }
	// logrus.Logger.Info("Redis: connected to database")
}

// 获取key的值
func (c *Client) Get(key string) (string, error) {
	ctx := context.Background()
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// 设置key, value和expiration
func (c *Client) Set(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
