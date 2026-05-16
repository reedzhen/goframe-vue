// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PayOrderExtension is the golang structure for table pay_order_extension.
type PayOrderExtension struct {
	Id                int64       `json:"id"                orm:"id"                  ` //
	No                string      `json:"no"                orm:"no"                  ` // 支付订单号
	PayOrderId        int64       `json:"payOrderId"        orm:"pay_order_id"        ` // 支付订单编号
	ChannelId         int64       `json:"channelId"         orm:"channel_id"          ` // 渠道编号
	ChannelCode       string      `json:"channelCode"       orm:"channel_code"        ` // 渠道编码
	UserIp            string      `json:"userIp"            orm:"user_ip"             ` // 用户 IP
	Status            int         `json:"status"            orm:"status"              ` // 支付状态 1未支付/10支付成功/20已退款/30支付关闭
	ChannelExtras     string      `json:"channelExtras"     orm:"channel_extras"      ` // 支付渠道的额外参数
	ChannelErrorCode  string      `json:"channelErrorCode"  orm:"channel_error_code"  ` // 渠道调用报错时，错误码
	ChannelErrorMsg   string      `json:"channelErrorMsg"   orm:"channel_error_msg"   ` // 渠道调用报错时，错误信息
	ChannelNotifyData string      `json:"channelNotifyData" orm:"channel_notify_data" ` // 支付渠道异步通知的内容
	CreatedBy         int64       `json:"createdBy"         orm:"created_by"          ` // 添加人
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          ` // 创建时间
	UpdatedBy         int64       `json:"updatedBy"         orm:"updated_by"          ` // 更新人
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          ` // 更新时间
	TenantId          int64       `json:"tenantId"          orm:"tenant_id"           ` // 租户ID
}
