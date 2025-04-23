package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_hash(rdb *redis.Client, ctx context.Context) {
	err := rdb.HSet(ctx, "hash1", "field1", "hash1_Val1", "field2", "hash1_Val2").Err()
	if err != nil {
		log.Fatalf("redis hset hash1 error: %s", err)
	}
}
