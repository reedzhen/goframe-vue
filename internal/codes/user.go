package codes

import "github.com/gogf/gf/v2/errors/gcode"

// 10:主系统 01:用户模块 XX:具体错误
var (
	CodeUserNotFound     = gcode.New(100101, "当前用户不存在", nil)
	CodeUserDuplicate    = gcode.New(100102, "账号已存在", nil)
	CodeUserDeleteDeny   = gcode.New(100103, "此账号不允许删除", nil)
	CodeUserDeleteSelf   = gcode.New(100104, "自己不能删除自己", nil)
	CodeUserOldPwdError  = gcode.New(100105, "旧密码输入错误", nil)
	CodeUserResetPwdDeny = gcode.New(100106, "此账号不允许重置密码", nil)
	CodeUserStatusDeny   = gcode.New(100107, "此账号不允许修改状态", nil)
	CodeUserStatusSelf   = gcode.New(100108, "自己不能修改自己的状态", nil)
)
