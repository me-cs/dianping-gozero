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

type AddVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVoucherLogic {
	return &AddVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVoucherLogic) AddVoucher(req *types.AddVoucherReq) (resp *types.AddVoucherRsp, err error) {
	// 调用 RPC 添加普通优惠券
	addReq := &pb.AddVoucherReq{
		ShopId:      req.ShopId,
		Title:       req.Title,
		SubTitle:    req.SubTitle,
		Rules:       req.Rules,
		PayValue:    req.PayValue,
		ActualValue: req.ActualValue,
		Type:        1, // 1=普通券
	}

	rsp, err := l.svcCtx.VoucherRpc.AddVoucher(l.ctx, addReq)
	if err != nil {
		l.Errorf("调用 VoucherRpc.AddVoucher 失败: %v", err)
		return &types.AddVoucherRsp{
			Success: false,
			ErrMsg:  "添加优惠券失败",
		}, nil
	}

	return &types.AddVoucherRsp{
		Success: true,
		Data:    rsp.GetVoucherId(),
	}, nil
}

