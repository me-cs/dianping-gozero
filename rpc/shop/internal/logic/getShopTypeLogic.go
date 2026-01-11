package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/shop/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/shop/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShopTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopTypeLogic {
	return &GetShopTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get shop type by ID
func (l *GetShopTypeLogic) GetShopType(in *pb.GetShopTypeReq) (*pb.GetShopTypeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetShopTypeResp{}, nil
}
