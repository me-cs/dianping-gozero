// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Mysql struct {
		DataSource string
	}

	Redis      redis.RedisConf
	CacheRedis cache.CacheConf

	Etcd struct {
		Hosts []string
		Key   string
	}

	UserRpc    zrpc.RpcClientConf
	ShopRpc    zrpc.RpcClientConf
	VoucherRpc zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf
	BlogRpc    zrpc.RpcClientConf

	
	Upload struct {
		Dir string // 上传文件保存目录
	}
}
