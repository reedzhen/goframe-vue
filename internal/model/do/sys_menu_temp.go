// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenuTemp is the golang structure of table sys_menu_temp for DAO operations like Where/Data.
type SysMenuTemp struct {
	g.Meta    `orm:"table:sys_menu_temp, do:true"`
	Id        any         // 菜单id
	ParentId  any         // 上级id, 0是顶级
	Title     any         // 菜单名称
	Path      any         // 菜单路由地址
	Component any         // 菜单组件地址, 目录可为空
	MenuType  any         // 类型 menu菜单/button按钮
	Sort      any         // 排序号
	Authority any         // 按钮权限
	Icon      any         // 菜单图标
	Hide      any         // 是否隐藏 1是/2否(仅注册路由不显示在左侧菜单)
	MenuMeta  *gjson.Json // 路由元信息
	ApiPath   *gjson.Json // 接口权限 默认存数组
	DeletedAt *gtime.Time // 删除时间
	CreatedBy any         // 添加人
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy any         // 更新人
	UpdatedAt *gtime.Time // 更新时间
}
