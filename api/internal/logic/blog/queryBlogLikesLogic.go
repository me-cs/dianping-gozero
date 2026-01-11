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

type QueryBlogLikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryBlogLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBlogLikesLogic {
	return &QueryBlogLikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryBlogLikesLogic) QueryBlogLikes(req *types.QueryBlogLikesReq) (resp *types.QueryBlogLikesRsp, err error) {
	// 调用 rpc 层查询博客点赞列表
	getBlogLikesReq := &blog.GetBlogLikesReq{
		BlogId: req.Id,
	}

	rsp, err := l.svcCtx.BlogRpc.GetBlogLikes(l.ctx, getBlogLikesReq)
	if err != nil {
		l.Errorf("调用 BlogRpc.GetBlogLikes 失败: %v", err)
		return &types.QueryBlogLikesRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	return &types.QueryBlogLikesRsp{
		Success: true,
		Data:    rsp.GetUserIds(),
	}, nil
}
