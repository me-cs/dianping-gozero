package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/order/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelVoucherOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelVoucherOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelVoucherOrderLogic {
	return &CancelVoucherOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CancelVoucherOrder cancels a voucher order
func (l *CancelVoucherOrderLogic) CancelVoucherOrder(in *pb.CancelVoucherOrderReq) (*pb.CancelVoucherOrderResp, error) {
	orderId := uint64(in.OrderId)
	userId := uint64(in.UserId)

	// 1. 查询订单是否存在
	order, err := l.svcCtx.Dao.FindVoucherOrderById(l.ctx, orderId)
	if err != nil {
		l.Logger.Errorf("find order failed: %v", err)
		return &pb.CancelVoucherOrderResp{Success: false}, nil
	}

	// 2. 验证订单是否属于该用户
	if order.UserId != userId {
		l.Logger.Errorf("order not belong to user, orderId: %d, userId: %d", orderId, userId)
		return &pb.CancelVoucherOrderResp{Success: false}, nil
	}

	// 3. 验证订单状态（只有未支付状态才能取消）
	if order.Status != 1 {
		l.Logger.Errorf("order status invalid, orderId: %d, status: %d", orderId, order.Status)
		return &pb.CancelVoucherOrderResp{Success: false}, nil
	}

	// 4. 更新订单状态为已取消（4）
	err = l.svcCtx.Dao.UpdateOrderStatus(l.ctx, orderId, 4)
	if err != nil {
		l.Logger.Errorf("update order status failed: %v", err)
		return &pb.CancelVoucherOrderResp{Success: false}, nil
	}

	return &pb.CancelVoucherOrderResp{Success: true}, nil
}
