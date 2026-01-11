package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbVoucherModel = (*customTbVoucherModel)(nil)

type (
	// TbVoucherModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbVoucherModel.
	TbVoucherModel interface {
		tbVoucherModel
		FindByShopId(ctx context.Context, shopId uint64) ([]*TbVoucher, error)
	}

	customTbVoucherModel struct {
		*defaultTbVoucherModel
	}
)

// NewTbVoucherModel returns a model for the database table.
func NewTbVoucherModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbVoucherModel {
	return &customTbVoucherModel{
		defaultTbVoucherModel: newTbVoucherModel(conn, c, opts...),
	}
}

// FindByShopId 根据商铺ID查询优惠券列表
func (m *customTbVoucherModel) FindByShopId(ctx context.Context, shopId uint64) ([]*TbVoucher, error) {
	var vouchers []*TbVoucher
	query := fmt.Sprintf("select %s from %s where shop_id = ? order by create_time desc", tbVoucherRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &vouchers, query, shopId)
	switch err {
	case nil:
		return vouchers, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
