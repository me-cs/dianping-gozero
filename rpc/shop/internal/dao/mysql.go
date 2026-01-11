package dao

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/shop/internal/config"
	"github.com/me-cs/dianping-gozero/rpc/shop/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type MySQL struct {
	shopModel     model.TbShopModel
	shopTypeModel model.TbShopTypeModel
}

func NewMySQL(c config.Config) *MySQL {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &MySQL{
		shopModel:     model.NewTbShopModel(conn, c.Cache),
		shopTypeModel: model.NewTbShopTypeModel(conn, c.Cache),
	}
}

// FindShopById 根据ID查询商铺
func (d *Dao) FindShopById(ctx context.Context, id uint64) (*model.TbShop, error) {
	return d.mysql.shopModel.FindOne(ctx, id)
}

// UpdateShop 更新商铺信息
func (d *Dao) UpdateShop(ctx context.Context, shop *model.TbShop) error {
	return d.mysql.shopModel.Update(ctx, shop)
}

// FindShopsByType 根据类型查询商铺列表
func (d *Dao) FindShopsByType(ctx context.Context, typeId uint64, page, pageSize int) ([]*model.TbShop, error) {
	return d.mysql.shopModel.FindByTypeId(ctx, typeId, page, pageSize)
}

// SearchShopsByName 根据名称搜索商铺
func (d *Dao) SearchShopsByName(ctx context.Context, name string, page, pageSize int) ([]*model.TbShop, error) {
	return d.mysql.shopModel.FindByName(ctx, name, page, pageSize)
}

// FindShopTypeById 根据ID查询商铺类型
func (d *Dao) FindShopTypeById(ctx context.Context, id uint64) (*model.TbShopType, error) {
	return d.mysql.shopTypeModel.FindOne(ctx, id)
}

// FindAllShopTypes 查询所有商铺类型
func (d *Dao) FindAllShopTypes(ctx context.Context) ([]*model.TbShopType, error) {
	return d.mysql.shopTypeModel.FindAll(ctx)
}
