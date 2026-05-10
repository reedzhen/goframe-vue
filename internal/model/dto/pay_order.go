package dto

import (
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/model/entity"
)

// NotifyCallFuncInput 异步通知回调，用于异步通知验签通过后回调到具体的业务中
type NotifyCallFuncInput struct {
	PayOrder *entity.PayOrder
}

// CreateOrderInput 统一创建订单入口
type CreateOrderInput struct {
	TradeType string // 交易类型 saobei(mini小程序) wxpay(mp公众号/mini小程序/appAPP/scan二维码扫码) alipay(web网页/appAPP/scan二维码扫码)
	UserIP    string // 用户 IP

	// 商户相关字段
	OutTradeNo string // 对应 PayOrderExtensionDO 的 no 字段
	Subject    string // 商品标题 max = 32
	Body       string // 商品描述信息 max = 128
	NotifyUrl  string // 支付结果的 notify 回调地址
	ReturnUrl  string // 支付结果的 return 回调地址

	// 订单相关字段
	Price      int         // 支付金额，单位：分
	ExpireTime *gtime.Time // 支付过期时间

	// 拓展参数
	Openid string // 微信支付 扫呗支付需要
}

type CreateOrderOutput struct {
	TradeType string `json:"tradeType" ` // 交易类型
	//PayURL     string                 `json:"payURL" `     // 支付地址
	//OutTradeNo string                 `json:"outTradeNo" ` // 商户订单号
	JsApi *wechat.JSAPIPayParams `json:"jsApi"  ` // jsapi支付参数
}

// NotifyInput 统一异步通知处理入口
type NotifyInput struct {
}

type NotifyOutput struct {
	OutTradeNo    string      `json:"outTradeNo"`    // 商户订单号
	TransactionId string      `json:"transactionId"` // 第三方订单号
	PayAt         *gtime.Time `json:"payAt"`         // 支付时间
	ActualAmount  int         `json:"actualAmount"`  // 实付金额 单位：分
}

// RefundInput 统一退款处理入口
type RefundInput struct {
	TransactionId string // 微信、扫呗的订单号
	RefundNo      string // 商户自定义退款单号
	TotalMoney    int    // 支付总额 微信退款用
	RefundMoney   int    // 退款金额
	Reason        string // 申请退款原因
	//Remark        string // 退款备注
}

type RefundOutput struct {
}

// PayOrderCreateInput 创建支付订单
type PayOrderCreateInput struct {
	Subject string // 订单标题
	//Detail     *gjson.Json // 支付商品详情
	MerchantOrderId int64  // 商户订单号
	OrderGroup      string // 组别 mall商城
	//Openid     string      // openid
	PayType   string // wxpay微信支付/alipay支付宝/saobei扫呗
	TradeType string // 交易类型  saobei(mini小程序) wxpay(mp公众号/mini小程序/appAPP/scan二维码扫码) alipay(web网页/appAPP/scan二维码扫码)
	PayAmount int    // 支付金额 单位分
	//ReturnUrl string // 买家付款成功跳转地址
	UserId int64  // 用户ID order.user_id
	UserIp string // 用户IP
	//PartnerId int64  // 联盟ID
	//PlaceId   int64  // 门店ID
	//CompanyId int64  // 公司ID
}

// PayOrderGenNotifyURLInput 生成支付通知地址
type PayOrderGenNotifyURLInput struct {
	PayType string // wxpay微信支付/alipay支付宝/saobei扫呗
	No      string // 支付订单号
}

// PayOrderSubmitInput 提交支付 此时，会发起支付渠道的调用
type PayOrderSubmitInput struct {
	PayOrderId int64  // 支付订单Id
	UserIp     string // 用户IP
	ReturnUrl  string // 回跳地址  回跳地址的格式必须是 URL
	Openid     string // openid
}

// PayOrderSubmitOutput 提交支付返回
type PayOrderSubmitOutput struct {
	TradeType string `json:"tradeType" ` // 交易类型
	//PayURL     string                 `json:"payURL" `     // 支付地址
	//OutTradeNo string                 `json:"outTradeNo" ` // 商户订单号
	JsApi *wechat.JSAPIPayParams `json:"jsApi"  ` // jsapi支付参数

	Status int // 支付状态 0未支付/10支付成功/20已退款/30支付关闭
}

// PayOrderNotifyInput 异步通知
type PayOrderNotifyInput struct {
	PayType string `json:"payType"` // 支付类型
	No      string `json:"no"`      // 支付订单号
}

// PayOrderNotifyOutput 异步通知返回
type PayOrderNotifyOutput struct {
	PayType string `json:"payType"` // 支付类型
	Code    string `json:"code"`    //  状态码
	Message string `json:"message"` // 响应文本消息
}
