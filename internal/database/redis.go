package database

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(ctx context.Context, options map[string]string) (RedisClient, error) {
	addr := options["addr"]
	password := options["password"]
	dbS := options["db"]
	db, err := strconv.Atoi(dbS)
	if err != nil {
		return RedisClient{}, err
	}
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return RedisClient{}, err
	}

	return RedisClient{
		Client: client,
	}, nil
}

func (r *RedisClient) Get(ctx context.Context, ip string) (string, error) {
	return r.Client.Get(ctx, ip).Result()

}

func (r *RedisClient) Set(ctx context.Context, ip string, json []byte) (string, error) {
	return r.Client.Set(ctx, ip, json, 0).Result()

}

func (r *RedisClient) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.Client.Keys(ctx, pattern).Result()

}

func (r *RedisClient) Del(ctx context.Context, key string) (int64, error) {
	return r.Client.Del(ctx, key).Result()
}
