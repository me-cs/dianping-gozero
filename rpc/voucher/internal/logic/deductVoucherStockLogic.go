package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/voucher/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductVoucherStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductVoucherStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductVoucherStockLogic {
	return &DeductVoucherStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Deduct voucher stock
func (l *DeductVoucherStockLogic) DeductVoucherStock(in *pb.DeductVoucherStockReq) (*pb.DeductVoucherStockResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeductVoucherStockResp{}, nil
}
