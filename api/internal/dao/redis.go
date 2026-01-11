package dao

import (
	"context"
	"time"
)

const (
	LoginCodeKey = "login:code:"
	LoginCodeTTL = 2 * time.Minute
)

// SaveCode 保存验证码到 Redis
func (d *Dao) SaveCode(ctx context.Context, phone, code string) error {
	key := LoginCodeKey + phone
	return d.Redis.SetexCtx(ctx, key, code, int(LoginCodeTTL.Seconds()))
}
