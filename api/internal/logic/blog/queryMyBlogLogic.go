// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/blog/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMyBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMyBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMyBlogLogic {
	return &QueryMyBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMyBlogLogic) QueryMyBlog(req *types.QueryMyBlogReq) (resp *types.QueryMyBlogRsp, err error) {
	// 调用 rpc 层查询我的博客（当前用户的博客）
	queryBlogsByUserReq := &blog.QueryBlogsByUserReq{
		Current: int32(req.Current),
	}

	rsp, err := l.svcCtx.BlogRpc.QueryBlogsByUser(l.ctx, queryBlogsByUserReq)
	if err != nil {
		l.Errorf("调用 BlogRpc.QueryBlogsByUser 失败: %v", err)
		return &types.QueryMyBlogRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	// 转换 []*pb.BlogInfo 为 []types.Blog
	pbBlogs := rsp.GetBlogs()
	blogs := make([]types.Blog, 0, len(pbBlogs))
	for _, pbBlog := range pbBlogs {
		blogs = append(blogs, types.Blog{
			Id:         pbBlog.Id,
			ShopId:     pbBlog.ShopId,
			UserId:     pbBlog.UserId,
			Title:      pbBlog.Title,
			Images:     pbBlog.Images,
			Content:    pbBlog.Content,
			Liked:      int(pbBlog.Liked),
			Comments:   int(pbBlog.Comments),
			CreateTime: pbBlog.CreateTime,
			UpdateTime: pbBlog.UpdateTime,
			Icon:       pbBlog.Icon,
			Name:       pbBlog.Name,
			IsLike:     pbBlog.IsLike,
		})
	}

	return &types.QueryMyBlogRsp{
		Success: true,
		Data:    blogs,
	}, nil
}
