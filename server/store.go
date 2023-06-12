package healthyshopper

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// Cookie Config
const cookieName = "hh"
const cookieExpiration = 24 * time.Hour * 30
const httpOnly = true
const secure = false
const sameSite = http.SameSiteLaxMode

type RedisStore struct {
	client    *redis.Client
	ctx       context.Context
	resWriter http.ResponseWriter
	req       *http.Request
}

func randBase64String(entropyBytes int) string {
	b := make([]byte, entropyBytes)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

// NewRedisStore creates a new RedisStore
func NewRedisStore(addr string, password string, db int, resWriter http.ResponseWriter, req http.Request) *RedisStore {
	redisStore := &RedisStore{}
	redisStore.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	redisStore.ctx = context.Background()
	redisStore.resWriter = resWriter
	redisStore.req = &req

	return redisStore
}

// Set adds a key-value pair to the Redis store
func (r *RedisStore) SetCookie(value string) error {

	cookieValue := randBase64String(33) // 33 bytes entropy

	err := r.client.Set(r.ctx, cookieValue, value, cookieExpiration).Err()

	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		Expires:  time.Now().Add(cookieExpiration),
		HttpOnly: httpOnly,
		Secure:   secure,
		Path:     "/",
		SameSite: sameSite,
		Domain:   "localhost",
	}

	http.SetCookie(r.resWriter, cookie)

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
func (r *RedisStore) DeleteCookie(key string) error {
	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		return err
	}

	http.SetCookie(r.resWriter, &http.Cookie{
		Name:   key,
		MaxAge: -1,
	})

	return nil
}
