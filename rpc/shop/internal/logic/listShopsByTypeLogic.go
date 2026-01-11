package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/shop/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/shop/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/shop/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type ListShopsByTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListShopsByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListShopsByTypeLogic {
	return &ListShopsByTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListShopsByType 根据类型查询商铺列表（支持地理位置搜索）
func (l *ListShopsByTypeLogic) ListShopsByType(in *pb.ListShopsByTypeReq) (*pb.ListShopsByTypeResp, error) {
	pageSize := 10
	if in.Current <= 0 {
		in.Current = 1
	}

	// 如果提供了坐标，进行地理位置搜索
	if in.X != 0 && in.Y != 0 {
		// TODO: 实现地理位置搜索
		// shopIds, distances, err := l.svcCtx.Dao.GeoSearch(l.ctx, uint64(in.TypeId), in.X, in.Y, 5000)
		// 目前先降级到普通查询
	}

	// 普通数据库查询
	shops, err := l.svcCtx.Dao.FindShopsByType(l.ctx, uint64(in.TypeId), int(in.Current), pageSize)
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.ListShopsByTypeResp{
				Shops: []*pb.ShopInfo{},
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("find shops by type failed: %v", err)
		return nil, err
	}

	// 转换为 pb.ShopInfo
	shopInfos := make([]*pb.ShopInfo, 0, len(shops))
	for _, shop := range shops {
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

		shopInfos = append(shopInfos, shopInfo)
	}

	return &pb.ListShopsByTypeResp{
		Shops: shopInfos,
		Total: int64(len(shopInfos)),
	}, nil
}
