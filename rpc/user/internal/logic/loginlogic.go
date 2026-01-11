package logic

import (
	"context"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/utils"
	"github.com/me-cs/dianping-gozero/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// User login
func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 1. 校验手机号
	if utils.IsPhoneInvalid(in.Phone) {
		return &pb.LoginResp{
			Success: false,
			Message: "手机号格式错误！",
		}, nil
	}

	// 2. 从Redis获取验证码并校验
	cacheCode, err := l.svcCtx.Dao.GetCode(l.ctx, in.GetPhone())
	if err != nil {
		return &pb.LoginResp{
			Success: false,
			Message: "验证码已过期",
		}, nil
	}

	if cacheCode != in.Code {
		return &pb.LoginResp{
			Success: false,
			Message: "验证码错误",
		}, nil
	}

	// 3. 根据手机号查询用户
	user, err := l.svcCtx.Dao.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && err != sqlc.ErrNotFound && err != model.ErrNotFound {
		l.Logger.Errorf("query user by phone failed: %v", err)
		return &pb.LoginResp{
			Success: false,
			Message: "系统错误",
		}, err
	}

	// 4. 如果用户不存在，创建新用户
	if user == nil {
		user, err = l.createUserWithPhone(in.Phone)
		if err != nil {
			l.Logger.Errorf("create user failed: %v", err)
			return &pb.LoginResp{
				Success: false,
				Message: "创建用户失败",
			}, err
		}
		if user == nil {
			return &pb.LoginResp{
				Success: false,
				Message: "创建用户失败",
			}, nil
		}
	}

	// 5. 生成token
	// gen token
	token, err := getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire)
	if err != nil {
		l.Logger.Errorf("getJwtToken failed: %v", err)
		return &pb.LoginResp{
			Success: false,
			Message: "生成token错误",
		}, err
	}
	// 6. 将用户信息保存到Redis
	userInfo := map[string]interface{}{
		"id":       user.Id,
		"nickName": user.NickName,
		"icon":     user.Icon,
	}
	userJSON, _ := json.Marshal(userInfo)
	err = l.svcCtx.Dao.SetUserInfo(l.ctx, token, string(userJSON))
	if err != nil {
		l.Logger.Errorf("save user to redis failed: %v", err)
		return &pb.LoginResp{
			Success: false,
			Message: "系统错误",
		}, err
	}

	// 7. 返回token
	return &pb.LoginResp{
		Success: true,
		Token:   token,
		Message: "登录成功",
	}, nil
}

// createUserWithPhone 创建用户
func (l *LoginLogic) createUserWithPhone(phone string) (*model.TbUser, error) {
	user := &model.TbUser{
		Phone:    phone,
		NickName: utils.UserNickNamePrefix + utils.RandomString(10),
	}

	err := l.svcCtx.Dao.InsertUserInfo(l.ctx, user)
	if err != nil {
		return nil, err
	}

	// 重新查询用户获取ID
	user, err = l.svcCtx.Dao.FindOneByPhone(l.ctx, phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getJwtToken(secretKey string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
