package repository

import (
	"context"
	"fmt"

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

func (r *Repository) Save(ctx context.Context, id uint64, url string) error{
	err := r.rdb.Set(ctx, fmt.Sprintf("url:%d", id), url)
	if err!= nil{
		return err
	}
	return 
}