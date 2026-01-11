package dao

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	// SeckillStockKey 秒杀库存key前缀
	SeckillStockKey = "seckill:stock:"
)

// GetSeckillStock 获取秒杀库存
func (d *Dao) GetSeckillStock(ctx context.Context, voucherId uint64) (int, error) {
	key := fmt.Sprintf("%s%d", SeckillStockKey, voucherId)
	val, err := d.Redis.GetCtx(ctx, key)
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}
	stock, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return stock, nil
}

// InitSeckillStock 初始化秒杀库存到Redis
func (d *Dao) InitSeckillStock(ctx context.Context, voucherId uint64, stock int) error {
	key := fmt.Sprintf("%s%d", SeckillStockKey, voucherId)
	return d.Redis.SetCtx(ctx, key, strconv.Itoa(stock))
}

// DecrSeckillStock 扣减秒杀库存（原子操作）
// 返回扣减后的库存数量
func (d *Dao) DecrSeckillStock(ctx context.Context, voucherId uint64) (int64, error) {
	key := fmt.Sprintf("%s%d", SeckillStockKey, voucherId)
	return d.Redis.DecrCtx(ctx, key)
}

// CheckSeckillStock 检查秒杀库存是否充足
func (d *Dao) CheckSeckillStock(ctx context.Context, voucherId uint64) (bool, error) {
	stock, err := d.GetSeckillStock(ctx, voucherId)
	if err != nil {
		return false, err
	}
	return stock > 0, nil
}
