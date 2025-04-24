package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_string_ByLua(rdb *redis.Client, ctx context.Context) {
	luaScript_string1 := `
		redis.call("SET", "string1", "string1_Val1")
		return "success"
	`
	_, err := rdb.Eval(ctx, luaScript_string1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis set string1 error: %s", err)
	}
	log.Printf("redis set string1")
}
