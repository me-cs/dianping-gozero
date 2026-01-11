package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/voucher/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVoucherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVoucherLogic {
	return &AddVoucherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Add voucher
func (l *AddVoucherLogic) AddVoucher(in *pb.AddVoucherReq) (*pb.AddVoucherResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddVoucherResp{}, nil
}
