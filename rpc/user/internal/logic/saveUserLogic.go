package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserLogic {
	return &SaveUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create or update user
func (l *SaveUserLogic) SaveUser(in *pb.SaveUserReq) (*pb.SaveUserResp, error) {
	user := &model.TbUser{
		Phone:    in.Phone,
		NickName: in.NickName,
		Icon:     in.Icon,
	}

	err := l.svcCtx.Dao.InsertUserInfo(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("insert user failed: %v", err)
		return nil, err
	}

	// 重新查询获取ID
	savedUser, err := l.svcCtx.Dao.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		l.Logger.Errorf("find user after insert failed: %v", err)
		return nil, err
	}

	return &pb.SaveUserResp{
		UserId: int64(savedUser.Id),
	}, nil
}
