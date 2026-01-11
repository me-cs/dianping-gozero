package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbSignModel = (*customTbSignModel)(nil)

type (
	// TbSignModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbSignModel.
	TbSignModel interface {
		tbSignModel
	}

	customTbSignModel struct {
		*defaultTbSignModel
	}
)

// NewTbSignModel returns a model for the database table.
func NewTbSignModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbSignModel {
	return &customTbSignModel{
		defaultTbSignModel: newTbSignModel(conn, c, opts...),
	}
}
