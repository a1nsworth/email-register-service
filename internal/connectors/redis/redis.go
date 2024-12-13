package redis

import "email-register-service/internal/config"

type Redis struct {
	config config.Redis
}

func NewRedis(config config.Redis) *Redis {
	return &Redis{
		config: config,
	}
}
