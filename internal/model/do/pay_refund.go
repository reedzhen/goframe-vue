// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayRefund is the golang structure of table pay_refund for DAO operations like Where/Data.
type PayRefund struct {
	g.Meta            `orm:"table:pay_refund, do:true"`
	Id                any         //
	No                any         // 退款单号
	AppId             any         // 应用ID
	ChannelId         any         // 渠道ID
	ChannelCode       any         // 渠道编码 例如:alipay_pc/wx_lite
	PayOrderId        any         // 支付订单ID
	PayOrderNo        any         // 支付订单编号
	MerchantOrderId   any         // 商户订单ID 例：商城订单ID
	MerchantRefundId  any         // 商户退款订单ID
	NotifyUrl         any         // 异步通知商户地址
	Status            any         // 退款状态 0未退款/10退款成功/20退款失败
	PayPrice          any         // 支付金额 单位分
	RefundPrice       any         // 退款金额 单位分
	Reason            any         // 申请退款原因
	UserIp            any         // 用户IP
	ChannelOrderNo    any         // 渠道订单号，pay_order 中的 channel_order_no 对应
	ChannelRefundNo   any         // 渠道退款单号，渠道返回
	SuccessTime       *gtime.Time // 退款成功时间
	ChannelErrorCode  any         // 渠道调用报错时，错误码
	ChannelErrorMsg   any         // 渠道调用报错时，错误信息
	ChannelNotifyData any         // 支付渠道异步通知的内容
	Remark            any         // 退款备注
	CreatedBy         any         // 添加人
	CreatedAt         *gtime.Time // 创建时间
	UpdatedBy         any         // 更新人
	UpdatedAt         *gtime.Time // 更新时间
	TenantId          any         // 租户ID
}
