package dao

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/order/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/order/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type MySQL struct {
	voucherOrderModel model.TbVoucherOrderModel
}

func NewMySQL(c config.Config) *MySQL {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &MySQL{
		voucherOrderModel: model.NewTbVoucherOrderModel(conn, c.Cache),
	}
}

// InsertVoucherOrder 插入优惠券订单
func (d *Dao) InsertVoucherOrder(ctx context.Context, order *model.TbVoucherOrder) error {
	_, err := d.mysql.voucherOrderModel.Insert(ctx, order)
	return err
}

// FindVoucherOrderById 根据ID查询订单
func (d *Dao) FindVoucherOrderById(ctx context.Context, id uint64) (*model.TbVoucherOrder, error) {
	return d.mysql.voucherOrderModel.FindOne(ctx, int64(id))
}

// FindOrderByUserIdAndVoucherId 查询用户是否已经购买过该优惠券
func (d *Dao) FindOrderByUserIdAndVoucherId(ctx context.Context, userId, voucherId uint64) (*model.TbVoucherOrder, error) {
	return d.mysql.voucherOrderModel.FindOneByUserIdAndVoucherId(ctx, userId, voucherId)
}

// FindOrdersByUserId 根据用户ID查询订单列表
func (d *Dao) FindOrdersByUserId(ctx context.Context, userId uint64, status int) ([]*model.TbVoucherOrder, error) {
	return d.mysql.voucherOrderModel.FindByUserId(ctx, userId, status)
}

// UpdateOrderStatus 更新订单状态
func (d *Dao) UpdateOrderStatus(ctx context.Context, orderId uint64, status int) error {
	return d.mysql.voucherOrderModel.UpdateStatus(ctx, orderId, status)
}

// UpdateOrderStatusAndPayTime 更新订单状态和支付时间
func (d *Dao) UpdateOrderStatusAndPayTime(ctx context.Context, orderId uint64, status int, payType uint64) error {
	return d.mysql.voucherOrderModel.UpdateStatusAndPayTime(ctx, orderId, status, payType)
}

// UpdateOrderStatusAndUseTime 更新订单状态和核销时间
func (d *Dao) UpdateOrderStatusAndUseTime(ctx context.Context, orderId uint64, status int) error {
	return d.mysql.voucherOrderModel.UpdateStatusAndUseTime(ctx, orderId, status)
}
