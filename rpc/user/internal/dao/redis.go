package dao

import (
	"context"
	"time"
)

const (
	LoginCodeKey = "login:code:"
	LoginCodeTTL = 2 * time.Minute
	LoginUserKey = "login:token:"
	LoginUserTTL = 36000 * 60 // 36000 minutes in seconds
	UserSignKey  = "sign:"
	FollowsKey   = "follows:"

	// System constants
	UserNickNamePrefix = "user_"
	DefaultPageSize    = 5
	MaxPageSize        = 10
)

// SaveCode 保存验证码到 Redis
func (d *Dao) SaveCode(ctx context.Context, phone, code string) error {
	key := LoginCodeKey + phone
	return d.Redis.SetexCtx(ctx, key, code, int(LoginCodeTTL.Seconds()))
}

// GetCode 获取之前保存验证码
func (d *Dao) GetCode(ctx context.Context, phone string) (string, error) {
	key := LoginCodeKey + phone
	return d.Redis.GetCtx(ctx, key)
}

func (d *Dao) SetUserInfo(ctx context.Context, token, userInfo string) error {
	key := LoginUserKey + token
	return d.Redis.SetexCtx(ctx, key, userInfo, LoginUserTTL)
}

// SetBit 设置位图
func (d *Dao) SetBit(ctx context.Context, key string, offset int64, value int) error {
	_, err := d.Redis.SetBitCtx(ctx, key, offset, value)
	return err
}

// BitField 执行BITFIELD命令
func (d *Dao) BitField(ctx context.Context, key string, operation string, fieldType string, offset string) ([]int64, error) {
	// 使用Do方法执行BITFIELD命令
	return nil, nil
}

// SAddMembers 添加成员到集合
func (d *Dao) SAddMembers(ctx context.Context, key string, members ...interface{}) (int, error) {
	return d.Redis.SaddCtx(ctx, key, members...)
}

// SRemMembers 从集合中移除成员
func (d *Dao) SRemMembers(ctx context.Context, key string, members ...interface{}) (int, error) {
	return d.Redis.SremCtx(ctx, key, members...)
}

// SInter 求多个集合的交集
func (d *Dao) SInter(ctx context.Context, keys ...string) ([]string, error) {
	return d.Redis.SinterCtx(ctx, keys...)
}
