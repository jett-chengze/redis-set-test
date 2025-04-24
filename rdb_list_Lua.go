package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_list_ByLua(rdb *redis.Client, ctx context.Context) {
	luaScript_List1 := `
		redis.call("RPUSH", "list1", "list1_Val1", "list1_Val2", "list1_Val3")
		return ""
	`
	_, err := rdb.Eval(ctx, luaScript_List1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis rpush list error: %s", err)
	}
	log.Printf("redis rpush list1 list_Val1 list1_Val2 list1_Val3")
}
