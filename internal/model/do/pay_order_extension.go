// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayOrderExtension is the golang structure of table pay_order_extension for DAO operations like Where/Data.
type PayOrderExtension struct {
	g.Meta            `orm:"table:pay_order_extension, do:true"`
	Id                any         //
	No                any         // 支付订单号
	PayOrderId        any         // 支付订单编号
	ChannelId         any         // 渠道编号
	ChannelCode       any         // 渠道编码
	UserIp            any         // 用户 IP
	Status            any         // 支付状态 1未支付/10支付成功/20已退款/30支付关闭
	ChannelExtras     any         // 支付渠道的额外参数
	ChannelErrorCode  any         // 渠道调用报错时，错误码
	ChannelErrorMsg   any         // 渠道调用报错时，错误信息
	ChannelNotifyData any         // 支付渠道异步通知的内容
	CreatedBy         any         // 添加人
	CreatedAt         *gtime.Time // 创建时间
	UpdatedBy         any         // 更新人
	UpdatedAt         *gtime.Time // 更新时间
	TenantId          any         // 租户ID
}
