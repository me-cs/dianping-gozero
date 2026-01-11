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

type IsFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowLogic {
	return &IsFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsFollowLogic) IsFollow(req *types.IsFollowReq) (resp *types.IsFollowRsp, err error) {
	// 调用 rpc 层检查是否关注
	isFollowingReq := &user.IsFollowingReq{
		TargetUserId: req.Id,
	}

	rsp, err := l.svcCtx.UserRpc.IsFollowing(l.ctx, isFollowingReq)
	if err != nil {
		l.Errorf("调用 UserRpc.IsFollowing 失败: %v", err)
		return &types.IsFollowRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	return &types.IsFollowRsp{
		Success: true,
		Data:    rsp.GetIsFollowing(),
	}, nil
}
