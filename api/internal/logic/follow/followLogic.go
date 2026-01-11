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

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowRsp, err error) {
	// 根据 isFollow 决定调用关注还是取关接口
	if req.IsFollow {
		// 关注用户
		followReq := &user.FollowUserReq{
			FollowUserId: req.Id,
		}

		rsp, err := l.svcCtx.UserRpc.FollowUser(l.ctx, followReq)
		if err != nil {
			l.Errorf("调用 UserRpc.FollowUser 失败: %v", err)
			return &types.FollowRsp{
				Success: false,
				ErrMsg:  "系统错误",
			}, err
		}

		return &types.FollowRsp{
			Success: rsp.GetSuccess(),
		}, nil
	} else {
		// 取消关注
		unfollowReq := &user.UnfollowUserReq{
			FollowUserId: req.Id,
		}

		rsp, err := l.svcCtx.UserRpc.UnfollowUser(l.ctx, unfollowReq)
		if err != nil {
			l.Errorf("调用 UserRpc.UnfollowUser 失败: %v", err)
			return &types.FollowRsp{
				Success: false,
				ErrMsg:  "系统错误",
			}, err
		}

		return &types.FollowRsp{
			Success: rsp.GetSuccess(),
		}, nil
	}
}
