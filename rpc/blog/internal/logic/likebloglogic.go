package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/blog/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/blog/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeBlogLogic {
	return &LikeBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Like blog
func (l *LikeBlogLogic) LikeBlog(in *pb.LikeBlogReq) (*pb.LikeBlogResp, error) {
	blogId := uint64(in.BlogId)
	userId := uint64(in.UserId)

	// 1. 判断当前用户是否已经点赞
	isLiked, err := l.svcCtx.Dao.IsLiked(l.ctx, blogId, userId)
	if err != nil {
		l.Logger.Errorf("check like status failed: %v", err)
		return nil, err
	}

	// 2. 查询博客（用于更新点赞数）
	blog, err := l.svcCtx.Dao.FindBlogById(l.ctx, blogId)
	if err != nil {
		l.Logger.Errorf("find blog failed: %v", err)
		return nil, err
	}

	if !isLiked {
		// 3. 如果未点赞，执行点赞操作
		// 3.1 数据库点赞数 +1
		blog.Liked++
		err = l.svcCtx.Dao.UpdateBlog(l.ctx, blog)
		if err != nil {
			l.Logger.Errorf("update blog liked count failed: %v", err)
			return nil, err
		}

		// 3.2 保存用户到Redis的ZSet集合
		err = l.svcCtx.Dao.LikeBlog(l.ctx, blogId, userId)
		if err != nil {
			l.Logger.Errorf("redis like blog failed: %v", err)
			// 回滚数据库操作
			blog.Liked--
			_ = l.svcCtx.Dao.UpdateBlog(l.ctx, blog)
			return nil, err
		}
	} else {
		// 4. 如果已点赞，取消点赞
		// 4.1 数据库点赞数 -1
		if blog.Liked > 0 {
			blog.Liked--
		}
		err = l.svcCtx.Dao.UpdateBlog(l.ctx, blog)
		if err != nil {
			l.Logger.Errorf("update blog liked count failed: %v", err)
			return nil, err
		}

		// 4.2 把用户从Redis的ZSet集合移除
		err = l.svcCtx.Dao.UnlikeBlog(l.ctx, blogId, userId)
		if err != nil {
			l.Logger.Errorf("redis unlike blog failed: %v", err)
			// 回滚数据库操作
			blog.Liked++
			_ = l.svcCtx.Dao.UpdateBlog(l.ctx, blog)
			return nil, err
		}
	}

	return &pb.LikeBlogResp{
		Success: true,
	}, nil
}
