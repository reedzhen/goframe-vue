// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenuTemp is the golang structure for table sys_menu_temp.
type SysMenuTemp struct {
	Id        int64       `json:"id"        orm:"id"         ` // 菜单id
	ParentId  int64       `json:"parentId"  orm:"parent_id"  ` // 上级id, 0是顶级
	Title     string      `json:"title"     orm:"title"      ` // 菜单名称
	Path      string      `json:"path"      orm:"path"       ` // 菜单路由地址
	Component string      `json:"component" orm:"component"  ` // 菜单组件地址, 目录可为空
	MenuType  string      `json:"menuType"  orm:"menu_type"  ` // 类型 menu菜单/button按钮
	Sort      int         `json:"sort"      orm:"sort"       ` // 排序号
	Authority string      `json:"authority" orm:"authority"  ` // 按钮权限
	Icon      string      `json:"icon"      orm:"icon"       ` // 菜单图标
	Hide      int         `json:"hide"      orm:"hide"       ` // 是否隐藏 1是/2否(仅注册路由不显示在左侧菜单)
	MenuMeta  *gjson.Json `json:"menuMeta"  orm:"menu_meta"  ` // 路由元信息
	ApiPath   *gjson.Json `json:"apiPath"   orm:"api_path"   ` // 接口权限 默认存数组
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // 删除时间
	CreatedBy int64       `json:"createdBy" orm:"created_by" ` // 添加人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" ` // 更新人
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
}
