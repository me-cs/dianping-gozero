package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/order/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVoucherOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVoucherOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVoucherOrderLogic {
	return &CreateVoucherOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create voucher order (seckill)
func (l *CreateVoucherOrderLogic) CreateVoucherOrder(in *pb.CreateVoucherOrderReq) (*pb.CreateVoucherOrderResp, error) {
	userId := uint64(in.UserId)
	voucherId := uint64(in.VoucherId)

	// 1. 生成订单ID
	orderId, err := l.svcCtx.IDGen.NextID(l.ctx, "order")
	if err != nil {
		l.Logger.Errorf("generate order id failed: %v", err)
		return &pb.CreateVoucherOrderResp{
			Success: false,
			Message: "生成订单ID失败",
		}, nil
	}

	// 2. 执行Lua脚本进行秒杀判断
	// 返回值：0=成功，1=库存不足，2=用户已下单
	result, err := l.svcCtx.Dao.SeckillVoucher(l.ctx, voucherId, userId, orderId)
	if err != nil {
		l.Logger.Errorf("seckill voucher failed: %v", err)
		return &pb.CreateVoucherOrderResp{
			Success: false,
			Message: "秒杀失败，请稍后重试",
		}, nil
	}

	// 3. 判断结果
	switch result {
	case 0:
		// 成功：订单信息已发送到Redis Stream队列，异步处理
		return &pb.CreateVoucherOrderResp{
			OrderId: int64(orderId),
			Success: true,
			Message: "下单成功",
		}, nil
	case 1:
		// 库存不足
		return &pb.CreateVoucherOrderResp{
			Success: false,
			Message: "库存不足",
		}, nil
	case 2:
		// 用户已下单
		return &pb.CreateVoucherOrderResp{
			Success: false,
			Message: "不能重复下单",
		}, nil
	default:
		// 未知错误
		return &pb.CreateVoucherOrderResp{
			Success: false,
			Message: "下单失败",
		}, nil
	}
}
