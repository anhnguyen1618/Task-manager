package controllers

import (
	"github.com/anhnguyen300795/Task-manager/interfaces"
)

type Controllers struct {
	// DB      *sql.DB
	// RedisDB *redis.Client
	*interfaces.Env
}
