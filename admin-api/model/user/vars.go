package user

import (
	"github.com/go-redis/redis/v8"
	"github.com/haozheyu/oam_system/admin-api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass, // no password set
		DB:       0,            // use default DB
	})
}
