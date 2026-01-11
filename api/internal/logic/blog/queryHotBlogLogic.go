package blog

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/blog/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHotBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHotBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHotBlogLogic {
	return &QueryHotBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryHotBlogLogic) QueryHotBlog(req *types.QueryHotBlogReq) (resp *types.QueryHotBlogRsp, err error) {
	// 调用 rpc 层查询热门博客
	queryHotBlogsReq := &blog.QueryHotBlogsReq{
		Current: int32(req.Current),
	}

	rsp, err := l.svcCtx.BlogRpc.QueryHotBlogs(l.ctx, queryHotBlogsReq)
	if err != nil {
		l.Errorf("调用 BlogRpc.QueryHotBlogs 失败: %v", err)
		return &types.QueryHotBlogRsp{
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

	return &types.QueryHotBlogRsp{
		Success: true,
		Data:    blogs,
	}, nil
}
