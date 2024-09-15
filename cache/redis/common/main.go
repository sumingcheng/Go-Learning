package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := NewRedisClient()

	// String
	err := rdb.Set(ctx, "UserName", "admin1", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "UserName").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("UserName:", val)

	// 使用 List
	rdb.RPush(ctx, "list", "element1")
	rdb.RPush(ctx, "list", "element2")
	listVal, err := rdb.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", listVal)

	// 使用 Set
	rdb.SAdd(ctx, "set", "member1")
	rdb.SAdd(ctx, "set", "member2")
	setVal, err := rdb.SMembers(ctx, "set").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("set:", setVal)

	// 使用 Sorted Set
	rdb.ZAdd(ctx, "sortedset", redis.Z{Score: 1, Member: "one"})
	rdb.ZAdd(ctx, "sortedset", redis.Z{Score: 2, Member: "two"})
	sortedSetVal, err := rdb.ZRange(ctx, "sortedset", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("sorted set:", sortedSetVal)

	// 使用 Hash
	rdb.HSet(ctx, "hash", "field1", "value1")
	rdb.HSet(ctx, "hash", "field2", "value2")
	hashVal, err := rdb.HGetAll(ctx, "hash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("hash:", hashVal)

	// 关闭连接
	rdb.Close()
}
