package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	// BeginTimestamp 开始时间戳 (2022-01-01 00:00:00 UTC)
	BeginTimestamp = 1640995200
	// CountBits 序列号的位数
	CountBits = 32
)

// IDGenerator Redis ID生成器
type IDGenerator struct {
	redis *redis.Redis
}

// NewIDGenerator 创建ID生成器
func NewIDGenerator(rds *redis.Redis) *IDGenerator {
	return &IDGenerator{
		redis: rds,
	}
}

// NextID 生成全局唯一ID
// keyPrefix: ID前缀，用于区分不同业务
// 返回：64位ID，格式为 [时间戳(32位)][序列号(32位)]
func (g *IDGenerator) NextID(ctx context.Context, keyPrefix string) (uint64, error) {
	// 1. 生成时间戳
	now := time.Now()
	timestamp := now.Unix() - BeginTimestamp

	// 2. 生成序列号
	// 2.1. 获取当前日期，精确到天
	date := now.Format("2006:01:02")
	// 2.2. 自增长
	key := fmt.Sprintf("icr:%s:%s", keyPrefix, date)
	count, err := g.redis.IncrCtx(ctx, key)
	if err != nil {
		return 0, err
	}

	// 3. 拼接并返回
	// 时间戳左移32位，然后与序列号进行或运算
	id := uint64(timestamp<<CountBits) | uint64(count)
	return id, nil
}
