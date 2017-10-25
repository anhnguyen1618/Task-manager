package database

import (
	"github.com/go-redis/redis"
)

func InitializeRedis() *redis.Client {
	RedisConn := redis.NewClient(&redis.Options{
		Addr:     "redis-17404.c1.eu-west-1-3.ec2.cloud.redislabs.com:17404",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisConn.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RedisConn
}
