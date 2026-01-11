package rpcx

import (
	"context"

	"github.com/me-cs/dianping-gozero/library/auth"
	"github.com/me-cs/dianping-gozero/library/metadatax"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UserInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		user := UserFromIncomingContext(ctx)
		ec, err := auth.EncodeUser(user)
		if err == nil {
			ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(metadatax.UserInner, ec))
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func MustNewClient(c zrpc.RpcClientConf, options ...zrpc.ClientOption) zrpc.Client {
	return zrpc.MustNewClient(c, zrpc.WithUnaryClientInterceptor(UserInterceptor()))
}
