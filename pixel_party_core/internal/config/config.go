package config

import "github.com/redis/go-redis/v9"

func getRedisConfig() redis.Options {
return redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
        Protocol: 2,
    }
}
