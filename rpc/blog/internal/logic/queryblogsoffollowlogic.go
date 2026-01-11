package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/blog/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/blog/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/blog/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type QueryBlogsOfFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBlogsOfFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBlogsOfFollowLogic {
	return &QueryBlogsOfFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Query blogs of following users - 滚动分页查询关注用户的博客Feed
func (l *QueryBlogsOfFollowLogic) QueryBlogsOfFollow(in *pb.QueryBlogsOfFollowReq) (*pb.QueryBlogsOfFollowResp, error) {
	// 1. 获取当前用户的Feed流
	// lastId 是上一次查询的最小时间戳（第一次查询时传入当前时间戳）
	// offset 是偏移量，用于处理相同时间戳的情况
	maxScore := float64(in.LastId)
	if maxScore == 0 {
		// 第一次查询，使用当前时间戳
		maxScore = float64(9999999999999) // 使用一个很大的值
	}

	// 2. 从Redis中滚动查询Feed流
	// count 默认为2（Java代码中是2）
	count := int64(2)
	blogIds, minTime, newOffset, err := l.svcCtx.Dao.ScrollFeed(
		l.ctx,
		uint64(in.UserId),
		maxScore,
		0,
		in.Offset,
		count,
	)
	if err != nil {
		l.Logger.Errorf("scroll feed failed: %v", err)
		return nil, err
	}

	// 3. 如果没有数据，直接返回
	if len(blogIds) == 0 {
		return &pb.QueryBlogsOfFollowResp{
			Blogs:   []*pb.BlogInfo{},
			MinTime: 0,
			Offset:  0,
		}, nil
	}

	// 4. 根据博客ID批量查询博客详情
	blogInfos := make([]*pb.BlogInfo, 0, len(blogIds))
	for _, blogId := range blogIds {
		blog, err := l.svcCtx.Dao.FindBlogById(l.ctx, blogId)
		if err != nil {
			if err == sqlc.ErrNotFound || err == model.ErrNotFound {
				continue
			}
			l.Logger.Errorf("find blog %d failed: %v", blogId, err)
			continue
		}

		blogInfo := &pb.BlogInfo{
			Id:         int64(blog.Id),
			ShopId:     blog.ShopId,
			UserId:     int64(blog.UserId),
			Title:      blog.Title,
			Images:     blog.Images,
			Content:    blog.Content,
			Liked:      int32(blog.Liked),
			CreateTime: blog.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: blog.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		if blog.Comments.Valid {
			blogInfo.Comments = int32(blog.Comments.Int64)
		}

		// 查询当前用户是否点赞
		if in.UserId > 0 {
			isLiked, err := l.svcCtx.Dao.IsLiked(l.ctx, blogId, uint64(in.UserId))
			if err == nil {
				blogInfo.IsLike = isLiked
			}
		}

		// TODO: 查询博客作者信息
		// userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userpb.GetUserReq{UserId: blogInfo.UserId})

		blogInfos = append(blogInfos, blogInfo)
	}

	return &pb.QueryBlogsOfFollowResp{
		Blogs:   blogInfos,
		MinTime: int64(minTime),
		Offset:  int64(newOffset),
	}, nil
}
