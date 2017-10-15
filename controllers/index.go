package controllers

import (
	"../interfaces"
)

type Controllers struct {
	// DB      *sql.DB
	// RedisDB *redis.Client
	*interfaces.Env
}
