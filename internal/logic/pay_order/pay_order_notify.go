package pay_order

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/payment"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
)

// Notify 异步通知
func (s *sPayOrder) Notify(ctx context.Context, in dto.PayOrderNotifyInput) (res *dto.PayOrderNotifyOutput, err error) {
	// 获取支付拓展单
	extension, err := s.GetExtensionInfo(ctx, in.No)
	if err != nil {
		return nil, err
	}
	if extension.Status == consts.PayOrderStatusSuccess { //  如果已经是成功，直接返回，不用重复更新
		return nil, nil
	}
	if extension.Status != consts.PayOrderStatusUnpaid {
		return nil, gerror.New("支付订单状态须待支付")
	}

	// 解析出支付配置
	var payCfg dto.PayConfig
	if err = gjson.DecodeTo(extension.ChannelExtras, &payCfg); err != nil {
		return nil, err
	}

	// 设置当前支付配置
	payment.SetConfig(&payCfg)

	// 获取支付回调通知参数
	data, err := payment.New(in.PayType).Notify(ctx, dto.NotifyInput{})
	if err != nil {
		return
	}

	// 更新 pay_order_extension
	extensionUp, err := dao.PayOrderExtension.Ctx(ctx).
		WherePri(extension.Id).
		Where("status", extension.Status).
		Data(do.PayOrderExtension{
			Status:            consts.PayOrderStatusSuccess,
			ChannelNotifyData: gjson.MustEncodeString(data),
		}).
		Update()
	if err != nil {
		return nil, err
	}
	if affectCnt, _ := extensionUp.RowsAffected(); affectCnt == 0 {
		return nil, gerror.New("支付交易拓展单不处于待支付")
	}

	// 查询支付订单
	payOrder, err := s.GetInfo(ctx, extension.PayOrderId)
	if err != nil {
		return nil, err
	}
	if payOrder.Status == consts.PayOrderStatusSuccess && payOrder.ExtensionId == extension.Id { //  如果已经是成功，直接返回，不用重复更新
		return nil, nil
	}

	if payOrder.Status != consts.PayOrderStatusUnpaid {
		return nil, gerror.Newf("商户订单号[%v]已被处理，请勿重复操作", data.OutTradeNo)
	}

	// 更新支付订单
	payOrderUp, err := dao.PayOrder.Ctx(ctx).
		WherePri(payOrder.Id).
		Where("status", consts.PayOrderStatusUnpaid).
		Data(do.PayOrder{
			Status:         consts.PayOrderStatusSuccess,
			SuccessTime:    data.PayAt,
			ExtensionId:    extension.Id,
			No:             extension.No,
			ChannelOrderNo: data.TransactionId,
		}).Update()
	if err != nil {
		return
	}

	if affectCnt2, _ := payOrderUp.RowsAffected(); affectCnt2 == 0 {
		return nil, gerror.New("支付订单状态须待支付")
	}

	// 回调业务订单同步通知, 注册请看：mall插件 main.go
	payOrder.Status = consts.PayOrderStatusSuccess
	payOrder.SuccessTime = data.PayAt
	payOrder.ExtensionId = extension.Id
	payOrder.No = extension.No
	payOrder.ChannelOrderNo = data.TransactionId
	if err = payment.NotifyCall(ctx, dto.NotifyCallFuncInput{PayOrder: payOrder}); err != nil {
		return
	}
	return
}
