package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisConn *redis.Client

func InitializeRedis() {
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisConn.Ping().Result()
	fmt.Println(pong, err)
}
