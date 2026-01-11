package svc

import (
	"github.com/me-cs/dianping-gozero/rpc/shop/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/shop/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	Dao    *dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
	}
}
