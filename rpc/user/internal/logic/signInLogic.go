package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/utils"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// User sign in
func (l *SignInLogic) SignIn(in *pb.SignInReq) (*pb.SignInResp, error) {
	// 1. 获取当前日期
	now := time.Now()

	// 2. 拼接key  sign:userId:yyyyMM
	keySuffix := now.Format(":200601") // 格式化为 :YYYYMM
	key := utils.UserSignKey + fmt.Sprintf("%d", in.UserId) + keySuffix

	// 3. 获取今天是本月的第几天 (从1开始，需要减1变成从0开始的offset)
	dayOfMonth := now.Day()

	// 4. 写入Redis SETBIT key offset 1
	err := l.svcCtx.Dao.SetBit(l.ctx, key, int64(dayOfMonth-1), 1)
	if err != nil {
		l.Logger.Errorf("sign in failed: %v", err)
		return &pb.SignInResp{
			Success: false,
		}, err
	}

	// 5. 获取连续签到天数
	getSignCountLogic := NewGetSignCountLogic(l.ctx, l.svcCtx)
	countResp, err := getSignCountLogic.GetSignCount(&pb.GetSignCountReq{
		UserId: in.UserId,
	})
	if err != nil {
		l.Logger.Errorf("get sign count failed: %v", err)
		return &pb.SignInResp{
			Success:   true,
			SignCount: 1,
		}, nil
	}

	return &pb.SignInResp{
		Success:   true,
		SignCount: countResp.Count,
	}, nil
}
