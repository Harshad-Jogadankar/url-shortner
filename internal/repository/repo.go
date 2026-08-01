package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

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

func (r *Repository) NextId(ctx context.Context) (uint64, error){
	id, err := r.rdb.Incr(ctx, "counter")
	if err != nil{
		return "", err
	}
	return id, nil
}