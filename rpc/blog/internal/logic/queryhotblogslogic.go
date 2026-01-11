package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/blog/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/blog/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/blog/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type QueryHotBlogsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHotBlogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHotBlogsLogic {
	return &QueryHotBlogsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Query hot blogs - 查询热门博客（按点赞数排序）
func (l *QueryHotBlogsLogic) QueryHotBlogs(in *pb.QueryHotBlogsReq) (*pb.QueryHotBlogsResp, error) {
	// 默认每页10条
	pageSize := 10
	if in.Current <= 0 {
		in.Current = 1
	}

	// 查询热门博客
	blogs, err := l.svcCtx.Dao.QueryHotBlogs(l.ctx, int(in.Current), pageSize)
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.QueryHotBlogsResp{
				Blogs: []*pb.BlogInfo{},
			}, nil
		}
		l.Logger.Errorf("query hot blogs failed: %v", err)
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

		// TODO: 查询用户信息和点赞状态
		// 这里可以批量调用 User RPC 获取用户信息

		blogInfos = append(blogInfos, blogInfo)
	}

	return &pb.QueryHotBlogsResp{
		Blogs: blogInfos,
	}, nil
}
