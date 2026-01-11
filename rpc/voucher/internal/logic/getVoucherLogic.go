package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/voucher/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoucherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoucherLogic {
	return &GetVoucherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get voucher by ID
func (l *GetVoucherLogic) GetVoucher(in *pb.GetVoucherReq) (*pb.GetVoucherResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetVoucherResp{}, nil
}
