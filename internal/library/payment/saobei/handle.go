package saobei

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/model/dto"
	"time"

	"gitee.com/cleanpay/cleanpay/saobei"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func New(cfg *dto.PayConfig) *saobeiPay {
	return &saobeiPay{config: cfg}
}

type saobeiPay struct {
	config *dto.PayConfig
}

// CreateOrder 创建订单
func (h *saobeiPay) CreateOrder(ctx context.Context, in dto.CreateOrderInput) (res *dto.CreateOrderOutput, err error) {
	switch in.TradeType {
	case consts.TradeTypeSaobeiMini:
		return h.miniPay(ctx, in)
	default:
		err = gerror.Newf("暂未支持的交易方式：%v", in.TradeType)
	}
	return
}

// PayOrder 支付接口
func (h *saobeiPay) miniPay(ctx context.Context, in dto.CreateOrderInput) (res *dto.CreateOrderOutput, err error) {
	client, err := h.getClient()
	if err != nil {
		return
	}
	//amount := decimal.NewFromFloat(in.PayAmount).Mul(decimal.NewFromInt(100)).IntPart()
	bm := make(gopay.BodyMap)
	bm.Set("pay_type", saobei.PayTypeWX).
		Set("terminal_ip", "127.0.0.1").
		Set("terminal_trace", in.OutTradeNo).
		Set("terminal_time", time.Now().Format("20060102150405")).
		Set("total_fee", in.Price).
		Set("sub_appid", h.config.WXSubAppId).
		Set("open_id", in.Openid).
		Set("notify_url", in.NotifyUrl)

	resp, err := client.MiniPay(ctx, bm)
	if err != nil {
		return nil, err
	}
	if resp.ResultCode != saobei.ResultCodeSuccess {
		return nil, gerror.New(resp.ReturnMsg)
	}

	res = new(dto.CreateOrderOutput)
	res.JsApi = &wechat.JSAPIPayParams{
		AppId:     resp.AppId,
		TimeStamp: resp.TimeStamp,
		NonceStr:  resp.NonceStr,
		Package:   resp.PackageStr,
		SignType:  resp.SignType,
		PaySign:   resp.PaySign,
	}
	return res, nil
}

// Notify 获取支付回调返回参数
func (h *saobeiPay) Notify(ctx context.Context, in dto.NotifyInput) (res *dto.NotifyOutput, err error) {
	var result *SaobeiNotifyReq
	if err = g.RequestFromCtx(ctx).Parse(&result); err != nil {
		return
	}
	res = &dto.NotifyOutput{
		OutTradeNo:    result.TerminalTrace, // 商户系统的订单号，系统原样返回
		TransactionId: result.OutTradeNo,    // 扫呗的订单号
		ActualAmount:  gconv.Int(result.TotalFee),
		PayAt:         gtime.Now(),
	}
	return
}

// Refund 退款
func (h *saobeiPay) Refund(ctx context.Context, in dto.RefundInput) (res *dto.RefundOutput, err error) {
	client, err := h.getClient()
	if err != nil {
		return
	}

	// 组织参数
	bm := make(gopay.BodyMap)
	bm.Set("pay_type", saobei.PayTypeWX).
		Set("terminal_trace", in.RefundNo).
		Set("terminal_time", time.Now().Format("20060102150405")).
		Set("refund_fee", in.RefundMoney).
		Set("out_trade_no", in.TransactionId)

	resp, err := client.Refund(context.Background(), bm)
	if err != nil {
		return nil, err
	}
	if resp.ResultCode != saobei.ResultCodeSuccess {
		g.Log().Error(ctx, "saobei.Refund", resp.ReturnMsg)
		return nil, gerror.New(resp.ReturnMsg)
	}
	return
}

//
//// AddSharer 分账人维护
//func (h *saobeiPay) AddSharer(ctx context.Context, in payvo.AddSharerInput) (res *payvo.AddSharerOutput, err error) {
//	client, err := h.getClient()
//	if err != nil {
//		return
//	}
//	//组织参数
//	bm := make(gopay.BodyMap)
//	bm.Set("pay_ver", "100").
//		Set("trace_no", time.Now().Format("20060102150405")+grand.Digits(18)).
//		Set("rule_list_json", client.GenAllocateRuleList([]saobei.AllocateRule{
//			{
//				AccountIn:     in.Account,
//				AllocateScale: int(in.Rate) * 100,
//			},
//		})).
//		Set("start_date", gtime.Now().Layout("20060102")).
//		Set("end_date", gtime.Now().AddDate(10, 0, 0).Layout("20060102"))
//	//对私账户手机号需要赋值
//	if in.Name != "" {
//		bm = bm.Set("phone_no", in.Name)
//	}
//	resp, err := client.GenerateContract(ctx, bm)
//	if err != nil {
//		return nil, err
//	}
//	if resp.ResultCode != saobei.ResultCodeSuccess {
//		return nil, gerror.New(resp.ReturnMsg)
//	}
//	return
//}

//// SignShare 签署分账
//func (h *saobeiPay) SignShare(ctx context.Context, in payvo.SignShareInput) (res *payvo.SignShareOutput, err error) {
//	client, err := h.getClient()
//	if err != nil {
//		return nil, err
//	}
//	//组织参数
//	bm := make(gopay.BodyMap)
//	bm.Set("trace_no", time.Now().Format("20060102150405")+grand.Digits(18)).
//		Set("contract_no", in.ContractNo)
//	if in.VerifyAmt > 0 {
//		bm = bm.Set("verify_amt", uint(in.VerifyAmt*100))
//	}
//	if strings.TrimSpace(in.VerifyNo) != "" {
//		bm = bm.Set("verify_no", in.VerifyNo)
//	}
//	resp, err := client.SignContract(ctx, bm)
//	if err != nil {
//		return nil, err
//	}
//	if resp.ResultCode != saobei.ResultCodeSuccess {
//		return nil, gerror.New(resp.ReturnMsg)
//	}
//	return
//}

//// DoShare 执行分账
//func (h *saobeiPay) DoShare(ctx context.Context, in payvo.DoShareInput) (res *payvo.DoShareOutput, err error) {
//	client, err := h.getClient()
//	if err != nil {
//		return nil, err
//	}
//	//组织参数
//	bm := make(gopay.BodyMap)
//	bm.Set("contract_no", in.ContractNo).
//		Set("terminal_trace", in.ShareNo).
//		Set("terminal_time", time.Now().Format("20060102150405")).
//		Set("out_trade_no", in.OutOrderNo).
//		Set("rule_list_json", DoAllocateRuleList(in.Items))
//	resp, err := client.DoAllocate(ctx, bm)
//	if err != nil {
//		return nil, err
//	}
//	if resp.ResultCode != saobei.ResultCodeSuccess {
//		return nil, gerror.New(resp.ReturnMsg)
//	}
//	err = gconv.Struct(resp, &res)
//	return
//}

func (h *saobeiPay) getClient() (client *saobei.Client, err error) {
	client, err = saobei.NewClient(h.config.SaobeiInstNo, h.config.SaobeiKey, h.config.PayNo,
		h.config.SaobeiTerminalId, h.config.SaobeiAccessToken, true)
	return
}
