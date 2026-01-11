package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/order/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/order/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type GetVoucherOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoucherOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoucherOrderLogic {
	return &GetVoucherOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get voucher order by ID
func (l *GetVoucherOrderLogic) GetVoucherOrder(in *pb.GetVoucherOrderReq) (*pb.GetVoucherOrderResp, error) {
	orderId := uint64(in.OrderId)

	// 查询订单
	order, err := l.svcCtx.Dao.FindVoucherOrderById(l.ctx, orderId)
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.GetVoucherOrderResp{}, nil
		}
		l.Logger.Errorf("find voucher order failed: %v", err)
		return nil, err
	}

	// 转换为pb格式
	return &pb.GetVoucherOrderResp{
		Order: l.orderToProto(order),
	}, nil
}

// orderToProto 将model转换为proto格式
func (l *GetVoucherOrderLogic) orderToProto(order *model.TbVoucherOrder) *pb.VoucherOrderInfo {
	info := &pb.VoucherOrderInfo{
		Id:         order.Id,
		UserId:     int64(order.UserId),
		VoucherId:  int64(order.VoucherId),
		PayType:    int32(order.PayType),
		Status:     int32(order.Status),
		CreateTime: order.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: order.UpdateTime.Format("2006-01-02 15:04:05"),
	}

	if order.PayTime.Valid {
		info.PayTime = order.PayTime.Time.Format("2006-01-02 15:04:05")
	}
	if order.UseTime.Valid {
		info.UseTime = order.UseTime.Time.Format("2006-01-02 15:04:05")
	}
	if order.RefundTime.Valid {
		info.RefundTime = order.RefundTime.Time.Format("2006-01-02 15:04:05")
	}

	return info
}
