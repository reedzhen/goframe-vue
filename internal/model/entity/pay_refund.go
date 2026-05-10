// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PayRefund is the golang structure for table pay_refund.
type PayRefund struct {
	Id                uint64      `json:"id"                orm:"id"                  ` //
	No                string      `json:"no"                orm:"no"                  ` // 退款单号
	AppId             int64       `json:"appId"             orm:"app_id"              ` // 应用ID
	ChannelId         int64       `json:"channelId"         orm:"channel_id"          ` // 渠道ID
	ChannelCode       string      `json:"channelCode"       orm:"channel_code"        ` // 渠道编码 例如:alipay_pc/wx_lite
	PayOrderId        int64       `json:"payOrderId"        orm:"pay_order_id"        ` // 支付订单ID
	PayOrderNo        string      `json:"payOrderNo"        orm:"pay_order_no"        ` // 支付订单编号
	MerchantOrderId   int64       `json:"merchantOrderId"   orm:"merchant_order_id"   ` // 商户订单ID 例：商城订单ID
	MerchantRefundId  int64       `json:"merchantRefundId"  orm:"merchant_refund_id"  ` // 商户退款订单ID
	NotifyUrl         string      `json:"notifyUrl"         orm:"notify_url"          ` // 异步通知商户地址
	Status            int         `json:"status"            orm:"status"              ` // 退款状态 0未退款/10退款成功/20退款失败
	PayPrice          int         `json:"payPrice"          orm:"pay_price"           ` // 支付金额 单位分
	RefundPrice       int         `json:"refundPrice"       orm:"refund_price"        ` // 退款金额 单位分
	Reason            string      `json:"reason"            orm:"reason"              ` // 申请退款原因
	UserIp            string      `json:"userIp"            orm:"user_ip"             ` // 用户IP
	ChannelOrderNo    string      `json:"channelOrderNo"    orm:"channel_order_no"    ` // 渠道订单号，pay_order 中的 channel_order_no 对应
	ChannelRefundNo   string      `json:"channelRefundNo"   orm:"channel_refund_no"   ` // 渠道退款单号，渠道返回
	SuccessTime       *gtime.Time `json:"successTime"       orm:"success_time"        ` // 退款成功时间
	ChannelErrorCode  string      `json:"channelErrorCode"  orm:"channel_error_code"  ` // 渠道调用报错时，错误码
	ChannelErrorMsg   string      `json:"channelErrorMsg"   orm:"channel_error_msg"   ` // 渠道调用报错时，错误信息
	ChannelNotifyData string      `json:"channelNotifyData" orm:"channel_notify_data" ` // 支付渠道异步通知的内容
	Remark            string      `json:"remark"            orm:"remark"              ` // 退款备注
	CreatedBy         int64       `json:"createdBy"         orm:"created_by"          ` // 添加人
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          ` // 创建时间
	UpdatedBy         int64       `json:"updatedBy"         orm:"updated_by"          ` // 更新人
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          ` // 更新时间
	TenantId          int64       `json:"tenantId"          orm:"tenant_id"           ` // 租户ID
}
