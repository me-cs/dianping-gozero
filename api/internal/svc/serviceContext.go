// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/me-cs/dianping-gozero/api/internal/config"
	"github.com/me-cs/dianping-gozero/api/internal/dao"
	"github.com/me-cs/dianping-gozero/api/internal/middleware"
	"github.com/me-cs/dianping-gozero/library/rpcx"
	"github.com/me-cs/dianping-gozero/rpc/blog/blog"
	"github.com/me-cs/dianping-gozero/rpc/order/order"
	"github.com/me-cs/dianping-gozero/rpc/shop/shop"
	"github.com/me-cs/dianping-gozero/rpc/user/user"
	"github.com/me-cs/dianping-gozero/rpc/voucher/voucher"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                 config.Config
	Dao                    *dao.Dao
	UserRpc                user.User
	BlogRpc                blog.Blog
	OrderRpc               order.Order
	VoucherRpc             voucher.Voucher
	ShopRpc                shop.Shop
	RefreshTokenMiddleware rest.Middleware
	AuthMiddleware         rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	d := dao.New(c)
	return &ServiceContext{
		Config:                 c,
		Dao:                    d,
		UserRpc:                user.NewUser(rpcx.MustNewClient(c.UserRpc)),
		BlogRpc:                blog.NewBlog(rpcx.MustNewClient(c.BlogRpc)),
		OrderRpc:               order.NewOrder(rpcx.MustNewClient(c.OrderRpc)),
		VoucherRpc:             voucher.NewVoucher(rpcx.MustNewClient(c.VoucherRpc)),
		ShopRpc:                shop.NewShop(rpcx.MustNewClient(c.ShopRpc)),
		RefreshTokenMiddleware: middleware.NewRefreshTokenMiddleware(d.Redis).Handle,
		AuthMiddleware:         middleware.NewAuthMiddleware().Handle,
	}
}
