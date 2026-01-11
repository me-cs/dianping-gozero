package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbFollowModel = (*customTbFollowModel)(nil)

type (
	// TbFollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbFollowModel.
	TbFollowModel interface {
		tbFollowModel
		FindOneByUserIdAndFollowUserId(ctx context.Context, userId, followUserId uint64) (*TbFollow, error)
		DeleteByUserIdAndFollowUserId(ctx context.Context, userId, followUserId uint64) error
		Count(ctx context.Context, userId, followUserId uint64) (int64, error)
	}

	customTbFollowModel struct {
		*defaultTbFollowModel
	}
)

// NewTbFollowModel returns a model for the database table.
func NewTbFollowModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TbFollowModel {
	return &customTbFollowModel{
		defaultTbFollowModel: newTbFollowModel(conn, c, opts...),
	}
}

// FindOneByUserIdAndFollowUserId 根据用户ID和被关注用户ID查询
func (m *customTbFollowModel) FindOneByUserIdAndFollowUserId(ctx context.Context, userId, followUserId uint64) (*TbFollow, error) {
	var resp TbFollow
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `follow_user_id` = ? limit 1", tbFollowRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, followUserId)
	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// DeleteByUserIdAndFollowUserId 根据用户ID和被关注用户ID删除
func (m *customTbFollowModel) DeleteByUserIdAndFollowUserId(ctx context.Context, userId, followUserId uint64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ? and `follow_user_id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, userId, followUserId)
	return err
}

// Count 统计关注关系数量
func (m *customTbFollowModel) Count(ctx context.Context, userId, followUserId uint64) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where `user_id` = ? and `follow_user_id` = ?", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, userId, followUserId)
	return count, err
}
