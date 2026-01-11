package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get user info by ID
func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// 查询用户基本信息
	user, err := l.svcCtx.Dao.FindOneById(l.ctx, uint64(in.UserId))
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.GetUserInfoResp{}, nil
		}
		l.Logger.Errorf("find user failed: %v", err)
		return nil, err
	}

	resp := &pb.GetUserInfoResp{
		User: &pb.UserInfo{
			Id:         int64(user.Id),
			Phone:      user.Phone,
			NickName:   user.NickName,
			Icon:       user.Icon,
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdateTime.Format("2006-01-02 15:04:05"),
		},
	}

	// 查询用户详细信息
	userInfo, err := l.svcCtx.Dao.FindUserInfoById(l.ctx, uint64(in.UserId))
	if err != nil && err != sqlc.ErrNotFound && err != model.ErrNotFound {
		l.Logger.Errorf("find user info failed: %v", err)
		return nil, err
	}

	if userInfo != nil {
		birthday := ""
		if userInfo.Birthday.Valid {
			birthday = userInfo.Birthday.Time.Format("2006-01-02")
		}
		introduce := ""
		if userInfo.Introduce.Valid {
			introduce = userInfo.Introduce.String
		}

		resp.UserInfo = &pb.UserDetailInfo{
			UserId:    int64(userInfo.UserId),
			City:      userInfo.City,
			Introduce: introduce,
			Fans:      int32(userInfo.Fans),
			Followee:  int32(userInfo.Followee),
			Gender:    int32(userInfo.Gender),
			Birthday:  birthday,
			Credits:   int32(userInfo.Credits),
			Level:     int32(userInfo.Level),
		}
	}
	logx.Infof("userInfo: %v", resp.String())
	return resp, nil
}
