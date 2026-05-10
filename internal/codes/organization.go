package codes

import "github.com/gogf/gf/v2/errors/gcode"

// 10:主系统 04:组织模块 XX:具体错误
var (
	CodeOrgNotFound      = gcode.New(100401, "组织不存在", nil)
	CodeOrgParentSelf    = gcode.New(100402, "上级组织不能是自己", nil)
	CodeOrgHasChild      = gcode.New(100403, "请先删除子级组织", nil)
	CodeOrgInUse         = gcode.New(100404, "当前组织已被使用", nil)
	CodeOrgChildAsParent = gcode.New(100405, "不能将子级组织设置为上级组织", nil)
)
