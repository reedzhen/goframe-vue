package payment

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-vben/internal/model/dto"
	"sync"
)

type NotifyCallFunc func(ctx context.Context, pay dto.NotifyCallFuncInput) (err error)

var (
	notifyCall = make(map[string]NotifyCallFunc)
	ncLock     sync.Mutex
)

// RegisterNotifyCall 注册支付成功回调方法
func RegisterNotifyCall(group string, f NotifyCallFunc) {
	ncLock.Lock()
	defer ncLock.Unlock()
	if _, ok := notifyCall[group]; ok {
		panic("notifyCall repeat registration, group:" + group)
	}
	notifyCall[group] = f
}

// RegisterNotifyCallMap 注册支付成功回调方法
func RegisterNotifyCallMap(calls map[string]NotifyCallFunc) {
	for group, f := range calls {
		RegisterNotifyCall(group, f)
	}
}

// NotifyCall 执行订单分组的同步回调
func NotifyCall(ctx context.Context, in dto.NotifyCallFuncInput) error {
	f, ok := notifyCall[in.PayOrder.OrderGroup]
	if ok {
		ctx = gctx.NeverDone(ctx)
		if err := f(ctx, in); err != nil {
			g.Log().Warningf(ctx, "payments.NotifyCall in:%+v exec err:%+v", gjson.New(in.PayOrder).String(), err)
			return err
		}
		//ctx =gctx.NeverDone(ctx)
		//simple.SafeGo(ctx, func(ctx context.Context) {
		//	if err := f(ctx, in); err != nil {
		//		g.Log().Warningf(ctx, "payment.NotifyCall in:%+v exec err:%+v", gjson.New(in.PayOrder).String(), err)
		//	}
		//})
	}

	return nil
}
