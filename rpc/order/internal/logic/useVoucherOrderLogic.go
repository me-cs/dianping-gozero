package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/order/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseVoucherOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUseVoucherOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseVoucherOrderLogic {
	return &UseVoucherOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UseVoucherOrder marks a voucher order as used
func (l *UseVoucherOrderLogic) UseVoucherOrder(in *pb.UseVoucherOrderReq) (*pb.UseVoucherOrderResp, error) {
	orderId := uint64(in.OrderId)
	userId := uint64(in.UserId)

	// 1. 查询订单是否存在
	order, err := l.svcCtx.Dao.FindVoucherOrderById(l.ctx, orderId)
	if err != nil {
		l.Logger.Errorf("find order failed: %v", err)
		return &pb.UseVoucherOrderResp{Success: false}, nil
	}

	// 2. 验证订单是否属于该用户
	if order.UserId != userId {
		l.Logger.Errorf("order not belong to user, orderId: %d, userId: %d", orderId, userId)
		return &pb.UseVoucherOrderResp{Success: false}, nil
	}

	// 3. 验证订单状态（只有已支付状态才能核销）
	if order.Status != 2 {
		l.Logger.Errorf("order status invalid, orderId: %d, status: %d", orderId, order.Status)
		return &pb.UseVoucherOrderResp{Success: false}, nil
	}

	// 4. 更新订单状态为已核销（3），并设置核销时间
	err = l.svcCtx.Dao.UpdateOrderStatusAndUseTime(l.ctx, orderId, 3)
	if err != nil {
		l.Logger.Errorf("update order status and use time failed: %v", err)
		return &pb.UseVoucherOrderResp{Success: false}, nil
	}

	return &pb.UseVoucherOrderResp{Success: true}, nil
}
