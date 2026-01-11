// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package voucher

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/voucher/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryVoucherOfShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryVoucherOfShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryVoucherOfShopLogic {
	return &QueryVoucherOfShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryVoucherOfShopLogic) QueryVoucherOfShop(req *types.QueryVoucherOfShopReq) (resp *types.QueryVoucherOfShopRsp, err error) {
	// 调用 RPC 查询商铺的优惠券列表
	listReq := &pb.ListVouchersByShopReq{
		ShopId: req.ShopId,
	}

	rsp, err := l.svcCtx.VoucherRpc.ListVouchersByShop(l.ctx, listReq)
	if err != nil {
		l.Errorf("调用 VoucherRpc.ListVouchersByShop 失败: %v", err)
		return &types.QueryVoucherOfShopRsp{
			Success: false,
			ErrMsg:  "查询优惠券列表失败",
		}, nil
	}

	// 转换 RPC 响应为 API 响应
	vouchers := make([]types.Voucher, 0, len(rsp.GetVouchers()))
	for _, v := range rsp.GetVouchers() {
		vouchers = append(vouchers, types.Voucher{
			Id:          v.GetId(),
			ShopId:      v.GetShopId(),
			Title:       v.GetTitle(),
			SubTitle:    v.GetSubTitle(),
			Rules:       v.GetRules(),
			PayValue:    v.GetPayValue(),
			ActualValue: v.GetActualValue(),
			Type:        int(v.GetType()),
			Status:      int(v.GetStatus()),
			CreateTime:  v.GetCreateTime(),
			UpdateTime:  v.GetUpdateTime(),
		})
	}

	return &types.QueryVoucherOfShopRsp{
		Success: true,
		Data:    vouchers,
	}, nil
}

