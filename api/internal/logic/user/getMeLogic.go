// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/library/rpcx"
	"github.com/me-cs/dianping-gozero/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMeLogic {
	return &GetMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMeLogic) GetMe(req *types.GetMeReq) (resp *types.GetMeRsp, err error) {
	// 从context获取userId
	u := rpcx.UserFromIncomingContext(l.ctx)
	if u == nil {
		l.Errorf("获取user失败")
		return &types.GetMeRsp{
			Success: false,
			ErrMsg:  "用户信息获取失败",
		}, nil
	}
	userId := u.GetUid()

	// 调用RPC层获取用户信息
	getMeReq := &user.GetMeReq{
		UserId: userId,
	}

	rsp, err := l.svcCtx.UserRpc.GetMe(l.ctx, getMeReq)
	if err != nil {
		l.Errorf("调用 UserRpc.GetMe 失败: %v", err)
		return &types.GetMeRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}
	l.Infof("获取到用户信息，id:%v,phone:%v,nickname:%v,icon:%v", rsp.GetUser().GetId(), rsp.GetUser().GetPhone(),
		rsp.GetUser().GetNickName(), rsp.GetUser().GetIcon())

	// 检查是否获取到用户信息
	if rsp.User == nil {
		return &types.GetMeRsp{
			Success: false,
			ErrMsg:  "用户不存在",
		}, nil
	}

	// 构造返回响应
	return &types.GetMeRsp{
		Success: true,
		Data: &types.UserInfo{
			Id:       rsp.User.Id,
			Phone:    rsp.User.Phone,
			NickName: rsp.User.NickName,
			Icon:     rsp.User.Icon,
		},
	}, nil
}
