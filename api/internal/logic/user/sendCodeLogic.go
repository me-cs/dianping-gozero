// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"math/rand"
	"regexp"
	"time"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

// isPhoneInvalid 校验手机号格式
func isPhoneInvalid(phone string) bool {
	if phone == "" {
		return true
	}
	return !phoneRegex.MatchString(phone)
}

// generateCode 生成指定长度的数字验证码
func generateCode(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	digits := make([]byte, length)
	for i := range digits {
		digits[i] = byte('0' + rnd.Intn(10))
	}
	return string(digits)
}

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeReq) (resp *types.SendCodeRsp, err error) {
	// 调用 rpc 层发送验证码
	sendCodeReq := &user.SendCodeReq{
		Phone: req.Phone,
	}

	rsp, err := l.svcCtx.UserRpc.SendCode(l.ctx, sendCodeReq)
	if err != nil {
		l.Errorf("调用 UserRpc.SendCode 失败: %v", err)
		return &types.SendCodeRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	return &types.SendCodeRsp{
		Success: rsp.GetSuccess(),
		ErrMsg:  rsp.GetMessage(),
	}, nil
}
