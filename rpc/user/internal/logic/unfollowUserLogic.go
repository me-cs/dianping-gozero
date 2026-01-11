package logic

import (
	"context"
	"fmt"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/utils"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowUserLogic {
	return &UnfollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Unfollow user
func (l *UnfollowUserLogic) UnfollowUser(in *pb.UnfollowUserReq) (*pb.UnfollowUserResp, error) {
	// 1. 删除关注关系
	err := l.svcCtx.Dao.DeleteFollowByUserIdAndFollowUserId(l.ctx, uint64(in.UserId), uint64(in.FollowUserId))
	if err != nil {
		l.Logger.Errorf("unfollow user failed: %v", err)
		return &pb.UnfollowUserResp{
			Success: false,
		}, err
	}

	// 2. 把关注用户的id从Redis集合中移除
	key := utils.FollowsKey + fmt.Sprintf("%d", in.UserId)
	_, err = l.svcCtx.Dao.SRemMembers(l.ctx, key, fmt.Sprintf("%d", in.FollowUserId))
	if err != nil {
		l.Logger.Errorf("remove from redis failed: %v", err)
		// Redis失败不影响主流程，只记录日志
	}

	return &pb.UnfollowUserResp{
		Success: true,
	}, nil
}
