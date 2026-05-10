package codes

import "github.com/gogf/gf/v2/errors/gcode"

// 10:主系统 02:菜单模块 XX:具体错误
var (
	CodeMenuNotFound  = gcode.New(100201, "菜单不存在", nil)
	CodeMenuHasChild  = gcode.New(100202, "请先删除所有子节点", nil)
	CodeMenuNoApiAuth = gcode.New(100203, "请先配置接口权限", nil)
)
