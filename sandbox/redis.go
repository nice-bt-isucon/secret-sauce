package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func main() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 1000,
	})

	rdb.Set(ctx, "mykey1", "hoge", 0)           // キー名 mykey1で文字列hogeをセット
	ret, err := rdb.Get(ctx, "mykey1").Result() // キー名mykey1を取得
	if err != nil {
		println("Error: ", err)
		return
	}

	println("Result: ", ret)
}
