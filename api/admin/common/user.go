package common

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CallbackDemoReq struct {
	g.Meta     `path:"/callback/demo" method:"post" tags:"通用接口" summary:"测试回调用" noAuth:"true"`
	ClientId   string `json:"clientId"`   // 客户编号
	ModuleCode string `json:"moduleCode"` // 模块
	Type       string `json:"type"`       // 类型 openModule/paySuccess
}
type CallbackDemoRes struct {
}
