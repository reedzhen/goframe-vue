package common

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SaobeiPayNotifyReq struct {
	g.Meta      `path:"/pay/notify/saobei/:extensionNo" method:"post" tags:"通用接口" summary:"扫呗支付回调" noAuth:"true"`
	ExtensionNo string `json:"extensionNo" dc:"支付订单号"`
}
type SaobeiPayNotifyRes struct{}

type WxPayNotifyReq struct {
	g.Meta      `path:"/pay/notify/wxpay/:extensionNo" method:"post" tags:"通用接口" summary:"微信支付回调" noAuth:"true"`
	ExtensionNo string `json:"extensionNo" dc:"支付订单号"`
}
type WxPayNotifyRes struct{}
