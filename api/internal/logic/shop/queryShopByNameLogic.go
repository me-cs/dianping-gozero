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

type QueryShopByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryShopByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryShopByNameLogic {
	return &QueryShopByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryShopByNameLogic) QueryShopByName(req *types.QueryShopByNameReq) (resp *types.QueryShopByNameRsp, err error) {
	// 调用 rpc 层搜索商铺
	searchShopsReq := &shop.SearchShopsReq{
		Name:    req.Name,
		Current: int32(req.Current),
	}

	rsp, err := l.svcCtx.ShopRpc.SearchShops(l.ctx, searchShopsReq)
	if err != nil {
		l.Errorf("调用 ShopRpc.SearchShops 失败: %v", err)
		return &types.QueryShopByNameRsp{
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

	return &types.QueryShopByNameRsp{
		Success: true,
		Data:    shops,
	}, nil
}
