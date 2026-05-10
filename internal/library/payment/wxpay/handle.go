package wxpay

import (
	"context"
	"crypto/rsa"
	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/model/dto"
	"time"
)

func New(config *dto.PayConfig) *wxPay {
	return &wxPay{
		config: config,
	}
}

type wxPay struct {
	config *dto.PayConfig
}

// Refund 订单退款
func (h *wxPay) Refund(ctx context.Context, in dto.RefundInput) (res *dto.RefundOutput, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.Set("sub_mchid", h.config.WXSubMchId)   // 子商户ID
	bm.Set("transaction_id", in.TransactionId) // 微信、扫呗的订单号
	bm.Set("out_refund_no", in.RefundNo)       // 商户自定义退款单号
	bm.Set("reason", in.Reason)                // 申请退款原因
	bm.Set("amount", map[string]interface{}{
		"total":    in.TotalMoney,
		"refund":   in.RefundMoney,
		"currency": "CNY",
	})

	refund, err := client.V3Refund(ctx, bm)
	if err != nil {
		return
	}

	if refund.Error != "" {
		g.Log().Errorf(ctx, "微信退款失败,原因：%v", refund)
		return nil, gerror.Newf("微信支付发起退款失败,原因：%v", refund.Error)
	}

	if refund.Response.Status != "SUCCESS" && refund.Response.Status != "PROCESSING" {
		g.Log().Errorf(ctx, "微信退款失败,原因：%v", refund)
		return nil, gerror.Newf("微信支付发起退款失败,状态码：%v", refund.Response.Status)
	}
	return
}

// Notify 异步通知
func (h *wxPay) Notify(ctx context.Context, in dto.NotifyInput) (res *dto.NotifyOutput, err error) {
	notifyReq, err := wechat.V3ParseNotify(ghttp.RequestFromCtx(ctx).Request)
	if err != nil {
		return
	}

	client, err := GetClient(h.config)
	if err != nil {
		return
	}

	// 获取微信平台证书
	certMap, err := getPublicKeyMap(client)
	if err != nil {
		return
	}

	// 验证异步通知的签名
	if err = notifyReq.VerifySignByPKMap(certMap); err != nil {
		return
	}

	notify, err := notifyReq.DecryptPayCipherText(h.config.WXPayApiKey)
	if err != nil {
		return
	}

	if notify.TradeState != "SUCCESS" {
		err = gerror.New("非交易支付成功状态，无需处理！")
		// 这里如果相对非交易支付成功状态进行处理，可自行调整此处逻辑
		// ...
		return
	}

	if notify.OutTradeNo == "" {
		err = gerror.New("订单中没有找到商户单号！")
		return
	}

	res = new(dto.NotifyOutput)
	res.TransactionId = notify.TransactionId
	res.OutTradeNo = notify.OutTradeNo
	res.PayAt = gtime.New(notify.SuccessTime)
	res.ActualAmount = notify.Amount.PayerTotal // 单位：分
	return
}

// CreateOrder 创建订单
func (h *wxPay) CreateOrder(ctx context.Context, in dto.CreateOrderInput) (res *dto.CreateOrderOutput, err error) {
	switch in.TradeType {
	//case consts.TradeTypeWxScan:
	//	return h.scan(ctx, in)
	case consts.TradeTypeWxMP, consts.TradeTypeWxMini:
		return h.jsapi(ctx, in)
	//case consts.TradeTypeWxH5:
	//	return h.h5(ctx, in)
	default:
		err = gerror.Newf("暂未支持的交易方式：%v", in.TradeType)
	}
	return
}

func GetClient(config *dto.PayConfig) (client *wechat.ClientV3, err error) {
	client, err = wechat.NewClientV3(config.WXSpMchId, config.WXPaySerial, config.WXPayApiKey, config.WxPayPrivateKey)
	if err != nil {
		return
	}

	if _, _, err = client.GetAndSelectNewestCertALL(); err != nil {
		return nil, err
	}

	serialNo, snCertMap, err := client.GetAndSelectNewestCert()
	if err != nil {
		return
	}
	snPkMap := make(map[string]*rsa.PublicKey)
	for sn, cert := range snCertMap {
		pubKey, err := xpem.DecodePublicKey([]byte(cert))
		if err != nil {
			return nil, err
		}
		snPkMap[sn] = pubKey
	}

	client.SnCertMap = snPkMap
	client.WxSerialNo = serialNo

	// 打开Debug开关，输出日志，默认关闭
	if config.Debug {
		client.DebugSwitch = gopay.DebugOn
	}
	return
}

func getPublicKeyMap(client *wechat.ClientV3) (wxPublicKeyMap map[string]*rsa.PublicKey, err error) {
	serialNo, snCertMap, err := client.GetAndSelectNewestCert()
	if err != nil {
		return
	}

	snPkMap := make(map[string]*rsa.PublicKey)
	for sn, cert := range snCertMap {
		pubKey, err := xpem.DecodePublicKey([]byte(cert))
		if err != nil {
			return nil, err
		}
		snPkMap[sn] = pubKey
	}
	client.SnCertMap = snPkMap
	client.WxSerialNo = serialNo

	wxPublicKeyMap = client.WxPublicKeyMap()
	return
}

