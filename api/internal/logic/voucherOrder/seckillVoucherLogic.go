// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package voucherOrder

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillVoucherLogic {
	return &SeckillVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillVoucherLogic) SeckillVoucher(req *types.SeckillVoucherReq) (resp *types.SeckillVoucherRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
