package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_set_ByLua(rdb *redis.Client, ctx context.Context) {
	luaScript_Set1 := `
		redis.call("SADD", "set1", "set1_Val1")
		redis.call("SADD", "set1", "set1_Val2")
		redis.call("SADD", "set1", "set1_Val3")
		return "success"
	`
	_, err := rdb.Eval(ctx, luaScript_Set1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis sadd set1 error: %s", err)
	}
	log.Printf("redis sadd set set1_Val1 set1_Val2 set1_Val3")
}
