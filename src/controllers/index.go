package controllers

import (
	"github.com/anhnguyen300795/Task-manager/src/interfaces"
)

type Controllers struct {
	// DB      *sql.DB
	// RedisDB *redis.Client
	*interfaces.Env
}
