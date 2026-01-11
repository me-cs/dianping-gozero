package dao

import (
	"context"
	"fmt"
	"strconv"
)

const (
	// SeckillStockKey 秒杀库存key
	SeckillStockKey = "seckill:stock:"

	// SeckillOrderKey 秒杀订单记录key（用于判断用户是否已下单）
	SeckillOrderKey = "seckill:order:"
)

// Lua脚本：秒杀下单（判断库存、判断重复下单、发送消息队列）
// 参数：
// ARGV[1]: 优惠券ID
// ARGV[2]: 用户ID
// ARGV[3]: 订单ID
// 返回值：
// 0: 成功
// 1: 库存不足
// 2: 用户已下单
const seckillLuaScript = `
-- 1.参数列表
local voucherId = ARGV[1]
local userId = ARGV[2]
local orderId = ARGV[3]

-- 2.数据key
local stockKey = 'seckill:stock:' .. voucherId
local orderKey = 'seckill:order:' .. voucherId

-- 3.脚本业务
-- 3.1.判断库存是否充足
if(tonumber(redis.call('get', stockKey)) <= 0) then
    return 1  -- 库存不足
end

-- 3.2.判断用户是否下单
if(redis.call('sismember', orderKey, userId) == 1) then
    return 2  -- 用户已下单
end

-- 3.3.扣库存
redis.call('incrby', stockKey, -1)

-- 3.4.下单（保存用户）
redis.call('sadd', orderKey, userId)

-- 3.5.发送消息到队列中
redis.call('xadd', 'stream.orders', '*', 'userId', userId, 'voucherId', voucherId, 'id', orderId)

return 0  -- 成功
`

// SeckillVoucher 执行秒杀（使用Lua脚本保证原子性）
// 返回值：0=成功，1=库存不足，2=用户已下单
func (d *Dao) SeckillVoucher(ctx context.Context, voucherId, userId, orderId uint64) (int64, error) {
	voucherIdStr := strconv.FormatUint(voucherId, 10)
	userIdStr := strconv.FormatUint(userId, 10)
	orderIdStr := strconv.FormatUint(orderId, 10)

	// 执行Lua脚本（不需要KEYS，直接在脚本内构造）
	result, err := d.Redis.EvalCtx(ctx, seckillLuaScript, []string{}, voucherIdStr, userIdStr, orderIdStr)
	if err != nil {
		return -1, err
	}

	// 转换结果
	code, ok := result.(int64)
	if !ok {
		return -1, fmt.Errorf("unexpected result type")
	}

	return code, nil
}

// CheckUserOrdered 检查用户是否已下单
func (d *Dao) CheckUserOrdered(ctx context.Context, voucherId, userId uint64) (bool, error) {
	orderKey := fmt.Sprintf("%s%d", SeckillOrderKey, voucherId)
	userIdStr := strconv.FormatUint(userId, 10)

	return d.Redis.SismemberCtx(ctx, orderKey, userIdStr)
}

// AddOrderRecord 添加订单记录到Redis Set
func (d *Dao) AddOrderRecord(ctx context.Context, voucherId, userId uint64) error {
	orderKey := fmt.Sprintf("%s%d", SeckillOrderKey, voucherId)
	userIdStr := strconv.FormatUint(userId, 10)

	_, err := d.Redis.SaddCtx(ctx, orderKey, userIdStr)
	return err
}
