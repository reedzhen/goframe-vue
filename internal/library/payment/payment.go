package payment

import (
	"context"
	"fmt"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/payment/saobei"
	"goframe-vben/internal/library/payment/wxpay"
	"goframe-vben/internal/model/dto"
)

// PayClient 支付客户端
type PayClient interface {
	// CreateOrder 创建订单
	CreateOrder(ctx context.Context, in dto.CreateOrderInput) (res *dto.CreateOrderOutput, err error)
	// Notify 异步通知
	Notify(ctx context.Context, in dto.NotifyInput) (res *dto.NotifyOutput, err error)
	// Refund 订单退款
	Refund(ctx context.Context, in dto.RefundInput) (res *dto.RefundOutput, err error)
}

func New(name ...string) PayClient {
	var (
		payType = consts.PayTypeWxPay
		client  PayClient
	)

	if len(name) > 0 && name[0] != "" {
		payType = name[0]
	}

	switch payType {
	case consts.PayTypeWxPay:
		client = wxpay.New(config)
	case consts.PayTypeSaobeiPay:
		client = saobei.New(config)
	default:
		panic(fmt.Sprintf("暂不支持的支付方式:%v", payType))
	}
	return client
}
