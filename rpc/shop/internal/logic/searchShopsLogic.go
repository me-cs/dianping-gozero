package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/shop/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/shop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchShopsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchShopsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchShopsLogic {
	return &SearchShopsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Search shops
func (l *SearchShopsLogic) SearchShops(in *pb.SearchShopsReq) (*pb.SearchShopsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchShopsResp{}, nil
}
