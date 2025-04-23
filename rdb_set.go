package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_set(rdb *redis.Client, ctx context.Context) {
	err := rdb.SAdd(ctx, "set1", "set1_Val1", "set1_Val2").Err()
	if err != nil {
		log.Fatalf("redis sadd set1 error: %s", err)
	}
	err = rdb.SAdd(ctx, "set1", "set1_Val3").Err()
	if err != nil {
		log.Fatalf("redis sadd set1 error: %s", err)
	}
}
