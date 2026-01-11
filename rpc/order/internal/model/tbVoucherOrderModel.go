package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbVoucherOrderModel = (*customTbVoucherOrderModel)(nil)

type (
	// TbVoucherOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbVoucherOrderModel.
	TbVoucherOrderModel interface {
		tbVoucherOrderModel
		FindOneByUserIdAndVoucherId(ctx context.Context, userId, voucherId uint64) (*TbVoucherOrder, error)
		FindByUserId(ctx context.Context, userId uint64, status int) ([]*TbVoucherOrder, error)
		UpdateStatus(ctx context.Context, orderId uint64, status int) error
		UpdateStatusAndPayTime(ctx context.Context, orderId uint64, status int, payType uint64) error
		UpdateStatusAndUseTime(ctx context.Context, orderId uint64, status int) error
	}

	customTbVoucherOrderModel struct {
		*defaultTbVoucherOrderModel
	}
)

// NewTbVoucherOrderModel returns a model for the database table.
func NewTbVoucherOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbVoucherOrderModel {
	return &customTbVoucherOrderModel{
		defaultTbVoucherOrderModel: newTbVoucherOrderModel(conn, c, opts...),
	}
}

// FindOneByUserIdAndVoucherId 查询用户是否已购买该优惠券
func (m *customTbVoucherOrderModel) FindOneByUserIdAndVoucherId(ctx context.Context, userId, voucherId uint64) (*TbVoucherOrder, error) {
	var resp TbVoucherOrder
	query := fmt.Sprintf("select %s from %s where user_id = ? and voucher_id = ? limit 1", tbVoucherOrderRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, voucherId)
	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindByUserId 根据用户ID查询订单列表
func (m *customTbVoucherOrderModel) FindByUserId(ctx context.Context, userId uint64, status int) ([]*TbVoucherOrder, error) {
	var orders []*TbVoucherOrder
	var query string
	var err error

	if status > 0 {
		query = fmt.Sprintf("select %s from %s where user_id = ? and status = ? order by create_time desc", tbVoucherOrderRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &orders, query, userId, status)
	} else {
		query = fmt.Sprintf("select %s from %s where user_id = ? order by create_time desc", tbVoucherOrderRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &orders, query, userId)
	}

	switch err {
	case nil:
		return orders, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// UpdateStatus 更新订单状态
func (m *customTbVoucherOrderModel) UpdateStatus(ctx context.Context, orderId uint64, status int) error {
	query := fmt.Sprintf("update %s set status = ? where id = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, status, orderId)
	return err
}

// UpdateStatusAndPayTime 更新订单状态和支付时间
func (m *customTbVoucherOrderModel) UpdateStatusAndPayTime(ctx context.Context, orderId uint64, status int, payType uint64) error {
	query := fmt.Sprintf("update %s set status = ?, pay_type = ?, pay_time = NOW() where id = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, status, payType, orderId)
	return err
}

// UpdateStatusAndUseTime 更新订单状态和核销时间
func (m *customTbVoucherOrderModel) UpdateStatusAndUseTime(ctx context.Context, orderId uint64, status int) error {
	query := fmt.Sprintf("update %s set status = ?, use_time = NOW() where id = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, status, orderId)
	return err
}
