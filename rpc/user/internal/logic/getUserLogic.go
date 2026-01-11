package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get user by ID
func (l *GetUserLogic) GetUser(in *pb.GetUserReq) (*pb.GetUserResp, error) {
	user, err := l.svcCtx.Dao.FindOneById(l.ctx, uint64(in.UserId))
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.GetUserResp{}, nil
		}
		l.Logger.Errorf("find user failed: %v", err)
		return nil, err
	}

	return &pb.GetUserResp{
		User: &pb.UserInfo{
			Id:         int64(user.Id),
			Phone:      user.Phone,
			NickName:   user.NickName,
			Icon:       user.Icon,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
