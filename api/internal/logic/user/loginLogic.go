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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRsp, err error) {
	// 调用 rpc 层登录
	loginReq := &user.LoginReq{
		Phone: req.Phone,
		Code:  req.Code,
	}

	rsp, err := l.svcCtx.UserRpc.Login(l.ctx, loginReq)
	if err != nil {
		l.Errorf("调用 UserRpc.Login 失败: %v", err)
		return &types.LoginRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}
	l.Infof("登录成功: success=%v, token=%s, message=%s", rsp.GetSuccess(), rsp.GetToken(), rsp.GetMessage())

	return &types.LoginRsp{
		Success: rsp.GetSuccess(),
		Data:    rsp.GetToken(),
		ErrMsg:  rsp.GetMessage(),
	}, nil
}
