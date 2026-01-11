package dao

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/user/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/user/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type MySQL struct {
	userModel     model.TbUserModel
	userInfoModel model.TbUserInfoModel
	followModel   model.TbFollowModel
}

func NewMySQL(c config.Config) *MySQL {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &MySQL{
		userModel:     model.NewTbUserModel(conn, c.Cache),
		userInfoModel: model.NewTbUserInfoModel(conn, c.Cache),
		followModel:   model.NewTbFollowModel(conn, c.Cache),
	}
}

func (d *Dao) InsertUserInfo(c context.Context, user *model.TbUser) error {
	_, err := d.mysql.userModel.Insert(c, user)
	return err
}

func (d *Dao) FindOneByPhone(c context.Context, phone string) (*model.TbUser, error) {
	return d.mysql.userModel.FindOneByPhone(c, phone)
}

func (d *Dao) FindOneById(c context.Context, id uint64) (*model.TbUser, error) {
	return d.mysql.userModel.FindOne(c, id)
}

// FindUserInfoById 根据用户ID查询用户详细信息
func (d *Dao) FindUserInfoById(c context.Context, id uint64) (*model.TbUserInfo, error) {
	return d.mysql.userInfoModel.FindOne(c, id)
}

// InsertFollow 新增关注关系
func (d *Dao) InsertFollow(c context.Context, follow *model.TbFollow) error {
	_, err := d.mysql.followModel.Insert(c, follow)
	return err
}

// DeleteFollowByUserIdAndFollowUserId 根据用户ID和被关注用户ID删除关注关系
func (d *Dao) DeleteFollowByUserIdAndFollowUserId(c context.Context, userId, followUserId uint64) error {
	return d.mysql.followModel.DeleteByUserIdAndFollowUserId(c, userId, followUserId)
}

// CountFollow 统计关注关系数量
func (d *Dao) CountFollow(c context.Context, userId, followUserId uint64) (int64, error) {
	return d.mysql.followModel.Count(c, userId, followUserId)
}
