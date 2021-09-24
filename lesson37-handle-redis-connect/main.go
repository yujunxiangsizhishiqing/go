package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)


var rdb *redis.Client
var ctx = context.Background()
func initializeRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "yj950627",
		DB:       0,
		PoolSize: 100,
	})

	_, err = rdb.Ping(ctx).Result()
	return

}

func main() {
	if err := initializeRedisClient(); err != nil {
		fmt.Printf("connect 2 redis failed err:%v", err)
		panic(err)
	}
	fmt.Println("connect rediss success")

}
