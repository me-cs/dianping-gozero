// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package shop

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/shop/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryShopByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryShopByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryShopByTypeLogic {
	return &QueryShopByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryShopByTypeLogic) QueryShopByType(req *types.QueryShopByTypeReq) (resp *types.QueryShopByTypeRsp, err error) {
	// 调用 rpc 层按类型查询商铺
	listShopsByTypeReq := &shop.ListShopsByTypeReq{
		TypeId:  req.TypeId,
		Current: int32(req.Current),
		X:       req.X,
		Y:       req.Y,
	}

	rsp, err := l.svcCtx.ShopRpc.ListShopsByType(l.ctx, listShopsByTypeReq)
	if err != nil {
		l.Errorf("调用 ShopRpc.ListShopsByType 失败: %v", err)
		return &types.QueryShopByTypeRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	// 转换 []*pb.ShopInfo 为 []types.Shop
	pbShops := rsp.GetShops()
	shops := make([]types.Shop, 0, len(pbShops))
	for _, pbShop := range pbShops {
		shops = append(shops, types.Shop{
			Id:         pbShop.Id,
			Name:       pbShop.Name,
			TypeId:     pbShop.TypeId,
			Images:     pbShop.Images,
			Area:       pbShop.Area,
			Address:    pbShop.Address,
			X:          pbShop.X,
			Y:          pbShop.Y,
			AvgPrice:   pbShop.AvgPrice,
			Sold:       int(pbShop.Sold),
			Comments:   int(pbShop.Comments),
			Score:      int(pbShop.Score),
			OpenHours:  pbShop.OpenHours,
			CreateTime: pbShop.CreateTime,
			UpdateTime: pbShop.UpdateTime,
			Distance:   pbShop.Distance,
		})
	}

	return &types.QueryShopByTypeRsp{
		Success: true,
		Data:    shops,
	}, nil
}
