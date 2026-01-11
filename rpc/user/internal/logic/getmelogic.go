package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMeLogic {
	return &GetMeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get current user info
func (l *GetMeLogic) GetMe(in *pb.GetMeReq) (*pb.GetMeResp, error) {
	// 直接复用 GetUserInfo 逻辑
	getUserInfoLogic := NewGetUserInfoLogic(l.ctx, l.svcCtx)
	userInfoResp, err := getUserInfoLogic.GetUserInfo(&pb.GetUserInfoReq{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetMeResp{
		User:     userInfoResp.User,
		UserInfo: userInfoResp.UserInfo,
	}, nil
}
