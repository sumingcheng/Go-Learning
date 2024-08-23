package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func setValue(
	ctx context.Context,
	client *redis.Client,
	key string,
	value string,
) {
	err := client.Set(ctx, key, value, 30*time.Second).Err()
	CheckError(err)
}

func getValue(
	ctx context.Context,
	client *redis.Client,
	key string,
) string {
	value, err := client.Get(ctx, key).Result()
	CheckError(err)
	return value
}

func deleteValue(
	ctx context.Context,
	client *redis.Client,
	key string,
) {
	err := client.Del(ctx, key).Err()
	CheckError(err)
}

func setExpiration(
	ctx context.Context,
	client *redis.Client,
	key string,
	duration time.Duration,
) {
	err := client.Expire(ctx, key, duration).Err()
	CheckError(err)
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	key := "name"
	value := "sumingcheng"

	setValue(ctx, client, key, value)
	setExpiration(ctx, client, key, 60*time.Second) // Set expiration to 60 seconds
	retrievedValue := getValue(ctx, client, key)
	fmt.Println("Retrieved Value:", retrievedValue)
	deleteValue(ctx, client, key)
}

func CheckError(err error) {
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("Key does not exist")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
