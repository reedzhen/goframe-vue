package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int    `json:"code" dc:"错误码 0成功/大于1失败"`
	Message string `json:"message" dc:"消息提示"`
	Data    any    `json:"data" dc:"返回数据"`
	//Timestamp int64       `json:"timestamp" dc:"服务器时间戳"`       // 服务器时间戳
	//TraceId   string      `json:"trace_id" dc:"链路ID"`          // 链路ID
}

// Json 返回标准JSON数据
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}

	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
		//Timestamp: gtime.Timestamp(),
		//TraceId:   gctx.CtxId(r.Context()),
	})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}
