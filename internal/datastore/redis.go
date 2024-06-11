package datastore

import (
	"TMS-GIN/config"
	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func InitRedis(redisCfg *config.Redis) {
	addr := redisCfg.Host + ":" + redisCfg.Port
	Cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
}
