package payment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

const (
	TradeOrderNoPrefix = "o"
	AfterSaleNoPrefix  = "r"
	KeyTradeNo         = "trade_no:" // 业务订单号的缓存

	OrderNoPrefix  = "P"       // 支付订单 no 的前缀
	RefundNoPrefix = "R"       // 退款订单 no 的前缀
	KeyPayNo       = "pay_no:" // 支付订单号的缓存
)

// GenOrderNo 生成订单号的通用方法
func GenOrderNo(ctx context.Context, prefix string, cachePrefix string, expireTimes ...int) (string, error) {
	// 设置默认过期时间为 60 秒
	expireTime := 60
	if len(expireTimes) > 0 {
		expireTime = expireTimes[0]
	}

	noPrefix := fmt.Sprintf("%s%s", prefix, gtime.Now().Format("ymdHis"))
	cacheKey := cachePrefix + noPrefix

	// 递增序号并获取结果 同一秒多个订单号递增
	incrementResult, err := g.Redis().Do(ctx, "INCR", cacheKey)
	if err != nil {
		return "", err
	}

	// 设置过期时间
	if _, err := g.Redis().Do(ctx, "EXPIRE", cacheKey, expireTime); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%d", noPrefix, incrementResult.Int()), nil
}

// GenTradeOrderNo 生成业务订单号
func GenTradeOrderNo(ctx context.Context, prefix string) (string, error) {
	return GenOrderNo(ctx, prefix, KeyTradeNo)
}

// GenPayOrderNo 生成支付订单号或支付退款单号
func GenPayOrderNo(ctx context.Context, prefix string) (string, error) {
	return GenOrderNo(ctx, prefix, KeyPayNo)
}

//// GenTradeOrderNo 生成业务订单号
//func GenTradeOrderNo(ctx context.Context, prefix string) (string, error) {
//	noPrefix := fmt.Sprintf("%s%s", prefix, gtime.Now().Format("ymdHis"))
//	cacheKey := KeyTradeNo + noPrefix
//
//	// 递增序号并获取结果 同一秒多个订单号递增
//	incrementResult, err := g.Redis().Do(ctx, "Incr", cacheKey)
//	if err != nil {
//		return "", err
//	}
//
//	// 设置过期时间 秒
//	if _, err := g.Redis().Do(ctx, "EXPIRE", cacheKey, 60); err != nil {
//		return "", err
//	}
//
//	return fmt.Sprintf("%s%d", noPrefix, incrementResult.Int()), nil
//}
//
//// GenPayOrderNo 生成支付订单号or支付退款单号
//func GenPayOrderNo(ctx context.Context, prefix string) (string, error) {
//	noPrefix := fmt.Sprintf("%s%s", prefix, gtime.Now().Format("ymdHis"))
//	cacheKey := KeyPayNo + noPrefix
//
//	// 递增序号并获取结果 同一秒多个订单号递增
//	incrementResult, err := g.Redis().Do(ctx, "Incr", cacheKey)
//	if err != nil {
//		return "", err
//	}
//
//	// 设置过期时间
//	if _, err := g.Redis().Do(ctx, "EXPIRE", cacheKey, 60); err != nil {
//		return "", err
//	}
//
//	return fmt.Sprintf("%s%d", noPrefix, incrementResult.Int()), nil
//}
