package repository

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type repository struct {
	redis redis.Client
}

func initDb(addr string, dbNumber int) redis.Client {
	return *redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	})
}

func New() *repository {
	return &repository{redis: initDb("127.0.0.1:6379", 2)}
}

func (r *repository) Download(FileId string) error {
	log.Println(os.Getenv("TOKEN") + "/getFile?file_id=" + FileId)
	return nil
}

func (r *repository) Upload(FileId string) error { return nil }

func (r *repository) Create(userId string, password string) error {
	ctx := context.Background()

	if err := r.redis.Set(ctx, userId, password, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *repository) ReadPassword(userId string) (string, error) {
	value, err := r.redis.Get(context.Background(), userId).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}
