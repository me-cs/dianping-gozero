package svc

import (
	"github.com/me-cs/dianping-gozero/library/rpcx"
	"github.com/me-cs/dianping-gozero/rpc/order/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/order/internal/dao"
	"github.com/me-cs/dianping-gozero/rpc/order/internal/utils"
	"github.com/me-cs/dianping-gozero/rpc/voucher/voucher"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config           config.Config
	Dao              *dao.Dao
	IDGen            *utils.IDGenerator
	VoucherRpcClient voucher.Voucher
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.RpcServerConf.Redis.RedisConf)
	return &ServiceContext{
		Config:           c,
		Dao:              dao.New(c),
		IDGen:            utils.NewIDGenerator(rds),
		VoucherRpcClient: voucher.NewVoucher(rpcx.MustNewClient(c.VoucherRpc)),
	}
}
