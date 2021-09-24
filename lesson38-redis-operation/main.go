package main

import (
	"context"
	"fmt"

	//"github.com/go-redis/redis"
	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

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

func hGetAllDemo() {
	v := rdb.HGetAll(ctx, "user").Val()
	fmt.Println(v)

	v2 := rdb.HMGet(ctx, "user", "name", "age").Val()
	fmt.Println(v2)

	v3 := rdb.HGet(ctx, "user", "age")
	fmt.Println(v3)
}

////resource code
//func (cmd *StringCmd) Val() string {
//	return cmd.val
//}
//
//func (cmd *StringCmd) Result() (string, error) {
//	return cmd.Val(), cmd.err
//}

func main() {
	if err := initializeRedisClient(); err != nil {
		fmt.Printf("connect 2 redis failed,err %v\n", err)
		panic(err)
	}
	fmt.Println("connect 2 redis success")

	//set
	err := rdb.Set(ctx, "name", "wangxu", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis set success")

	//get
	//1
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("get success name:%v\n", val)
	//2
	val2 := rdb.Get(ctx, "name").Val()
	fmt.Println("name:", val2)

	//redis.nil
	val3, err := rdb.Get(ctx, "testKey").Result()
	if err == redis.Nil {
		fmt.Println("testKey does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("testKey:", val3)
	}

	hGetAllDemo()
}
