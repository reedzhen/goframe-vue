// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PayOrder is the golang structure for table pay_order.
type PayOrder struct {
	Id              int64       `json:"id"              orm:"id"                ` //
	AppId           int64       `json:"appId"           orm:"app_id"            ` // 应用ID
	ChannelId       int64       `json:"channelId"       orm:"channel_id"        ` // 渠道ID
	ChannelCode     string      `json:"channelCode"     orm:"channel_code"      ` // 渠道编码 例如:alipay_pc/wx_lite
	MerchantOrderId int64       `json:"merchantOrderId" orm:"merchant_order_id" ` // 商户订单ID 例如：trade_order.id
	OrderGroup      string      `json:"orderGroup"      orm:"order_group"       ` // 组别 mall商城
	Subject         string      `json:"subject"         orm:"subject"           ` // 商品标题
	Body            string      `json:"body"            orm:"body"              ` // 商品描述
	NotifyUrl       string      `json:"notifyUrl"       orm:"notify_url"        ` // 支付通知回调地址
	PayType         string      `json:"payType"         orm:"pay_type"          ` // 支付类型 wxpay微信支付/alipay支付宝/saobei扫呗
	TradeType       string      `json:"tradeType"       orm:"trade_type"        ` // 交易类型 saobei(mini小程序) wxpay(mp公众号/mini小程序/appAPP/scan二维码扫码) alipay(web网页/appAPP/scan二维码扫码)
	Price           int         `json:"price"           orm:"price"             ` // 支付金额 单位:分
	ChannelFeeRate  float64     `json:"channelFeeRate"  orm:"channel_fee_rate"  ` // 渠道手续费率 单位:百分比
	ChannelFeePrice int         `json:"channelFeePrice" orm:"channel_fee_price" ` // 渠道手续费金额 单位:分
	Status          int         `json:"status"          orm:"status"            ` // 支付状态 1未支付/10支付成功/20已退款/30支付关闭
	UserIp          string      `json:"userIp"          orm:"user_ip"           ` // 用户IP
	ExpireTime      *gtime.Time `json:"expireTime"      orm:"expire_time"       ` // 订单失效时间
	SuccessTime     *gtime.Time `json:"successTime"     orm:"success_time"      ` // 订单支付成功时间
	ExtensionId     int64       `json:"extensionId"     orm:"extension_id"      ` // 订单拓展单编号 支付成功后写入
	No              string      `json:"no"              orm:"no"                ` // 支付订单号 extension.no 支付成功后写入
	RefundPrice     int         `json:"refundPrice"     orm:"refund_price"      ` // 退款总金额 单位：分
	ChannelUserId   string      `json:"channelUserId"   orm:"channel_user_id"   ` // 渠道用户编号
	ChannelOrderNo  string      `json:"channelOrderNo"  orm:"channel_order_no"  ` // 渠道订单号
	CreatedBy       int64       `json:"createdBy"       orm:"created_by"        ` // 添加人
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        ` // 创建时间
	UpdatedBy       int64       `json:"updatedBy"       orm:"updated_by"        ` // 更新人
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        ` // 更新时间
	TenantId        int64       `json:"tenantId"        orm:"tenant_id"         ` // 租户ID
}
