package consts

const (
	PayTypeWxPay     = "wxpay"  // 微信支付
	PayTypeAliPay    = "alipay" // 支付宝
	PayTypeSaobeiPay = "saobei" // 扫呗支付

	TradeTypeWxMP       = "mp"   // 微信交易类型 公众号
	TradeTypeWxMini     = "mini" // 微信交易类型 小程序
	TradeTypeWxApp      = "app"  // 微信交易类型 APP
	TradeTypeWxScan     = "scan" // 微信交易类型 二维码扫码
	TradeTypeWxPos      = "pos"  // 微信交易类型 二维码收款
	TradeTypeWxH5       = "h5"   // 微信交易类型  H5
	TradeTypeAliWeb     = "web"  // 支付宝 电脑网页
	TradeTypeAliApp     = "app"  // 支付宝 APP
	TradeTypeAliScan    = "scan" // 支付宝 二维码扫码
	TradeTypeAliWap     = "wap"  // 支付宝 手机网页
	TradeTypeAliPos     = "pos"  // 支付宝 二维码收款
	TradeTypeSaobeiMini = "mini" // 扫呗 小程序支付

	OrderGroupDefault   = "order"      // 普通订单
	OrderGroupMallOrder = "mall_order" // 商城订单

	PayOrderStatusUnpaid  = 1  // 未支付
	PayOrderStatusSuccess = 10 // 支付成功
	PayOrderStatusRefund  = 20 // 已退款
	PayOrderStatusClose   = 30 // 支付关闭

	PayRefundStatusWaiting = 1  // 退款状态 未退款
	PayRefundStatusSuccess = 10 // 退款状态 退款成功
	PayRefundStatusFail    = 20 // 退款状态 退款失败
)
