package utils

import (
	"context"
	"errors"
)

var (
	ErrUserNotLogin = errors.New("用户未登录")
)

// GetUserId 从context中获取用户ID
// 返回值：userId (int64), 是否存在 (bool)
func GetUserId(ctx context.Context) (int64, bool) {
	userId := ctx.Value("userId")
	if userId == nil {
		return 0, false
	}
	userIdInt64, ok := userId.(int64)
	return userIdInt64, ok
}

// MustGetUserId 从context中获取用户ID，如果不存在则返回错误
// 适用于必须登录的接口
func MustGetUserId(ctx context.Context) (int64, error) {
	userId, ok := GetUserId(ctx)
	if !ok {
		return 0, ErrUserNotLogin
	}
	return userId, nil
}

// GetNickName 从context中获取用户昵称
func GetNickName(ctx context.Context) (string, bool) {
	nickName := ctx.Value("nickName")
	if nickName == nil {
		return "", false
	}
	nickNameStr, ok := nickName.(string)
	return nickNameStr, ok
}

// GetIcon 从context中获取用户头像
func GetIcon(ctx context.Context) (string, bool) {
	icon := ctx.Value("icon")
	if icon == nil {
		return "", false
	}
	iconStr, ok := icon.(string)
	return iconStr, ok
}

// GetToken 从context中获取原始token
func GetToken(ctx context.Context) (string, bool) {
	token := ctx.Value("token")
	if token == nil {
		return "", false
	}
	tokenStr, ok := token.(string)
	return tokenStr, ok
}

// UserInfo 用户信息结构体
type UserInfo struct {
	UserId   int64
	NickName string
	Icon     string
	Token    string
}

// GetUserInfo 从context中获取完整的用户信息
func GetUserInfo(ctx context.Context) (*UserInfo, bool) {
	userId, ok := GetUserId(ctx)
	if !ok {
		return nil, false
	}

	userInfo := &UserInfo{
		UserId: userId,
	}

	if nickName, ok := GetNickName(ctx); ok {
		userInfo.NickName = nickName
	}

	if icon, ok := GetIcon(ctx); ok {
		userInfo.Icon = icon
	}

	if token, ok := GetToken(ctx); ok {
		userInfo.Token = token
	}

	return userInfo, true
}
