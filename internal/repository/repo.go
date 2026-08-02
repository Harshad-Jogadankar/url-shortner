package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	rdb *redis.Client
}

func NewRepository() *Repository {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)

	return &Repository{
		rdb: redisClient,
	}
}

func (r *Repository) NextId(ctx context.Context) (uint64, error) {
	id, err := r.rdb.Incr(ctx, "counter").Result()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (r *Repository) Save(ctx context.Context, shortenedUrl string, originalUrl string) error {
	_, err := r.rdb.Set(ctx, fmt.Sprintf("url:%s", shortenedUrl), originalUrl, 24*time.Hour).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Fom gets a different code than Bob shortening the same URL, e.g. for separate click analytics, etch(ctx context.Context, shortenedUrl string) (string, error) {
	originalUrl, err := r.rdb.Get(ctx, fmt.Sprintf("url:%s", shortenedUrl)).Result()
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}
