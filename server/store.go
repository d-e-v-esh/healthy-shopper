package healthyshopper

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisStore creates a new RedisStore
func NewRedisStore(addr string, password string, db int) *RedisStore {
	r := &RedisStore{}
	r.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	r.ctx = context.Background()
	return r
}

// Set adds a key-value pair to the Redis store
func (r *RedisStore) Set(key string, value string) error {
	err := r.client.Set(r.ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves the value for a key from the Redis store
func (r *RedisStore) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Delete removes a key from the Redis store
func (r *RedisStore) Delete(key string) error {
	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
