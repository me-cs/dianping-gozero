// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package follow

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowCommonsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowCommonsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowCommonsLogic {
	return &FollowCommonsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowCommonsLogic) FollowCommons(req *types.FollowCommonsReq) (resp *types.FollowCommonsRsp, err error) {
	// 调用 rpc 层获取共同关注
	getCommonFollowsReq := &user.GetCommonFollowsReq{
		TargetUserId: req.Id,
	}

	rsp, err := l.svcCtx.UserRpc.GetCommonFollows(l.ctx, getCommonFollowsReq)
	if err != nil {
		l.Errorf("调用 UserRpc.GetCommonFollows 失败: %v", err)
		return &types.FollowCommonsRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	// 将 rpc 返回的用户列表转换为 interface{}
	return &types.FollowCommonsRsp{
		Success: true,
		Data:    rsp.GetUsers(),
	}, nil
}
