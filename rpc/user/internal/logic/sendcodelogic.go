package logic

import (
	"context"
	"fmt"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/utils"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Send verification code
func (l *SendCodeLogic) SendCode(in *pb.SendCodeReq) (*pb.SendCodeResp, error) {
	// 1. 校验手机号
	if utils.IsPhoneInvalid(in.Phone) {
		return &pb.SendCodeResp{
			Success: false,
			Message: "手机号格式错误！",
		}, nil
	}

	// 2. 生成6位随机验证码
	code := utils.RandomNumbers(6)

	// 3. 保存验证码到Redis，设置过期时间
	err := l.svcCtx.Dao.SaveCode(l.ctx, in.GetPhone(), code)
	if err != nil {
		l.Logger.Errorf("save code to redis failed: %v", err)
		return &pb.SendCodeResp{
			Success: false,
			Message: "系统错误，请稍后重试",
		}, err
	}

	// 4. 发送验证码（这里只是记录日志，实际应该调用短信服务）
	l.Infof("发送短信验证码成功，手机号：%s，验证码：%s", in.Phone, code)

	return &pb.SendCodeResp{
		Success: true,
		Message: fmt.Sprintf("验证码发送成功：%s", code), // 生产环境不应该返回验证码
	}, nil
}
