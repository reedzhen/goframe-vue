package codes

import "github.com/gogf/gf/v2/errors/gcode"

var (
	// 10:主系统 07:定时任务模块 XX:具体错误
	CodeCronNotFound = gcode.New(100701, "定时任务不存在", nil)
)
