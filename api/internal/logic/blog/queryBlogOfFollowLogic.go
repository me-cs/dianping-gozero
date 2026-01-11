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

type QueryBlogOfFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryBlogOfFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBlogOfFollowLogic {
	return &QueryBlogOfFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryBlogOfFollowLogic) QueryBlogOfFollow(req *types.QueryBlogOfFollowReq) (resp *types.QueryBlogOfFollowRsp, err error) {
	// 调用 rpc 层查询关注用户的博客
	queryBlogsOfFollowReq := &blog.QueryBlogsOfFollowReq{
		LastId: req.LastId,
		Offset: req.Offset,
	}

	rsp, err := l.svcCtx.BlogRpc.QueryBlogsOfFollow(l.ctx, queryBlogsOfFollowReq)
	if err != nil {
		l.Errorf("调用 BlogRpc.QueryBlogsOfFollow 失败: %v", err)
		return &types.QueryBlogOfFollowRsp{
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

	// 构造返回数据，包含 blogs、minTime 和 offset
	data := map[string]interface{}{
		"list":    blogs,
		"minTime": rsp.GetMinTime(),
		"offset":  rsp.GetOffset(),
	}

	return &types.QueryBlogOfFollowRsp{
		Success: true,
		Data:    data,
	}, nil
}
