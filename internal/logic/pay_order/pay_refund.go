package pay_order

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/payment"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
)

func init() {
	service.RegisterPayRefund(NewRefund())
}

// sPayOrder 支付订单退款
type sPayRefund struct{}

func NewRefund() *sPayRefund {
	return &sPayRefund{}
}

// Refund 订单退款
func (s *sPayRefund) Refund(ctx context.Context, in dto.PayRefundInput) (res *entity.PayRefund, err error) {
	// 校验支付订单
	payOrder, err := s.validatePayOrderCanRefund(ctx, in)
	if err != nil {
		return nil, err
	}

	// 生成退款单号
	refundNo, err := payment.GenPayOrderNo(ctx, payment.RefundNoPrefix)
	if err != nil {
		return nil, err
	}

	// 获取并动态加载支付配置
	// 微信支付走服务商模式，只能获取联盟支付配置 payOrder.PartnerId
	// 扫呗支持获取门店支付配置 payOrder.PlaceId
	//mid := payOrder.PartnerId
	//if payOrder.PayType == consts.PayTypeSaobeiPay {
	//	mid = payOrder.PlaceId
	//}
	//payCfg, err := service.Pay().GetPayConfig(ctx, uint64(mid))
	//if err != nil {
	//	return nil, err
	//}

	// saobei本地测试用
	payCfg := &dto.PayConfig{
		Debug:             true,
		PaySource:         "saobei",
		WXSubAppId:        "xxx", // 子商户APPID
		SaobeiInstNo:      "",
		SaobeiKey:         "",
		SaobeiTerminalId:  "xxx", //支付系统：商户号终端号
		SaobeiAccessToken: "xx",  //支付系统： 令牌
		PayNo:             "xx",  // 商户号
		PayId:             0,
		FeeRate:           0,
	}
	payment.SetConfig(payCfg)

	// 创建第三方平台退款
	if _, err = payment.New(payOrder.PayType).Refund(ctx, dto.RefundInput{
		TransactionId: payOrder.ChannelOrderNo,
		TotalMoney:    payOrder.Price,
		RefundMoney:   in.RefundPrice,
		RefundNo:      refundNo,
		Reason:        in.Reason,
		//Remark:        in.Remark,
	}); err != nil {
		return
	}

	// 更新支付订单（使用原子操作增加退款金额）
	result, err := dao.PayOrder.Ctx(ctx).WherePri(payOrder.Id).Data(do.PayOrder{
		RefundPrice: gdb.Raw(fmt.Sprintf("refund_price + %d", in.RefundPrice)),
		Status:      consts.PayOrderStatusRefund,
	}).Update()
	if err != nil {
		return
	}
	ret, err := result.RowsAffected()
	if err != nil {
		return
	}
	if ret == 0 {
		return nil, gerror.Newf("业务订单号[%v]退款失败", in.MerchantOrderId)
	}

	// 创建退款记录
	if _, err = dao.PayRefund.Ctx(ctx).Data(do.PayRefund{
		No:                refundNo,
		PayOrderId:        payOrder.Id,
		PayOrderNo:        payOrder.No,
		MerchantOrderId:   payOrder.MerchantOrderId,
		MerchantRefundId:  0,
		NotifyUrl:         "",
		Status:            consts.PayRefundStatusSuccess, // 这里直接记录退款成功
		PayPrice:          payOrder.Price,
		RefundPrice:       in.RefundPrice,
		Reason:            in.Reason,
		UserIp:            g.RequestFromCtx(ctx).GetRemoteIp(),
		ChannelOrderNo:    payOrder.ChannelOrderNo,
		ChannelRefundNo:   nil,
		SuccessTime:       gtime.Now(),
		ChannelErrorCode:  nil,
		ChannelErrorMsg:   nil,
		ChannelNotifyData: nil,
		Remark:            in.Remark,
		CreatedBy:         in.CreatedBy,
	}).Insert(); err != nil {
		return
	}
	return
}

// validatePayOrderCanRefund 校验支付订单是否可以退款
func (s *sPayRefund) validatePayOrderCanRefund(ctx context.Context, in dto.PayRefundInput) (res *entity.PayOrder, err error) {
	payOrder, err := service.PayOrder().GetInfoByMerchantOrderId(ctx, in.MerchantOrderId)
	if err != nil {
		return nil, err
	}
	if payOrder == nil {
		return nil, gerror.New("支付订单不存在")
	}

	// 校验状态，必须是已支付、或者已退款
	if payOrder.Status != consts.PayOrderStatusRefund && payOrder.Status != consts.PayOrderStatusSuccess {
		return nil, gerror.New("支付订单退款失败，原因：状态不是已支付或已退款")
	}

	// 校验金额，退款金额不能大于原定的金额
	if payOrder.RefundPrice+in.RefundPrice > payOrder.Price {
		return nil, gerror.New("退款金额超过订单可退款金额")
	}

	// 是否有退款中的订单
	if cnt := s.selectCountByAppIdAndOrderId(ctx, dto.PayRefundExistInput{
		AppId:      payOrder.AppId,
		PayOrderId: payOrder.Id,
		Status:     consts.PayRefundStatusWaiting,
	}); cnt > 0 {
		return nil, gerror.New("已有退款在处理中")
	}

	return payOrder, nil
}

// selectCountByAppIdAndOrderId 查询当前支付订单是否有退款中的订单
func (s *sPayRefund) selectCountByAppIdAndOrderId(ctx context.Context, in dto.PayRefundExistInput) int {
	cnt, err := dao.PayRefund.Ctx(ctx).Where(do.PayRefund{
		AppId:      in.AppId,
		PayOrderId: in.PayOrderId,
		Status:     in.Status,
	}).Count()
	if err != nil {
		return 0
	}

	return cnt
}
