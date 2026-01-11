package dao

import (
	"database/sql"
	//"database/sql/driver"

	"github.com/me-cs/dianping-gozero/api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Dao struct {
	Redis *redis.Redis
	mysql *sql.DB
}

func New(c config.Config) *Dao {
	return &Dao{
		Redis: redis.MustNewRedis(c.Redis),
	}
}
