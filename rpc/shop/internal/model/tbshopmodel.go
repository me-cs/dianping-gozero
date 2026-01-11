package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbShopModel = (*customTbShopModel)(nil)

type (
	// TbShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbShopModel.
	TbShopModel interface {
		tbShopModel
		FindByTypeId(ctx context.Context, typeId uint64, page, pageSize int) ([]*TbShop, error)
		FindByName(ctx context.Context, name string, page, pageSize int) ([]*TbShop, error)
	}

	customTbShopModel struct {
		*defaultTbShopModel
	}
)

// NewTbShopModel returns a model for the database table.
func NewTbShopModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbShopModel {
	return &customTbShopModel{
		defaultTbShopModel: newTbShopModel(conn, c, opts...),
	}
}

// FindByTypeId 根据类型ID查询商铺列表
func (m *customTbShopModel) FindByTypeId(ctx context.Context, typeId uint64, page, pageSize int) ([]*TbShop, error) {
	var shops []*TbShop
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s where type_id = ? order by create_time desc limit ?,?", tbShopRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &shops, query, typeId, offset, pageSize)
	switch err {
	case nil:
		return shops, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindByName 根据名称模糊查询商铺列表
func (m *customTbShopModel) FindByName(ctx context.Context, name string, page, pageSize int) ([]*TbShop, error) {
	var shops []*TbShop
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s where name like ? order by create_time desc limit ?,?", tbShopRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &shops, query, "%"+name+"%", offset, pageSize)
	switch err {
	case nil:
		return shops, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
