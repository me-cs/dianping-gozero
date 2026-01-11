package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbBlogCommentsModel = (*customTbBlogCommentsModel)(nil)

type (
	// TbBlogCommentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbBlogCommentsModel.
	TbBlogCommentsModel interface {
		tbBlogCommentsModel
	}

	customTbBlogCommentsModel struct {
		*defaultTbBlogCommentsModel
	}
)

// NewTbBlogCommentsModel returns a model for the database table.
func NewTbBlogCommentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbBlogCommentsModel {
	return &customTbBlogCommentsModel{
		defaultTbBlogCommentsModel: newTbBlogCommentsModel(conn, c, opts...),
	}
}
