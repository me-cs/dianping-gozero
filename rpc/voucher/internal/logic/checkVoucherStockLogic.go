package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/voucher/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckVoucherStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckVoucherStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckVoucherStockLogic {
	return &CheckVoucherStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckVoucherStock 检查优惠券库存
func (l *CheckVoucherStockLogic) CheckVoucherStock(in *pb.CheckVoucherStockReq) (*pb.CheckVoucherStockResp, error) {
	stock, err := l.svcCtx.Dao.GetSeckillStock(l.ctx, uint64(in.VoucherId))
	if err != nil {
		l.Logger.Errorf("get seckill stock failed: %v", err)
		return nil, err
	}

	return &pb.CheckVoucherStockResp{
		HasStock: stock > 0,
		Stock:    int32(stock),
	}, nil
}
