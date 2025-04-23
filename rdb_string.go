package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_string(rdb *redis.Client, ctx context.Context) {
	err := rdb.Set(ctx, "string1", "string1_Val1", 0).Err()
	if err != nil {
		log.Fatalf("redis set string1 error: %s", err)
	}
}
