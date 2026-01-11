package main

import (
	"flag"
	"fmt"

	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/server"
	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/voucher/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/voucher.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterVoucherServer(grpcServer, server.NewVoucherServer(ctx))

		if c.RpcServerConf.Mode == service.DevMode || c.RpcServerConf.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting voucher rpc server at %s...\n", c.RpcServerConf.ListenOn)
	s.Start()
}
