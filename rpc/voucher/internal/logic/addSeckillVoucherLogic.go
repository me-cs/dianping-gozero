package logic

import (
	"context"
	"time"

	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/model"
	"github.com/me-cs/dianping-gozero/rpc/voucher/internal/svc"
	"github.com/me-cs/dianping-gozero/rpc/voucher/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSeckillVoucherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSeckillVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillVoucherLogic {
	return &AddSeckillVoucherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSeckillVoucher 添加秒杀券（并初始化Redis库存）
func (l *AddSeckillVoucherLogic) AddSeckillVoucher(in *pb.AddSeckillVoucherReq) (*pb.AddSeckillVoucherResp, error) {
	// 1. 解析时间
	beginTime, err := time.Parse("2006-01-02 15:04:05", in.BeginTime)
	if err != nil {
		l.Logger.Errorf("parse begin time failed: %v", err)
		return &pb.AddSeckillVoucherResp{Success: false}, nil
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", in.EndTime)
	if err != nil {
		l.Logger.Errorf("parse end time failed: %v", err)
		return &pb.AddSeckillVoucherResp{Success: false}, nil
	}

	// 2. 创建秒杀券对象
	now := time.Now()
	seckillVoucher := &model.TbSeckillVoucher{
		VoucherId:  uint64(in.VoucherId),
		Stock:      int64(uint64(in.Stock)),
		BeginTime:  beginTime,
		EndTime:    endTime,
		CreateTime: now,
		UpdateTime: now,
	}

	// 3. 保存到数据库
	err = l.svcCtx.Dao.InsertSeckillVoucher(l.ctx, seckillVoucher)
	if err != nil {
		l.Logger.Errorf("insert seckill voucher failed: %v", err)
		return nil, err
	}

	// 4. 初始化Redis库存
	err = l.svcCtx.Dao.InitSeckillStock(l.ctx, uint64(in.VoucherId), int(in.Stock))
	if err != nil {
		l.Logger.Errorf("init seckill stock failed: %v", err)
		return nil, err
	}

	return &pb.AddSeckillVoucherResp{Success: true}, nil
}
