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

type QueryUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserByIdLogic {
	return &QueryUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserByIdLogic) QueryUserById(req *types.QueryUserByIdReq) (resp *types.QueryUserByIdRsp, err error) {
	// 调用RPC层获取用户基本信息
	getUserReq := &user.GetUserReq{
		UserId: req.UserId,
	}

	rsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, getUserReq)
	if err != nil {
		l.Errorf("调用 UserRpc.GetUser 失败: %v", err)
		return &types.QueryUserByIdRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	// 检查是否获取到用户信息
	if rsp.User == nil {
		l.Infof("用户不存在，userId:%v", req.UserId)
		return &types.QueryUserByIdRsp{
			Success: true,
			Data:    nil,
		}, nil
	}

	l.Infof("获取到用户信息，userId:%v, phone:%v, nickname:%v",
		rsp.User.Id, rsp.User.Phone, rsp.User.NickName)

	// 构造返回响应
	return &types.QueryUserByIdRsp{
		Success: true,
		Data: &types.UserInfo{
			Id:       rsp.User.Id,
			Phone:    rsp.User.Phone,
			NickName: rsp.User.NickName,
			Icon:     rsp.User.Icon,
		},
	}, nil
}
