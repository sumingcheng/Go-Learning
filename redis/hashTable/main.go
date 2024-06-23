package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// HSET - 添加键值对到哈希表
	err := client.HSet(ctx, "user", "name", "John", "age", 30).Err()
	CheckError(err)

	// HGET - 获取哈希表中的值
	name, err := client.HGet(ctx, "user", "name").Result()
	CheckError(err)
	fmt.Println("Name:", name)

	// HGETALL - 获取哈希表中的所有键值对
	allFields, err := client.HGetAll(ctx, "user").Result()
	CheckError(err)
	fmt.Println("All fields:", allFields)

	//// HDEL - 删除哈希表中的键
	//err = client.HDel(ctx, "user", "age").Err()
	//CheckError(err)
	//
	//// HMSET (过时) - 修改哈希表中的键值对
	//err = client.HMSet(ctx, "user", map[string]interface{}{"age": "31", "gender": "male"}).Err()
	//CheckError(err)
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
