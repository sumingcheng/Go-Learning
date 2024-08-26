package main

import "github.com/redis/go-redis/v9"

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.22.220.64:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
