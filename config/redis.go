package config

import "github.com/redis/go-redis/v9"

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var Cache *redis.Client

func InitRedis() {
	addr := Cfg.Redis.Host + ":" + Cfg.Redis.Port
	Cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: Cfg.Redis.Password,
		DB:       Cfg.Redis.DB,
	})
}
