// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayOrder is the golang structure of table pay_order for DAO operations like Where/Data.
type PayOrder struct {
	g.Meta          `orm:"table:pay_order, do:true"`
	Id              any         //
	AppId           any         // 应用ID
	ChannelId       any         // 渠道ID
	ChannelCode     any         // 渠道编码 例如:alipay_pc/wx_lite
	MerchantOrderId any         // 商户订单ID 例如：trade_order.id
	OrderGroup      any         // 组别 mall商城
	Subject         any         // 商品标题
	Body            any         // 商品描述
	NotifyUrl       any         // 支付通知回调地址
	PayType         any         // 支付类型 wxpay微信支付/alipay支付宝/saobei扫呗
	TradeType       any         // 交易类型 saobei(mini小程序) wxpay(mp公众号/mini小程序/appAPP/scan二维码扫码) alipay(web网页/appAPP/scan二维码扫码)
	Price           any         // 支付金额 单位:分
	ChannelFeeRate  any         // 渠道手续费率 单位:百分比
	ChannelFeePrice any         // 渠道手续费金额 单位:分
	Status          any         // 支付状态 1未支付/10支付成功/20已退款/30支付关闭
	UserIp          any         // 用户IP
	ExpireTime      *gtime.Time // 订单失效时间
	SuccessTime     *gtime.Time // 订单支付成功时间
	ExtensionId     any         // 订单拓展单编号 支付成功后写入
	No              any         // 支付订单号 extension.no 支付成功后写入
	RefundPrice     any         // 退款总金额 单位：分
	ChannelUserId   any         // 渠道用户编号
	ChannelOrderNo  any         // 渠道订单号
	CreatedBy       any         // 添加人
	CreatedAt       *gtime.Time // 创建时间
	UpdatedBy       any         // 更新人
	UpdatedAt       *gtime.Time // 更新时间
	TenantId        any         // 租户ID
}
