package global

import (
	"time"
	"wecat/common/setting"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func NewRedisClient(s *setting.RedisSettingS) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:            s.Host,
		Password:        s.Password,
		DB:              0,
		MaxActiveConns:  s.MaxActive,
		ConnMaxIdleTime: time.Duration(time.Microsecond * time.Duration(s.IdleTimeout)),
	})

	return rdb, nil
}
