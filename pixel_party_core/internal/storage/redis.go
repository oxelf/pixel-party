package storage

import "github.com/redis/go-redis/v9"

type RedisConnection struct {
	*redis.Client
}

func StartRedis(customAddress *string) *RedisConnection {
	address := "localhost:6379"
	if customAddress != nil {
		address = *customAddress
	}
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return &RedisConnection{client}
}
