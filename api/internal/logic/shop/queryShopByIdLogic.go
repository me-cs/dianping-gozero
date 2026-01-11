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

type QueryShopByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryShopByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryShopByIdLogic {
	return &QueryShopByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryShopByIdLogic) QueryShopById(req *types.QueryShopByIdReq) (resp *types.QueryShopByIdRsp, err error) {
	// 调用 rpc 层查询商铺详情
	getShopReq := &shop.GetShopReq{
		ShopId: req.Id,
	}

	rsp, err := l.svcCtx.ShopRpc.GetShop(l.ctx, getShopReq)
	if err != nil {
		l.Errorf("调用 ShopRpc.GetShop 失败: %v", err)
		return &types.QueryShopByIdRsp{
			Success: false,
			ErrMsg:  "系统错误",
		}, err
	}

	// 转换 pb.ShopInfo 为 types.Shop
	shopInfo := rsp.GetShop()
	if shopInfo == nil {
		return &types.QueryShopByIdRsp{
			Success: false,
			ErrMsg:  "商铺不存在",
		}, nil
	}

	return &types.QueryShopByIdRsp{
		Success: true,
		Data: &types.Shop{
			Id:         shopInfo.Id,
			Name:       shopInfo.Name,
			TypeId:     shopInfo.TypeId,
			Images:     shopInfo.Images,
			Area:       shopInfo.Area,
			Address:    shopInfo.Address,
			X:          shopInfo.X,
			Y:          shopInfo.Y,
			AvgPrice:   shopInfo.AvgPrice,
			Sold:       int(shopInfo.Sold),
			Comments:   int(shopInfo.Comments),
			Score:      int(shopInfo.Score),
			OpenHours:  shopInfo.OpenHours,
			CreateTime: shopInfo.CreateTime,
			UpdateTime: shopInfo.UpdateTime,
			Distance:   shopInfo.Distance,
		},
	}, nil
}
