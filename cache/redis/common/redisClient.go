package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 密码，没有则留空
		DB:       0,                // 默认数据库编号
	})
	return &RedisClient{
		Client: rdb,
		Ctx:    context.Background(),
	}
}
