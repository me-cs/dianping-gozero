package dao

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/blog/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/blog/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type MySQL struct {
	blogModel         model.TbBlogModel
	blogCommentsModel model.TbBlogCommentsModel
}

func NewMySQL(c config.Config) *MySQL {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &MySQL{
		blogModel:         model.NewTbBlogModel(conn, c.Cache),
		blogCommentsModel: model.NewTbBlogCommentsModel(conn, c.Cache),
	}
}

// InsertBlog 插入博客
func (d *Dao) InsertBlog(ctx context.Context, blog *model.TbBlog) error {
	_, err := d.mysql.blogModel.Insert(ctx, blog)
	return err
}

// FindBlogById 根据ID查询博客
func (d *Dao) FindBlogById(ctx context.Context, id uint64) (*model.TbBlog, error) {
	return d.mysql.blogModel.FindOne(ctx, id)
}

// UpdateBlog 更新博客
func (d *Dao) UpdateBlog(ctx context.Context, blog *model.TbBlog) error {
	return d.mysql.blogModel.Update(ctx, blog)
}

// QueryHotBlogs 查询热门博客（按点赞数排序）
func (d *Dao) QueryHotBlogs(ctx context.Context, page, pageSize int) ([]*model.TbBlog, error) {
	return d.mysql.blogModel.FindHotBlogs(ctx, page, pageSize)
}

// QueryBlogsByUserId 根据用户ID查询博客
func (d *Dao) QueryBlogsByUserId(ctx context.Context, userId uint64, page, pageSize int) ([]*model.TbBlog, error) {
	return d.mysql.blogModel.FindByUserId(ctx, userId, page, pageSize)
}
