// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package shopType

import (
	"context"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/me-cs/dianping-gozero/rpc/shop/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTypeListLogic {
	return &QueryTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryTypeListLogic) QueryTypeList(req *types.QueryTypeListReq) (resp *types.QueryTypeListRsp, err error) {
	// todo: add your logic here and delete this line
	listRsp, err := l.svcCtx.ShopRpc.ListShopTypes(l.ctx, &shop.ListShopTypesReq{})
	if err != nil {
		l.Errorf("调用 ShopRpc.ListShopTypes 失败: %v", err)
		return &types.QueryTypeListRsp{
			Success: false,
			ErrMsg:  "获取店铺类型失败",
		}, nil
	}

	resp = &types.QueryTypeListRsp{
		Success: true,
	}
	for _, v := range listRsp.GetShopTypes() {
		resp.Data = append(resp.Data, types.ShopType{
			Id:   v.GetId(),
			Name: v.GetName(),
			Icon: v.GetIcon(),
			Sort: int(v.GetSort()),
		})
	}
	return resp, nil
}
