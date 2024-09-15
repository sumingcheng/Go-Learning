package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

func addToList(
	ctx context.Context,
	client *redis.Client,
	key string,
	values ...string,
) {
	err := client.RPush(ctx, key, values).Err()
	CheckError(err)
}

func removeFromList(
	ctx context.Context,
	client *redis.Client,
	key string,
	count int,
	value string,
) {
	err := client.LRem(ctx, key, int64(count), value).Err()
	CheckError(err)
}

func updateList(
	ctx context.Context,
	client *redis.Client,
	key string,
	index int64,
	value string,
) {
	err := client.LSet(ctx, key, index, value).Err()
	CheckError(err)
}

func getList(
	ctx context.Context,
	client *redis.Client,
	key string,
	start int64,
	stop int64,
) []string {
	values, err := client.LRange(ctx, key, start, stop).Result()
	CheckError(err)
	return values
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	key := "mylist"

	// Adding to list
	addToList(ctx, client, key, "element1", "element2")

	// Getting list
	values := getList(ctx, client, key, 0, -1)
	fmt.Println("List values:", values)

	// Updating list
	updateList(ctx, client, key, 1, "newElement2")

	// Removing from list
	removeFromList(ctx, client, key, 1, "element1")

	// Getting updated list
	updatedValues := getList(ctx, client, key, 0, -1)
	fmt.Println("Updated List values:", updatedValues)
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
