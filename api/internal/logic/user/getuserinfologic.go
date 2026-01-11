// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoRsp, err error) {
	// 调用RPC层获取用户详细信息
	getUserInfoReq := &user.GetUserInfoReq{
		UserId: req.UserId,
	}

	rsp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, getUserInfoReq)
	if err != nil {
		l.Errorf("调用 UserRpc.GetUserInfo 失败: %v", err)
		return &types.GetUserInfoRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	// 检查是否获取到用户信息
	if rsp.UserInfo == nil {
		l.Infof("用户详情不存在，userId:%v", req.UserId)
		return &types.GetUserInfoRsp{
			Success: true,
			Data:    nil,
		}, nil
	}

	l.Infof("获取到用户详情，userId:%v", req.UserId)

	// 构造返回响应
	return &types.GetUserInfoRsp{
		Success: true,
		Data: &types.UserDetailInfo{
			Id:        rsp.UserInfo.UserId,
			City:      rsp.UserInfo.City,
			Introduce: rsp.UserInfo.Introduce,
			Fans:      int(rsp.UserInfo.Fans),
			Followee:  int(rsp.UserInfo.Followee),
			Gender:    int(rsp.UserInfo.Gender),
			Birthday:  rsp.UserInfo.Birthday,
			Credits:   int(rsp.UserInfo.Credits),
			Level:     int(rsp.UserInfo.Level),
		},
	}, nil
}
