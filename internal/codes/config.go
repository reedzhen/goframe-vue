package codes

import "github.com/gogf/gf/v2/errors/gcode"

// 10:主系统 06:配置模块 XX:具体错误
var (
	CodeConfigGroupEmpty      = gcode.New(100601, "缓存分组名称不能为空", nil)
	CodeConfigModuleCodeExist = gcode.New(100602, "模块编码已存在", nil)
	CodeConfigModuleNotFound  = gcode.New(100603, "配置模块不存在", nil)
	CodeConfigModuleHasItem   = gcode.New(100604, "模块下存在配置项，无法删除", nil)
	CodeConfigItemKeyExist    = gcode.New(100605, "配置项键名在该模块下已存在", nil)
)
