package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbShopTypeModel = (*customTbShopTypeModel)(nil)

type (
	// TbShopTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbShopTypeModel.
	TbShopTypeModel interface {
		tbShopTypeModel
		FindAll(ctx context.Context) ([]*TbShopType, error)
	}

	customTbShopTypeModel struct {
		*defaultTbShopTypeModel
	}
)

// NewTbShopTypeModel returns a model for the database table.
func NewTbShopTypeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbShopTypeModel {
	return &customTbShopTypeModel{
		defaultTbShopTypeModel: newTbShopTypeModel(conn, c, opts...),
	}
}

// FindAll 查询所有商铺类型
func (m *customTbShopTypeModel) FindAll(ctx context.Context) ([]*TbShopType, error) {
	var shopTypes []*TbShopType
	query := fmt.Sprintf("select %s from %s order by sort asc", tbShopTypeRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &shopTypes, query)
	switch err {
	case nil:
		return shopTypes, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
