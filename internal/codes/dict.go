package codes

import "github.com/gogf/gf/v2/errors/gcode"

var (
	// 10:主系统 05:字典模块 XX:具体错误
	CodeDictHasChild = gcode.New(100501, "请先删除子级", nil)
)
