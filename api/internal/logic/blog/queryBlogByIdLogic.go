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

type QueryBlogByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryBlogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBlogByIdLogic {
	return &QueryBlogByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryBlogByIdLogic) QueryBlogById(req *types.QueryBlogByIdReq) (resp *types.QueryBlogByIdRsp, err error) {
	// 调用 rpc 层查询博客详情
	getBlogReq := &blog.GetBlogReq{
		BlogId: req.Id,
	}

	rsp, err := l.svcCtx.BlogRpc.GetBlog(l.ctx, getBlogReq)
	if err != nil {
		l.Errorf("调用 BlogRpc.GetBlog 失败: %v", err)
		return &types.QueryBlogByIdRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	// 转换 pb.BlogInfo 为 types.Blog
	blogInfo := rsp.GetBlog()
	if blogInfo == nil {
		return &types.QueryBlogByIdRsp{
			Success: false,
			ErrMsg:  "博客不存在",
		}, nil
	}

	return &types.QueryBlogByIdRsp{
		Success: true,
		Data: &types.Blog{
			Id:         blogInfo.Id,
			ShopId:     blogInfo.ShopId,
			UserId:     blogInfo.UserId,
			Title:      blogInfo.Title,
			Images:     blogInfo.Images,
			Content:    blogInfo.Content,
			Liked:      int(blogInfo.Liked),
			Comments:   int(blogInfo.Comments),
			CreateTime: blogInfo.CreateTime,
			UpdateTime: blogInfo.UpdateTime,
			Icon:       blogInfo.Icon,
			Name:       blogInfo.Name,
			IsLike:     blogInfo.IsLike,
		},
	}, nil
}
