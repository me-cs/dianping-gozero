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

type SignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignLogic {
	return &SignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignLogic) Sign(req *types.SignReq) (resp *types.SignRsp, err error) {
	// todo: add your logic here and delete this line
	signinReq := user.SignInReq{}
	rsp, err := l.svcCtx.UserRpc.SignIn(l.ctx, &signinReq)
	if err != nil {
		return &types.SignRsp{Success: false}, err
	}
	return &types.SignRsp{Success: rsp.GetSuccess()}, err
}
