package database

import "context"

type Persist interface {
	Get(ctx context.Context, ip string) (string, error)
	Set(ctx context.Context, ip string, json []byte) (string, error)
	Keys(ctx context.Context, pattern string) ([]string, error)
	Del(ctx context.Context, key string) (int64, error)
}