// // scan 创建扫码支付订单
//
//	func (h *wxPay) scan(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
//		client, err := GetClient(h.config)
//		if err != nil {
//			return
//		}
//
//		bm := make(gopay.BodyMap)
//		bm.Set("appid", h.config.WxPayAppId).
//			Set("mchid", h.config.WxPayMchId).
//			Set("description", in.PayOrder.Subject).
//			Set("out_trade_no", in.PayOrder.OutTradeNo).
//			Set("time_expire", time.Now().Add(2*time.Hour).Format(time.RFC3339)).
//			Set("notify_url", in.PayOrder.NotifyUrl).
//			SetBodyMap("amount", func(bm gopay.BodyMap) {
//				bm.Set("total", int64(in.PayOrder.PayAmount*100)).
//					Set("currency", "CNY")
//			})
//
//		wxRsp, err := client.V3TransactionNative(ctx, bm)
//		if err != nil {
//			return
//		}
//
//		if wxRsp.Code != 0 {
//			err = gerror.New(wxRsp.Error)
//			return
//		}
//
//		res = new(payin.CreateOrderModel)
//		res.TradeType = in.PayOrder.TradeType
//		res.PayURL = wxRsp.Response.CodeUrl
//		res.OutTradeNo = in.PayOrder.OutTradeNo
//		return
//	}
//
// // h5 创建H5支付订单
//
//	func (h *wxPay) h5(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error) {
//		client, err := GetClient(h.config)
//		if err != nil {
//			return
//		}
//
//		// 初始化参数Map
//		bm := make(gopay.BodyMap)
//		bm.Set("appid", h.config.WxPayAppId).
//			Set("mchid", h.config.WxPayMchId).
//			Set("description", in.PayOrder.Subject).
//			Set("out_trade_no", in.PayOrder.OutTradeNo).
//			Set("time_expire", time.Now().Add(2*time.Hour).Format(time.RFC3339)).
//			Set("notify_url", in.PayOrder.NotifyUrl).
//			SetBodyMap("amount", func(b gopay.BodyMap) {
//				b.Set("total", int64(in.PayOrder.PayAmount*100)).
//					Set("currency", "CNY")
//			}).
//			SetBodyMap("scene_info", func(b gopay.BodyMap) {
//				b.Set("payer_client_ip", in.PayOrder.CreateIp).
//					SetBodyMap("h5_info", func(b gopay.BodyMap) {
//						b.Set("type", "Wap")
//					})
//			})
//
//		// 请求支付下单，成功后得到结果
//		wxRsp, err := client.V3TransactionH5(ctx, bm)
//		if err != nil {
//			return
//		}
//
//		if wxRsp.Code != 0 {
//			err = gerror.New(wxRsp.Error)
//			return
//		}
//
//		res = new(payin.CreateOrderModel)
//		res.TradeType = in.PayOrder.TradeType
//		res.PayURL = wxRsp.Response.H5Url
//		res.OutTradeNo = in.PayOrder.OutTradeNo
//		return
//	}
//

// jsapi 创建jsapi支付订单
func (h *wxPay) jsapi(ctx context.Context, in dto.CreateOrderInput) (res *dto.CreateOrderOutput, err error) {
	client, err := GetClient(h.config)
	if err != nil {
		return
	}

	bm := make(gopay.BodyMap)
	bm.Set("description", in.Subject)
	bm.Set("out_trade_no", in.OutTradeNo)
	bm.Set("time_expire", time.Now().Add(59*time.Minute).Format(time.RFC3339))
	bm.Set("notify_url", in.NotifyUrl)
	bm.SetBodyMap("amount", func(bm gopay.BodyMap) {
		bm.Set("total", in.Price).
			Set("currency", "CNY")
	})
	bm.Set("sp_appid", h.config.WXSpAppId)
	bm.Set("sp_mchid", h.config.WXSpMchId)
	bm.Set("sub_mchid", h.config.WXSubMchId)
	bm.Set("sub_appid", h.config.WXSubAppId)
	bm.SetBodyMap("payer", func(bm gopay.BodyMap) {
		bm.Set("sub_openid", in.Openid)
	})
	bm.SetBodyMap("scene_info", func(b gopay.BodyMap) {
		b.Set("payer_client_ip", in.UserIP)
	})

	wxRsp, err := client.V3PartnerTransactionJsapi(ctx, bm)
	if err != nil {
		return
	}
	if wxRsp.Code != wechat.Success {
		err = gerror.New(wxRsp.Error)
		return
	}

	// 获取 JSAPI 支付所需要的参数
	jsApi, err := client.PaySignOfJSAPI(h.config.WXSubAppId, wxRsp.Response.PrepayId)
	if err != nil {
		return nil, err
	}

	return &dto.CreateOrderOutput{
		TradeType: in.TradeType,
		JsApi:     jsApi,
	}, nil
}
