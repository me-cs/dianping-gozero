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

type GetSignCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSignCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSignCountLogic {
	return &GetSignCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get sign in count
func (l *GetSignCountLogic) GetSignCount(in *pb.GetSignCountReq) (*pb.GetSignCountResp, error) {
	// 1. 获取当前日期
	now := time.Now()

	// 2. 拼接key  sign:userId:yyyyMM
	keySuffix := now.Format(":200601") // 格式化为 :YYYYMM
	key := utils.UserSignKey + fmt.Sprintf("%d", in.UserId) + keySuffix

	// 3. 获取今天是本月的第几天
	dayOfMonth := now.Day()

	// 4. 获取本月截止今天为止的所有的签到记录
	// 使用BITFIELD命令: BITFIELD key GET u{dayOfMonth} 0
	result, err := l.svcCtx.Dao.BitField(l.ctx, key, "GET", fmt.Sprintf("u%d", dayOfMonth), "0")
	if err != nil {
		l.Logger.Errorf("get sign count failed: %v", err)
		return &pb.GetSignCountResp{
			Count: 0,
		}, nil
	}

	if len(result) == 0 {
		return &pb.GetSignCountResp{
			Count: 0,
		}, nil
	}

	num := result[0]
	if num == 0 {
		return &pb.GetSignCountResp{
			Count: 0,
		}, nil
	}

	// 5. 循环遍历，统计连续签到天数
	count := int32(0)
	for {
		// 让这个数字与1做与运算，得到数字的最后一个bit位
		if (num & 1) == 0 {
			// 如果为0，说明未签到，结束
			break
		} else {
			// 如果不为0，说明已签到，计数器+1
			count++
		}
		// 把数字右移一位，抛弃最后一个bit位，继续下一个bit位
		num >>= 1
	}

	return &pb.GetSignCountResp{
		Count: count,
	}, nil
}
