package main

import (
	"flag"
	"fmt"

	"github.com/me-cs/dianping-gozero/rpc/order/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/order/internal/server"
	"github.com/me-cs/dianping-gozero/rpc/order/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServer(grpcServer, server.NewOrderServer(ctx))

		if c.RpcServerConf.Mode == service.DevMode || c.RpcServerConf.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting order rpc server at %s...\n", c.RpcServerConf.ListenOn)
	s.Start()
}
