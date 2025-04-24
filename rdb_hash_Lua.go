package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_hash_ByLua(rdb *redis.Client, ctx context.Context) {
	luaScript_Hash1 := `
		redis.call("HSET", "hash1", "field1", "hash1_Val1")
		redis.call("HSET", "hash1", "field2", "hash1_Val2")
		return "success"
	`
	_, err := rdb.Eval(ctx, luaScript_Hash1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis hset hash1 error: %s", err)
	}
	log.Printf("redis hset hash1 field1 hash1_Val1 field2 hash1_Val2")
}
