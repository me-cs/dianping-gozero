package dao

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type MySQL struct {
	voucherModel         model.TbVoucherModel
	seckillVoucherModel  model.TbSeckillVoucherModel
}

func NewMySQL(c config.Config) *MySQL {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &MySQL{
		voucherModel:        model.NewTbVoucherModel(conn, c.Cache),
		seckillVoucherModel: model.NewTbSeckillVoucherModel(conn, c.Cache),
	}
}

// InsertVoucher 插入优惠券
func (d *Dao) InsertVoucher(ctx context.Context, voucher *model.TbVoucher) error {
	_, err := d.mysql.voucherModel.Insert(ctx, voucher)
	return err
}

// InsertSeckillVoucher 插入秒杀券
func (d *Dao) InsertSeckillVoucher(ctx context.Context, seckillVoucher *model.TbSeckillVoucher) error {
	_, err := d.mysql.seckillVoucherModel.Insert(ctx, seckillVoucher)
	return err
}

// FindVoucherById 根据ID查询优惠券
func (d *Dao) FindVoucherById(ctx context.Context, id uint64) (*model.TbVoucher, error) {
	return d.mysql.voucherModel.FindOne(ctx, id)
}

// FindSeckillVoucherByVoucherId 根据优惠券ID查询秒杀券
func (d *Dao) FindSeckillVoucherByVoucherId(ctx context.Context, voucherId uint64) (*model.TbSeckillVoucher, error) {
	return d.mysql.seckillVoucherModel.FindOneByVoucherId(ctx, voucherId)
}

// FindVouchersByShopId 根据商铺ID查询优惠券列表
func (d *Dao) FindVouchersByShopId(ctx context.Context, shopId uint64) ([]*model.TbVoucher, error) {
	return d.mysql.voucherModel.FindByShopId(ctx, shopId)
}

// UpdateSeckillVoucherStock 更新秒杀券库存（扣减库存）
func (d *Dao) UpdateSeckillVoucherStock(ctx context.Context, voucherId uint64) error {
	return d.mysql.seckillVoucherModel.DecrStock(ctx, voucherId)
}
