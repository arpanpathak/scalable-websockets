package redisclient

import (
	"context"
	"fmt"

	"os"

	"github.com/go-redis/redis/v8"
)

var (
	ctx       = context.Background()
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")
)

// CreateNewRedisClient returns an instance of Redis client
func CreateNewRedisClient() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		DB:   0, // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return rdb
}
