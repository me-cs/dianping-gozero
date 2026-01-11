// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/library/rpcx"
	"github.com/me-cs/dianping-gozero/rpc/blog/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveBlogLogic {
	return &SaveBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveBlogLogic) SaveBlog(req *types.SaveBlogReq) (resp *types.SaveBlogRsp, err error) {
	// 从context获取userId
	u := rpcx.UserFromIncomingContext(l.ctx)
	if u == nil {
		l.Errorf("获取user失败")
		return &types.SaveBlogRsp{
			Success: false,
			ErrMsg:  "用户信息获取失败",
		}, nil
	}
	userId := u.GetUid()
	// 调用 rpc 层创建博客
	createBlogReq := &blog.CreateBlogReq{
		UserId:  userId,
		ShopId:  req.ShopId,
		Title:   req.Title,
		Images:  req.Images,
		Content: req.Content,
	}

	rsp, err := l.svcCtx.BlogRpc.CreateBlog(l.ctx, createBlogReq)
	if err != nil {
		l.Errorf("调用 BlogRpc.CreateBlog 失败: %v", err)
		return &types.SaveBlogRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	return &types.SaveBlogRsp{
		Success: true,
		Data:    rsp.GetBlogId(),
	}, nil
}
