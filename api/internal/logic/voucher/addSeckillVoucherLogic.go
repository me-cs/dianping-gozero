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

type AddSeckillVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSeckillVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillVoucherLogic {
	return &AddSeckillVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSeckillVoucherLogic) AddSeckillVoucher(req *types.AddSeckillVoucherReq) (resp *types.AddSeckillVoucherRsp, err error) {
	// Step 1: 先创建基础优惠券
	addVoucherReq := &pb.AddVoucherReq{
		ShopId:      req.ShopId,
		Title:       req.Title,
		SubTitle:    req.SubTitle,
		Rules:       req.Rules,
		PayValue:    req.PayValue,
		ActualValue: req.ActualValue,
		Type:        2, // 2=秒杀券
	}

	voucherRsp, err := l.svcCtx.VoucherRpc.AddVoucher(l.ctx, addVoucherReq)
	if err != nil {
		l.Errorf("调用 VoucherRpc.AddVoucher 失败: %v", err)
		return &types.AddSeckillVoucherRsp{
			Success: false,
			ErrMsg:  "创建优惠券失败",
		}, nil
	}

	// Step 2: 添加秒杀信息
	addSeckillReq := &pb.AddSeckillVoucherReq{
		VoucherId: voucherRsp.GetVoucherId(),
		Stock:     int32(req.Stock),
		BeginTime: req.BeginTime,
		EndTime:   req.EndTime,
	}

	seckillRsp, err := l.svcCtx.VoucherRpc.AddSeckillVoucher(l.ctx, addSeckillReq)
	if err != nil {
		l.Errorf("调用 VoucherRpc.AddSeckillVoucher 失败: %v", err)
		return &types.AddSeckillVoucherRsp{
			Success: false,
			ErrMsg:  "添加秒杀信息失败",
		}, nil
	}

	if !seckillRsp.GetSuccess() {
		return &types.AddSeckillVoucherRsp{
			Success: false,
			ErrMsg:  "添加秒杀优惠券失败",
		}, nil
	}

	return &types.AddSeckillVoucherRsp{
		Success: true,
		Data:    voucherRsp.GetVoucherId(),
	}, nil
}

