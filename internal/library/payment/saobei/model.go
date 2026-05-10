package saobei

import (
	"gitee.com/cleanpay/cleanpay/saobei"
)

// SaobeiNotifyReq 扫呗回调参数，应该移到gitee包里
type SaobeiNotifyReq struct {
	saobei.RspBase
	PayVer              string `json:"pay_ver"`               //接口内业务逻辑兼容版本号，可用值201、202
	PayType             string `json:"pay_type"`              //支付方式，010微信，020支付宝
	UserId              string `json:"user_id"`               //付款方用户id，“微信openid”、“支付宝账户”
	MerchantName        string `json:"merchant_name"`         //商户名称
	TerminalId          string `json:"terminal_id"`           //终端号
	DeviceNo            string `json:"device_no"`             //商户终端设备号(商户自定义，如门店编号),必须在平台已配置过
	TerminalTrace       string `json:"terminal_trace"`        //终端流水号，商户系统的订单号，系统原样返回
	TerminalTime        string `json:"terminal_time"`         //终端交易时间，yyyyMMddHHmmss，全局统一时间格式，系统原样返回
	PayTrace            string `json:"pay_trace"`             //当前支付终端流水号
	PayTime             string `json:"pay_time"`              //当前支付终端交易时间，yyyyMMddHHmmss，全局统一时间格式
	TotalFee            string `json:"total_fee"`             //金额，单位分
	EndTime             string `json:"end_time"`              //退款完成时间，yyyyMMddHHmmss，全局统一时间格式
	OutTradeNo          string `json:"out_trade_no"`          //平台唯一订单号
	ChannelTradeNo      string `json:"channel_trade_no"`      //通道订单号，微信订单号、支付宝订单号等
	ChannelOrderNo      string `json:"channel_order_no"`      //银行渠道订单号，微信支付时显示在支付成功页面的条码，可用作扫码查询和扫码退款时匹配
	Attach              string `json:"attach"`                //附加数据,原样返回
	ReceiptFee          string `json:"receipt_fee"`           //商家应结算金额,单位分
	BuyerPayFee         string `json:"buyer_pay_fee"`         //买家实付金额（分）pay_ver为202时返回
	PlatformDiscountFee string `json:"platform_discount_fee"` //平台优惠金额（分）pay_ver为202时返回
	MerchantDiscountFee string `json:"merchant_discount_fee"` //商家优惠金额（分）pay_ver为202时返回
	BankType            string `json:"bank_type"`             //银行类型，采用字符串类型的银行标识
	PromotionDetail     string `json:"promotion_detail"`      //官方营销详情,pay_ver=202时返回. 本交易支付时使用的所有优惠券信息 ，单品优惠功能字段，详情见
	OrderBody           string `json:"order_body"`            //订单标题描述
}
