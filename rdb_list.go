package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_list(rdb *redis.Client, ctx context.Context) {
	err := rdb.RPush(ctx, "list1", "list1_Val1", "list1_Val2", "list1_Val3").Err()
	if err != nil {
		log.Fatalf("redis rpush list1 error: %s", err)
	}
}
