// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignCountLogic {
	return &SignCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignCountLogic) SignCount(req *types.SignCountReq) (resp *types.SignCountRsp, err error) {
	// 调用 rpc 层获取签到统计
	getSignCountReq := &user.GetSignCountReq{}

	rsp, err := l.svcCtx.UserRpc.GetSignCount(l.ctx, getSignCountReq)
	if err != nil {
		l.Errorf("调用 UserRpc.GetSignCount 失败: %v", err)
		return &types.SignCountRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	return &types.SignCountRsp{
		Success: true,
		Data:    int(rsp.GetCount()),
	}, nil
}
