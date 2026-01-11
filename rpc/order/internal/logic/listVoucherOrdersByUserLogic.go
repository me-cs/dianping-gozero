package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/order/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/order/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type ListVoucherOrdersByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVoucherOrdersByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVoucherOrdersByUserLogic {
	return &ListVoucherOrdersByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// List voucher orders by user
func (l *ListVoucherOrdersByUserLogic) ListVoucherOrdersByUser(in *pb.ListVoucherOrdersByUserReq) (*pb.ListVoucherOrdersByUserResp, error) {
	userId := uint64(in.UserId)
	status := int(in.Status)

	// 查询订单列表
	orders, err := l.svcCtx.Dao.FindOrdersByUserId(l.ctx, userId, status)
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.ListVoucherOrdersByUserResp{
				Orders: []*pb.VoucherOrderInfo{},
			}, nil
		}
		l.Logger.Errorf("find orders by user id failed: %v", err)
		return nil, err
	}

	// 转换为pb格式
	orderInfos := make([]*pb.VoucherOrderInfo, 0, len(orders))
	for _, order := range orders {
		orderInfos = append(orderInfos, l.orderToProto(order))
	}

	return &pb.ListVoucherOrdersByUserResp{
		Orders: orderInfos,
	}, nil
}

// orderToProto 将model转换为proto格式
func (l *ListVoucherOrdersByUserLogic) orderToProto(order *model.TbVoucherOrder) *pb.VoucherOrderInfo {
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
