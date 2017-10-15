package database

import (
	"github.com/go-redis/redis"
)

func InitializeRedis() *redis.Client {
	RedisConn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisConn.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RedisConn
}
