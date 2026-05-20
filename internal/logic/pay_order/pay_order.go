package pay_order

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
	"goframe-vben/api/admin/common"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/payment"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"strings"
	"time"
)

func init() {
	service.RegisterPayOrder(New())
}

// sPayOrder 通用支付
type sPayOrder struct{}

func New() *sPayOrder {
	return &sPayOrder{}
}

// CreateOrder 创建支付订单(创建商城订单时会调用此方法)
func (s *sPayOrder) CreateOrder(ctx context.Context, in dto.PayOrderCreateInput) (out int64, err error) {
	payOrder, err := s.GetInfoByMerchantOrderId(ctx, in.MerchantOrderId)
	if err != nil {
		return
	}
	if payOrder != nil {
		g.Log().Warningf(ctx, "[创建支付订单] MerchantOrderNo(%d) 已经存在对应的支付单(%v)", in.MerchantOrderId, payOrder) // 理论来说，不会出现这个情况
		return payOrder.Id, nil
	}

	//// 生成支付通知地址
	//notifyURL, err := s.GenNotifyURL(ctx, in)
	//if err != nil {
	//	return
	//}

	data := &entity.PayOrder{
		MerchantOrderId: in.MerchantOrderId,
		OrderGroup:      in.OrderGroup,
		Subject:         in.Subject,
		Body:            "",
		//NotifyUrl:       notifyURL, // 每次发起支付的时候动态写入
		PayType:        in.PayType,
		TradeType:      in.TradeType,
		Price:          in.PayAmount,
		Status:         consts.PayOrderStatusUnpaid,
		UserIp:         in.UserIp,
		ExpireTime:     gtime.Now().Add(40 * time.Minute), // 这个时间一定要>trade_order 失效配置时间，pay_order 要后失效 todo 后期写到配置文件里
		SuccessTime:    nil,                               // 支付成功后回调填写
		ExtensionId:    0,                                 // 支付成功后写入
		No:             "",                                // 支付成功后写入
		ChannelUserId:  "",
		ChannelOrderNo: "", // 支付成功后写入
		CreatedBy:      in.UserId,
	}

	// 创建支付订单
	newId, err := dao.PayOrder.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return 0, err
	}

	return newId, nil
}

// SubmitOrder 提交支付 此时，会发起支付渠道的调用
func (s *sPayOrder) SubmitOrder(ctx context.Context, in dto.PayOrderSubmitInput) (out *dto.PayOrderSubmitOutput, err error) {
	// 校验支付订单
	payOrder, err := s.validateOrderCanSubmit(ctx, in.PayOrderId)
	if err != nil {
		return
	}

	// 生成支付单号
	no, err := payment.GenPayOrderNo(ctx, payment.OrderNoPrefix)
	if err != nil {
		return
	}

	// 动态生成支付通知地址
	notifyURL, err := s.GenNotifyURL(ctx, dto.PayOrderGenNotifyURLInput{
		PayType: payOrder.PayType,
		No:      no,
	})
	if err != nil {
		return
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
		SaobeiTerminalId:  "xx",  //支付系统：商户号终端号
		SaobeiAccessToken: "xxx", //支付系统： 令牌
		PayNo:             "xx",  // 商户号
		PayId:             0,
		FeeRate:           0,
	}
	payment.SetConfig(payCfg)

	var ret *dto.CreateOrderOutput
	if err = dao.PayOrderExtension.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 创建支付订单拓展
		if _, err = dao.PayOrderExtension.Ctx(ctx).Data(do.PayOrderExtension{
			No:            no,
			PayOrderId:    payOrder.Id,
			UserIp:        in.UserIp,
			Status:        consts.PayOrderStatusUnpaid,
			CreatedBy:     payOrder.CreatedBy,
			ChannelExtras: gjson.MustEncodeString(payCfg), // 支付配置先放到这里
		}).Insert(); err != nil {
			return err
		}

		// 编辑支付订单通知地址
		if _, err := dao.PayOrder.Ctx(ctx).Data(do.PayOrder{
			NotifyUrl: notifyURL,
		}).WherePri(payOrder.Id).Update(); err != nil {
			return err
		}

		// 调用三方支付接口
		ret, err = payment.New(payOrder.PayType).CreateOrder(ctx, dto.CreateOrderInput{
			TradeType:  payOrder.TradeType,
			UserIP:     in.UserIp,
			OutTradeNo: no,
			Subject:    payOrder.Subject,
			Body:       payOrder.Body,
			NotifyUrl:  notifyURL,
			ReturnUrl:  in.ReturnUrl,
			Price:      payOrder.Price,
			ExpireTime: payOrder.ExpireTime,
			Openid:     in.Openid,
		})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return
	}

	return &dto.PayOrderSubmitOutput{
		TradeType: ret.TradeType,
		//PayURL:     ret.PayURL,
		//OutTradeNo: ret.OutTradeNo,
		JsApi:  ret.JsApi,
		Status: consts.PayOrderStatusUnpaid,
	}, nil
}

// validateOrderCanSubmit 校验支付订单
func (s *sPayOrder) validateOrderCanSubmit(ctx context.Context, payOrderId int64) (out *entity.PayOrder, err error) {
	payOrder, err := s.GetInfo(ctx, payOrderId)
	if err != nil {
		return
	}

	if payOrder.Status == consts.PayOrderStatusSuccess {
		return nil, gerror.New("订单已支付，请刷新页面")
	}
	if payOrder.Status != consts.PayOrderStatusUnpaid {
		return nil, gerror.New("支付订单状态须待支付")
	}
	if payOrder.ExpireTime.Before(gtime.Now()) {
		return nil, gerror.New("支付订单已过期")
	}
	// 【重要】校验是否支付拓展单已支付
	if err = s.validateOrderActuallyPaid(ctx, payOrderId); err != nil {
		return
	}
	return payOrder, nil
}

