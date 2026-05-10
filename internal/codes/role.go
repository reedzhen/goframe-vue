package codes

import "github.com/gogf/gf/v2/errors/gcode"

// 10:主系统 03:角色模块 XX:具体错误
var (
	CodeRoleNotFound = gcode.New(100301, "角色不存在", nil)
	CodeRoleHasChild = gcode.New(100302, "请先删除子级角色", nil)
	CodeRoleInUse    = gcode.New(100303, "当前角色已被使用", nil)
)
