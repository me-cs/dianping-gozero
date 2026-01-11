package logic

import (
	"context"
	"encoding/json"

	"github.com/me-cs/dianping-gozero/rpc/shop/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/shop/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/shop/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type GetShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShopLogic {
	return &GetShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetShop 根据ID查询商铺（带缓存穿透处理）
func (l *GetShopLogic) GetShop(in *pb.GetShopReq) (*pb.GetShopResp, error) {
	shopId := uint64(in.ShopId)

	// 1. 从Redis查询商铺缓存
	cacheValue, err := l.svcCtx.Dao.GetShopCache(l.ctx, shopId)
	if err == nil && cacheValue != "" {
		// 缓存命中
		if cacheValue == "null" {
			// 空值缓存，防止缓存穿透
			return &pb.GetShopResp{}, nil
		}

		// 解析缓存数据
		var shop model.TbShop
		if err := json.Unmarshal([]byte(cacheValue), &shop); err == nil {
			return l.shopToResp(&shop), nil
		}
	}

	// 2. 缓存未命中，从数据库查询
	shop, err := l.svcCtx.Dao.FindShopById(l.ctx, shopId)
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			// 商铺不存在，设置空值缓存防止缓存穿透
			_ = l.svcCtx.Dao.SetShopCacheWithNullValue(l.ctx, shopId)
			return &pb.GetShopResp{}, nil
		}
		l.Logger.Errorf("find shop failed: %v", err)
		return nil, err
	}

	// 3. 写入缓存
	shopJson, _ := json.Marshal(shop)
	_ = l.svcCtx.Dao.SetShopCache(l.ctx, shopId, string(shopJson))

	return l.shopToResp(shop), nil
}

// shopToResp 转换 shop model 到 pb response
func (l *GetShopLogic) shopToResp(shop *model.TbShop) *pb.GetShopResp {
	shopInfo := &pb.ShopInfo{
		Id:         int64(shop.Id),
		Name:       shop.Name,
		TypeId:     int64(shop.TypeId),
		Images:     shop.Images,
		Address:    shop.Address,
		X:          shop.X,
		Y:          shop.Y,
		Sold:       int32(shop.Sold),
		Comments:   int32(shop.Comments),
		Score:      int32(shop.Score),
		CreateTime: shop.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: shop.UpdateTime.Format("2006-01-02 15:04:05"),
	}

	if shop.Area.Valid {
		shopInfo.Area = shop.Area.String
	}
	if shop.AvgPrice.Valid {
		shopInfo.AvgPrice = shop.AvgPrice.Int64
	}
	if shop.OpenHours.Valid {
		shopInfo.OpenHours = shop.OpenHours.String
	}

	return &pb.GetShopResp{Shop: shopInfo}
}
