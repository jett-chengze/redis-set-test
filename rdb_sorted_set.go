package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_sorted_set(rdb *redis.Client, ctx context.Context) {
	err := rdb.ZAdd(ctx, "sort_set1", redis.Z{Score: 2, Member: "sort_set1_Val2"}, redis.Z{Score: 1, Member: "sort_set1_Val1"}).Err()
	if err != nil {
		log.Fatalf("redis zadd sort_set1 error: %s", err)
	}
	log.Println("redis zadd sort_set1 1 sort_set1_Val1 2 sort_set1_Val2")
}
