package logic

import (
	"context"

	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/voucher/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type ListVouchersByShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVouchersByShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVouchersByShopLogic {
	return &ListVouchersByShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListVouchersByShop 查询商铺的优惠券列表
func (l *ListVouchersByShopLogic) ListVouchersByShop(in *pb.ListVouchersByShopReq) (*pb.ListVouchersByShopResp, error) {
	vouchers, err := l.svcCtx.Dao.FindVouchersByShopId(l.ctx, uint64(in.ShopId))
	if err != nil {
		if err == sqlc.ErrNotFound || err == model.ErrNotFound {
			return &pb.ListVouchersByShopResp{
				Vouchers: []*pb.VoucherInfo{},
			}, nil
		}
		l.Logger.Errorf("find vouchers by shop failed: %v", err)
		return nil, err
	}

	// 转换为 pb.VoucherInfo
	voucherInfos := make([]*pb.VoucherInfo, 0, len(vouchers))
	for _, voucher := range vouchers {
		voucherInfo := &pb.VoucherInfo{
			Id:          int64(voucher.Id),
			ShopId:      int64(voucher.ShopId.Int64),
			Title:       voucher.Title,
			SubTitle:    voucher.SubTitle.String,
			Rules:       voucher.Rules.String,
			PayValue:    int64(voucher.PayValue),
			ActualValue: int64(voucher.ActualValue),
			Type:        int32(voucher.Type),
			Status:      int32(voucher.Status),
			CreateTime:  voucher.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  voucher.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		voucherInfos = append(voucherInfos, voucherInfo)
	}

	return &pb.ListVouchersByShopResp{
		Vouchers: voucherInfos,
	}, nil
}
