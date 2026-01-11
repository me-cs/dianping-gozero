package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/blog/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/blog/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/blog/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type QueryBlogsByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBlogsByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBlogsByUserLogic {
	return &QueryBlogsByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Query blogs by user - 查询某个用户的博客列表
func (l *QueryBlogsByUserLogic) QueryBlogsByUser(in *pb.QueryBlogsByUserReq) (*pb.QueryBlogsByUserResp, error) {
	// 默认每页10条
	pageSize := 10
	if in.Current <= 0 {
		in.Current = 1
	}

	// 查询用户的博客
	blogs, err := l.svcCtx.Dao.QueryBlogsByUserId(l.ctx, uint64(in.UserId), int(in.Current), pageSize)
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.QueryBlogsByUserResp{
				Blogs: []*pb.BlogInfo{},
			}, nil
		}
		l.Logger.Errorf("query blogs by user failed: %v", err)
		return nil, err
	}

	// 转换为 pb.BlogInfo
	blogInfos := make([]*pb.BlogInfo, 0, len(blogs))
	for _, blog := range blogs {
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

		blogInfos = append(blogInfos, blogInfo)
	}

	return &pb.QueryBlogsByUserResp{
		Blogs: blogInfos,
	}, nil
}
