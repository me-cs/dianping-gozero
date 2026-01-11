package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/blog/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/blog/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlogLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlogLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogLikesLogic {
	return &GetBlogLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get blog likes - 获取前5个点赞的用户ID
func (l *GetBlogLikesLogic) GetBlogLikes(in *pb.GetBlogLikesReq) (*pb.GetBlogLikesResp, error) {
	// 查询top5的点赞用户
	userIds, err := l.svcCtx.Dao.GetTopLikers(l.ctx, uint64(in.BlogId), 5)
	if err != nil {
		l.Logger.Errorf("get top likers failed: %v", err)
		return nil, err
	}

	// 转换为 int64 切片
	likerIds := make([]int64, 0, len(userIds))
	for _, uid := range userIds {
		likerIds = append(likerIds, int64(uid))
	}

	// TODO: 可以调用 User RPC 批量获取用户信息
	// userInfos, err := l.svcCtx.UserRpc.BatchGetUsers(l.ctx, &userpb.BatchGetUsersReq{UserIds: likerIds})

	return &pb.GetBlogLikesResp{
		UserIds: likerIds,
	}, nil
}
