package codes

import "github.com/gogf/gf/v2/errors/gcode"

// 10:主系统 08:认证模块 XX:具体错误
var (
	CodeAuthCaptchaError    = gcode.New(100801, "验证码错误或已过期", nil)
	CodeAuthCredentialError = gcode.New(100802, "账号或密码错误", nil)
	CodeAuthUserFrozen      = gcode.New(100803, "您的账号已被冻结，请联系管理员", nil)
	CodeAuthRoleDisabled    = gcode.New(100804, "您的账号所属角色已被禁用，请联系管理员", nil)
)
