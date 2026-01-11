package logic

import (
	"context"
	"fmt"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/utils"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Follow user
func (l *FollowUserLogic) FollowUser(in *pb.FollowUserReq) (*pb.FollowUserResp, error) {
	// 1. 新增关注关系到数据库
	follow := &model.TbFollow{
		UserId:       uint64(in.UserId),
		FollowUserId: uint64(in.FollowUserId),
	}

	err := l.svcCtx.Dao.InsertFollow(l.ctx, follow)
	if err != nil {
		l.Logger.Errorf("follow user failed: %v", err)
		return &pb.FollowUserResp{
			Success: false,
		}, err
	}

	// 2. 把关注用户的id，放入redis的set集合 sadd userId followerUserId
	key := utils.FollowsKey + fmt.Sprintf("%d", in.UserId)
	_, err = l.svcCtx.Dao.SAddMembers(l.ctx, key, fmt.Sprintf("%d", in.FollowUserId))
	if err != nil {
		l.Logger.Errorf("add to redis failed: %v", err)
		// Redis失败不影响主流程，只记录日志
	}

	return &pb.FollowUserResp{
		Success: true,
	}, nil
}
