package rpcx

import (
	"context"

	"github.com/me-cs/dianping-gozero/library/auth"
	"github.com/me-cs/dianping-gozero/library/metadatax"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

func UserFromIncomingContext(ctx context.Context) auth.User {
	val := metadata.ValueFromIncomingContext(ctx, metadatax.UserInner)
	if len(val) == 0 {
		logx.Errorf(" len(val) == 0")
		return nil
	}
	//d, err := base64.StdEncoding.DecodeString(val[0])
	//if err != nil {
	//	logx.Errorf("  base64.StdEncoding.DecodeString")
	//	return nil
	//}
	u, err := auth.DecodeUser(val[0])
	if err != nil {
		logx.Errorf("auth.DecodeUser(string(d)),err=%v",err)
		return nil
	}
	return u
}
