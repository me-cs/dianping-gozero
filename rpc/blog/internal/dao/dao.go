package dao

import (
	"github.com/me-cs/dianping-gozero/rpc/blog/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Dao struct {
	Redis *redis.Redis
	mysql *MySQL
}

func New(c config.Config) *Dao {
	return &Dao{
		Redis: redis.MustNewRedis(c.RpcServerConf.Redis.RedisConf),
		mysql: NewMySQL(c),
	}
}
