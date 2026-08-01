package repository

import "github.com/redis/go-redis/v9"

type Repository struct {
	rdb *redis.Client
}

func NewRepository () Repository{
	redisClient := redis.NewClient(
		&redis.Options{
			Addr: config.get()
		}
	)

	return &Repository{
		rdb: redisClient
	}
}