package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_sroted_set_ByLua(rdb *redis.Client, ctx context.Context) {
	luaScript_sort_set1 := `
		redis.call("ZADD", "sort_set1", 2, "sort_set1_Val2")
		redis.call("ZADD", "sort_set1", 1, "sort_set1_Val1")
		return "success"
	`
	_, err := rdb.Eval(ctx, luaScript_sort_set1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis zadd sort_set1 error: %s", err)
	}
	log.Println("redis zadd sort_set1 1 sort_set1_Val1 2 sort_set1_Val2")
}
