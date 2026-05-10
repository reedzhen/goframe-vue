// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
)

type (
	IPayOrder interface {
		// CreateOrder 创建支付订单(创建商城订单时会调用此方法)
		CreateOrder(ctx context.Context, in dto.PayOrderCreateInput) (out int64, err error)
		// SubmitOrder 提交支付 此时，会发起支付渠道的调用
		SubmitOrder(ctx context.Context, in dto.PayOrderSubmitInput) (out *dto.PayOrderSubmitOutput, err error)
		// GenNotifyURL 生成支付通知地址
		GenNotifyURL(ctx context.Context, in dto.PayOrderGenNotifyURLInput) (notifyURL string, err error)
		// GetInfoByMerchantOrderId 获取支付订单详情 todo 实际情况下 app_id+merchant_order_id才能确定唯一一条记录，这里只有商城用，所以适用
		GetInfoByMerchantOrderId(ctx context.Context, merchantOrderId int64) (out *entity.PayOrder, err error)
		// GetInfo 获取支付订单详情
		GetInfo(ctx context.Context, payOrderId int64) (out *entity.PayOrder, err error)
		// GetExtensionInfo 获取支付订单拓展单详情
		GetExtensionInfo(ctx context.Context, no string) (out *entity.PayOrderExtension, err error)
		// ExpirePayOrder 将已过期的订单，状态修改为已关闭(cron)
		ExpirePayOrder(ctx context.Context)
		// Notify 异步通知
		Notify(ctx context.Context, in dto.PayOrderNotifyInput) (res *dto.PayOrderNotifyOutput, err error)
	}
	IPayRefund interface {
		// Refund 订单退款
		Refund(ctx context.Context, in dto.PayRefundInput) (res *entity.PayRefund, err error)
	}
)

var (
	localPayOrder  IPayOrder
	localPayRefund IPayRefund
)

func PayOrder() IPayOrder {
	if localPayOrder == nil {
		panic("implement not found for interface IPayOrder, forgot register?")
	}
	return localPayOrder
}

func RegisterPayOrder(i IPayOrder) {
	localPayOrder = i
}

func PayRefund() IPayRefund {
	if localPayRefund == nil {
		panic("implement not found for interface IPayRefund, forgot register?")
	}
	return localPayRefund
}

func RegisterPayRefund(i IPayRefund) {
	localPayRefund = i
}
