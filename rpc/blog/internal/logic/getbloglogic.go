package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/blog/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/blog/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/blog/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type GetBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogLogic {
	return &GetBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get blog by ID
func (l *GetBlogLogic) GetBlog(in *pb.GetBlogReq) (*pb.GetBlogResp, error) {
	// 1. 查询博客
	blog, err := l.svcCtx.Dao.FindBlogById(l.ctx, uint64(in.BlogId))
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.GetBlogResp{}, nil
		}
		l.Logger.Errorf("find blog failed: %v", err)
		return nil, err
	}

	// 2. 构建响应
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

	// 3. 查询评论数（如果有值）
	if blog.Comments.Valid {
		blogInfo.Comments = int32(blog.Comments.Int64)
	}

	// 4. 判断当前用户是否点赞（如果提供了userId）
	if in.UserId > 0 {
		isLiked, err := l.svcCtx.Dao.IsLiked(l.ctx, uint64(in.BlogId), uint64(in.UserId))
		if err != nil {
			l.Logger.Errorf("check like status failed: %v", err)
		} else {
			blogInfo.IsLike = isLiked
		}
	}

	// TODO: 5. 查询博客作者信息（需要调用 User RPC）
	// 这里暂时只返回 userId，后续可以通过 User RPC 获取用户昵称和头像
	// userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userpb.GetUserReq{UserId: blogInfo.UserId})
	// if err == nil && userInfo.User != nil {
	//     blogInfo.Name = userInfo.User.NickName
	//     blogInfo.Icon = userInfo.User.Icon
	// }

	return &pb.GetBlogResp{
		Blog: blogInfo,
	}, nil
}
