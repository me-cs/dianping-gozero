package logic

import (
	"context"
	"database/sql"
	"time"

	"github.com/me-cs/dianping-gozero/rpc/blog/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/blog/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/blog/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBlogLogic {
	return &CreateBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create blog - 创建博客并推送到粉丝的Feed流
func (l *CreateBlogLogic) CreateBlog(in *pb.CreateBlogReq) (*pb.CreateBlogResp, error) {
	// 1. 创建博客对象
	now := time.Now()
	blog := &model.TbBlog{
		ShopId:     in.ShopId,
		UserId:     uint64(in.UserId),
		Title:      in.Title,
		Images:     in.Images,
		Content:    in.Content,
		Liked:      0,
		Comments:   sql.NullInt64{Valid: false},
		CreateTime: now,
		UpdateTime: now,
	}

	// 2. 保存博客到数据库
	err := l.svcCtx.Dao.InsertBlog(l.ctx, blog)
	if err != nil {
		l.Logger.Errorf("insert blog failed: %v", err)
		return nil, err
	}

	// 3. TODO: 查询博客作者的所有粉丝并推送到他们的Feed流
	// 需要调用 User RPC 获取粉丝列表
	// followers, err := l.svcCtx.UserRpc.GetFollowers(l.ctx, &userpb.GetFollowersReq{UserId: in.UserId})
	// if err == nil && len(followers.UserIds) > 0 {
	//     for _, followerId := range followers.UserIds {
	//         // 推送博客到粉丝的Feed流
	//         err = l.svcCtx.Dao.PushFeed(l.ctx, uint64(followerId), blog.Id)
	//         if err != nil {
	//             l.Logger.Errorf("push feed to follower %d failed: %v", followerId, err)
	//         }
	//     }
	// }

	return &pb.CreateBlogResp{
		BlogId: int64(blog.Id),
	}, nil
}
