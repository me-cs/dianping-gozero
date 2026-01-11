package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbBlogModel = (*customTbBlogModel)(nil)

type (
	// TbBlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbBlogModel.
	TbBlogModel interface {
		tbBlogModel
		FindHotBlogs(ctx context.Context, page, pageSize int) ([]*TbBlog, error)
		FindByUserId(ctx context.Context, userId uint64, page, pageSize int) ([]*TbBlog, error)
	}

	customTbBlogModel struct {
		*defaultTbBlogModel
	}
)

// NewTbBlogModel returns a model for the database table.
func NewTbBlogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbBlogModel {
	return &customTbBlogModel{
		defaultTbBlogModel: newTbBlogModel(conn, c, opts...),
	}
}

// FindHotBlogs 查询热门博客（按点赞数倒序）
func (m *customTbBlogModel) FindHotBlogs(ctx context.Context, page, pageSize int) ([]*TbBlog, error) {
	var blogs []*TbBlog
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s order by liked desc limit ?,?", tbBlogRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &blogs, query, offset, pageSize)
	switch err {
	case nil:
		return blogs, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindByUserId 根据用户ID查询博客
func (m *customTbBlogModel) FindByUserId(ctx context.Context, userId uint64, page, pageSize int) ([]*TbBlog, error) {
	var blogs []*TbBlog
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s where user_id = ? order by create_time desc limit ?,?", tbBlogRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &blogs, query, userId, offset, pageSize)
	switch err {
	case nil:
		return blogs, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
