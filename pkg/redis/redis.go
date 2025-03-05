package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func SetAndGet() {
	err := rdb.Set(context.Background(), "my-key", "my-value", 0).Err()
	if err != nil {
		log.Fatalf("Redis SET failed: %v", err)
	}

	val, err := rdb.Get(context.Background(), "my-key").Result()
	if err != nil {
		log.Fatalf("Redis GET failed: %v", err)
	}

	if val != "my-value" {
		log.Fatalf("Expected 'myvalue', got %s", val)
	}

	fmt.Println("Redis test passed")
}
