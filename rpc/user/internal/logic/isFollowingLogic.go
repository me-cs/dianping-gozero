package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowingLogic {
	return &IsFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Check if following
func (l *IsFollowingLogic) IsFollowing(in *pb.IsFollowingReq) (*pb.IsFollowingResp, error) {
	// 查询是否关注
	count, err := l.svcCtx.Dao.CountFollow(l.ctx, uint64(in.UserId), uint64(in.TargetUserId))
	if err != nil {
		l.Logger.Errorf("check following failed: %v", err)
		return nil, err
	}

	return &pb.IsFollowingResp{
		IsFollowing: count > 0,
	}, nil
}
