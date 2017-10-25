package interfaces

import (
	"database/sql"

	"github.com/go-redis/redis"
)

type Env struct {
	DB      *sql.DB
	RedisDB *redis.Client
}
