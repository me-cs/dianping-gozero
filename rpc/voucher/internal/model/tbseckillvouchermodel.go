package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSeckillVoucherModel = (*customTbSeckillVoucherModel)(nil)

type (
	// TbSeckillVoucherModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSeckillVoucherModel.
	TbSeckillVoucherModel interface {
		tbSeckillVoucherModel
		FindOneByVoucherId(ctx context.Context, voucherId uint64) (*TbSeckillVoucher, error)
		DecrStock(ctx context.Context, voucherId uint64) error
	}

	customTbSeckillVoucherModel struct {
		*defaultTbSeckillVoucherModel
	}
)

// NewTbSeckillVoucherModel returns a model for the database table.
func NewTbSeckillVoucherModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbSeckillVoucherModel {
	return &customTbSeckillVoucherModel{
		defaultTbSeckillVoucherModel: newTbSeckillVoucherModel(conn, c, opts...),
	}
}

// FindOneByVoucherId 根据优惠券ID查询秒杀券
func (m *customTbSeckillVoucherModel) FindOneByVoucherId(ctx context.Context, voucherId uint64) (*TbSeckillVoucher, error) {
	var resp TbSeckillVoucher
	query := fmt.Sprintf("select %s from %s where voucher_id = ? limit 1", tbSeckillVoucherRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, voucherId)
	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// DecrStock 扣减秒杀券库存（只有库存>0时才能扣减）
func (m *customTbSeckillVoucherModel) DecrStock(ctx context.Context, voucherId uint64) error {
	query := fmt.Sprintf("update %s set stock = stock - 1 where voucher_id = ? and stock > 0", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, voucherId)
	return err
}