// validateOrderActuallyPaid 校验支付订单实际已支付
func (s *sPayOrder) validateOrderActuallyPaid(ctx context.Context, payOrderId int64) (err error) {
	list := make([]*entity.PayOrderExtension, 0)
	if err = dao.PayOrderExtension.Ctx(ctx).Where("pay_order_id", payOrderId).Scan(&list); err != nil {
		return
	}
	for _, v := range list {
		// 情况一：校验数据库中的 orderExtension 是不是已支付
		if v.Status == consts.PayOrderStatusSuccess {
			return gerror.New("订单已支付，请等待支付结果")
		}

		// todo 情况二：调用三方接口，查询支付单状态，是不是已支付
	}

	return
}

// GenNotifyURL 生成支付通知地址
func (s *sPayOrder) GenNotifyURL(ctx context.Context, in dto.PayOrderGenNotifyURLInput) (notifyURL string, err error) {
	path := ""
	var object interface{}
	switch in.PayType {
	case consts.PayTypeSaobeiPay:
		object = common.SaobeiPayNotifyReq{}
	case consts.PayTypeWxPay:
		object = common.WxPayNotifyReq{}
	default:
		err = gerror.Newf("未被支持的支付方式：%v", in.PayType)
		return
	}

	path = gmeta.Get(object, "path").String()
	path = strings.ReplaceAll(path, ":extensionNo", in.No)
	notifyURL = fmt.Sprintf("%s%s%s",
		g.Cfg().MustGet(ctx, "system.domain").String(),
		"/api",
		path,
	)

	return
}

// GetInfoByMerchantOrderId 获取支付订单详情 todo 实际情况下 app_id+merchant_order_id才能确定唯一一条记录，这里只有商城用，所以适用
func (s *sPayOrder) GetInfoByMerchantOrderId(ctx context.Context, merchantOrderId int64) (out *entity.PayOrder, err error) {
	if err = dao.PayOrder.Ctx(ctx).Where("merchant_order_id", merchantOrderId).Scan(&out); err != nil {
		return
	}
	return
}

// GetInfo 获取支付订单详情
func (s *sPayOrder) GetInfo(ctx context.Context, payOrderId int64) (out *entity.PayOrder, err error) {
	if err = dao.PayOrder.Ctx(ctx).WherePri(payOrderId).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.New("支付订单不存在")
	}
	return
}

// GetExtensionInfo 获取支付订单拓展单详情
func (s *sPayOrder) GetExtensionInfo(ctx context.Context, no string) (out *entity.PayOrderExtension, err error) {
	if err = dao.PayOrderExtension.Ctx(ctx).Where("no", strings.TrimSpace(no)).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.New("支付交易拓展单不存在")
	}
	return
}

// ExpirePayOrder 将已过期的订单，状态修改为已关闭(cron)
func (s *sPayOrder) ExpirePayOrder(ctx context.Context) {
	// 查询过期的待支付订单
	list := make([]*entity.PayOrder, 0)
	if err := dao.PayOrder.Ctx(ctx).
		Where("status", consts.PayOrderStatusUnpaid).
		WhereLT("expire_time", gtime.Now()).Scan(&list); err != nil || len(list) == 0 {
		return
	}

	cnt := 0
	for _, v := range list {
		if err := s.expireOneOrder(ctx, v); err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		cnt++
	}

	g.Log().Infof(ctx, "支付订单过期处理完成，共处理 %d 个", cnt)
	return
}

// expireOneOrder  将已过期的订单，状态修改为已关闭
func (s *sPayOrder) expireOneOrder(ctx context.Context, order *entity.PayOrder) error {
	// 获取支付拓展单列表
	list := make([]*entity.PayOrderExtension, 0)
	if err := dao.PayOrderExtension.Ctx(ctx).Where("pay_order_id", order.Id).Scan(&list); err != nil {
		return err
	}

	for _, v := range list {
		if v.Status == consts.PayOrderStatusClose {
			continue
		}
		// 情况一：校验数据库中的 orderExtension 是不是已支付
		if v.Status == consts.PayOrderStatusSuccess {
			return gerror.Newf("[expireOrder][order(%d) 的 extension(%d) 已支付，可能是数据不一致]", order.Id, v.Id)
		}

		// todo 情况二：调用三方接口，查询支付单状态，是不是已支付/已退款

		// 兜底逻辑：将支付拓展单更新为已关闭
		result, err := dao.PayOrderExtension.Ctx(ctx).
			WherePri(v.Id).
			Where("status", consts.PayOrderStatusUnpaid).
			Data(do.PayOrderExtension{Status: consts.PayOrderStatusClose}).Update()
		if err != nil {
			return err
		}

		if cnt, _ := result.RowsAffected(); cnt == 0 {
			return gerror.Newf("[expireOrder][extension(%d) 更新为支付关闭失败]", v.Id)
		}
	}

	// 都没有上述情况，可以安心更新为已关闭
	result1, err := dao.PayOrder.Ctx(ctx).
		WherePri(order.Id).
		Where("status", order.Status).
		Data(do.PayOrder{Status: consts.PayOrderStatusClose}).Update()
	if err != nil {
		return err
	}
	if cnt, _ := result1.RowsAffected(); cnt == 0 {
		return gerror.Newf("[expireOrder][order(%d) 更新为支付关闭失败]", order.Id)
	}
	return nil
}
