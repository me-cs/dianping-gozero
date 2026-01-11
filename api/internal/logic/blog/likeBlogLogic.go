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

type LikeBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeBlogLogic {
	return &LikeBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeBlogLogic) LikeBlog(req *types.LikeBlogReq) (resp *types.LikeBlogRsp, err error) {
	// 调用 rpc 层点赞博客
	likeBlogReq := &blog.LikeBlogReq{
		BlogId: req.Id,
	}

	rsp, err := l.svcCtx.BlogRpc.LikeBlog(l.ctx, likeBlogReq)
	if err != nil {
		l.Errorf("调用 BlogRpc.LikeBlog 失败: %v", err)
		return &types.LikeBlogRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	return &types.LikeBlogRsp{
		Success: rsp.GetSuccess(),
	}, nil
}
