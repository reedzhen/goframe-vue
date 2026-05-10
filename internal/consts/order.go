package consts

const (
	PaymentMethodWechat  = "wechat"
	PaymentMethodAlipay  = "alipay"
	PaymentMethodOffline = "offline" // 线下付款

	OrderStatusDefault = "default" // 未支付
	OrderStatusPaid    = "paid"    // 已支付

	RefundStatusPending    = "pending"    // 未退款
	RefundStatusApplied    = "applied"    // 已申请退款
	RefundStatusProcessing = "processing" // 退款中
	RefundStatusSuccess    = "success"    // 退款成功
	RefundStatusFailed     = "failed"     // 退款失败
)
