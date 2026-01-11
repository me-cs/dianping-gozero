package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/utils"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type GetCommonFollowsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommonFollowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommonFollowsLogic {
	return &GetCommonFollowsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get common follows
func (l *GetCommonFollowsLogic) GetCommonFollows(in *pb.GetCommonFollowsReq) (*pb.GetCommonFollowsResp, error) {
	// 1. 获取两个用户的关注列表key
	key1 := utils.FollowsKey + fmt.Sprintf("%d", in.UserId)
	key2 := utils.FollowsKey + fmt.Sprintf("%d", in.TargetUserId)

	// 2. 求交集
	intersect, err := l.svcCtx.Dao.SInter(l.ctx, key1, key2)
	if err != nil || len(intersect) == 0 {
		return &pb.GetCommonFollowsResp{
			Users: []*pb.UserInfo{},
		}, nil
	}

	// 3. 解析id集合，查询用户信息
	users := make([]*pb.UserInfo, 0, len(intersect))
	for _, idStr := range intersect {
		userId, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			continue
		}

		user, err := l.svcCtx.Dao.FindOneById(l.ctx, userId)
		if err != nil {
			if err == sqlc.ErrNotFound || err == model.ErrNotFound {
				continue
			}
			l.Logger.Errorf("find user failed: %v", err)
			continue
		}

		users = append(users, &pb.UserInfo{
			Id:         int64(user.Id),
			Phone:      user.Phone,
			NickName:   user.NickName,
			Icon:       user.Icon,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.GetCommonFollowsResp{
		Users: users,
	}, nil
}
