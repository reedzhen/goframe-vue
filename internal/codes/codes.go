package codes

import "github.com/gogf/gf/v2/errors/gcode"

// 客户端错误 >=400 < 500
// 服务端错误 500
// 自定义错误 >900
var (
	CodeOK                  = gcode.New(0, "操作成功", nil)
	CodeBadRequest          = gcode.New(400, "请求参数不正确", nil)
	CodeUnauthorized        = gcode.New(401, "账号未登录", nil)
	CodeForbidden           = gcode.New(403, "没有该操作权限", nil)
	CodeNotFound            = gcode.New(404, "请求未找到", nil)
	CodeMethodNotAllowed    = gcode.New(405, "请求方法不支持", nil)
	CodeLocked              = gcode.New(423, "请求失败，请稍后重试", nil)
	CodeTooManyRequests     = gcode.New(429, "请求过于频繁，请稍后重试", nil)
	CodeInternalServerError = gcode.New(500, "系统异常", nil)
	CodeNotImplemented      = gcode.New(501, "接口正在开发中，还未完成", nil)
	CodeRepeatRequest       = gcode.New(900, "重复请求，请稍后重试", nil)
	CodeDemoDeny            = gcode.New(901, "演示模式，禁止写操作", nil)
	CodeUnknown             = gcode.New(999, "未知错误", nil)
)
