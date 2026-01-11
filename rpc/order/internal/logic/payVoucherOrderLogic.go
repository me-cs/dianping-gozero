package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/order/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayVoucherOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayVoucherOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayVoucherOrderLogic {
	return &PayVoucherOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PayVoucherOrder pays for a voucher order
func (l *PayVoucherOrderLogic) PayVoucherOrder(in *pb.PayVoucherOrderReq) (*pb.PayVoucherOrderResp, error) {
	orderId := uint64(in.OrderId)
	userId := uint64(in.UserId)
	payType := uint64(in.PayType)

	// 1. 查询订单是否存在
	order, err := l.svcCtx.Dao.FindVoucherOrderById(l.ctx, orderId)
	if err != nil {
		l.Logger.Errorf("find order failed: %v", err)
		return &pb.PayVoucherOrderResp{Success: false}, nil
	}

	// 2. 验证订单是否属于该用户
	if order.UserId != userId {
		l.Logger.Errorf("order not belong to user, orderId: %d, userId: %d", orderId, userId)
		return &pb.PayVoucherOrderResp{Success: false}, nil
	}

	// 3. 验证订单状态（只有未支付状态才能支付）
	if order.Status != 1 {
		l.Logger.Errorf("order status invalid, orderId: %d, status: %d", orderId, order.Status)
		return &pb.PayVoucherOrderResp{Success: false}, nil
	}

	// 4. 更新订单状态为已支付（2），并设置支付时间和支付方式
	err = l.svcCtx.Dao.UpdateOrderStatusAndPayTime(l.ctx, orderId, 2, payType)
	if err != nil {
		l.Logger.Errorf("update order status and pay time failed: %v", err)
		return &pb.PayVoucherOrderResp{Success: false}, nil
	}

	return &pb.PayVoucherOrderResp{Success: true}, nil
}
