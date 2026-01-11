package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/shop/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/shop/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/shop/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type ListShopTypesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListShopTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListShopTypesLogic {
	return &ListShopTypesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListShopTypes 查询所有商铺类型
func (l *ListShopTypesLogic) ListShopTypes(in *pb.ListShopTypesReq) (*pb.ListShopTypesResp, error) {
	shopTypes, err := l.svcCtx.Dao.FindAllShopTypes(l.ctx)
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.ListShopTypesResp{
				ShopTypes: []*pb.ShopTypeInfo{},
			}, nil
		}
		l.Logger.Errorf("find all shop types failed: %v", err)
		return nil, err
	}

	// 转换为 pb.ShopTypeInfo
	shopTypeInfos := make([]*pb.ShopTypeInfo, 0, len(shopTypes))
	for _, shopType := range shopTypes {
		shopTypeInfo := &pb.ShopTypeInfo{
			Id:   int64(shopType.Id),
			Name: shopType.Name.String,
			Icon: shopType.Icon.String,
			Sort: int32(shopType.Sort.Int64),
		}
		shopTypeInfos = append(shopTypeInfos, shopTypeInfo)
	}

	return &pb.ListShopTypesResp{
		ShopTypes: shopTypeInfos,
	}, nil
}
