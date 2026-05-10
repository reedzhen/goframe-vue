package dto

// PayRefundInput 订单退款
type PayRefundInput struct {
	MerchantOrderId int64  // 商户订单ID
	RefundPrice     int    // 退款金额
	Reason          string // 申请退款原因
	Remark          string //退款备注
	CreatedBy       int64  // 创建人
}

// PayRefundExistInput 查询当前支付订单是否有退款中的订单
type PayRefundExistInput struct {
	AppId      int64 // 应用ID
	PayOrderId int64 // 支付订单ID
	Status     int   // 退款状态 0未退款/10退款成功/20退款失败
}
