package port

import (
	"context"
	"time"
)

type CacheProvider interface {
	SetObject(ctx context.Context, key string, data []byte, ttl time.Duration) error

	GetObject(ctx context.Context, key string) ([]byte, error)

	Del(ctx context.Context, key string) error

	SetString(ctx context.Context, key, str string, ttl time.Duration) error

	GetString(ctx context.Context, key string) (string, error)
}
