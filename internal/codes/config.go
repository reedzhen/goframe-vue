package codes

import "github.com/gogf/gf/v2/errors/gcode"

// 10:主系统 06:配置模块 XX:具体错误
var (
	CodeConfigGroupEmpty = gcode.New(100601, "缓存分组名称不能为空", nil)
)
