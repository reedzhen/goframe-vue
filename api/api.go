package api

import "github.com/gogf/gf/v2/encoding/gjson"

type MenuTree struct {
	Id        int64       `json:"id"        dc:"菜单id"        `
	ParentId  int64       `json:"parentId"  dc:"上级id, 0是顶级"  `
	Title     string      `json:"title"     dc:"菜单名称"     `
	Path      string      `json:"path"     dc:"菜单路由地址" `
	Component string      `json:"component" dc:"菜单组件地址, 目录可为空"  `
	MenuType  int         `json:"menuType"  dc:"类型 0菜单/1按钮"  `
	Sort      int         `json:"sort"      dc:"排序号"     `
	Authority string      `json:"authority" dc:"权限标识" `
	Icon      string      `json:"icon"      dc:"菜单图标"      `
	Hide      int         `json:"hide"      dc:"是否隐藏 0否/1是(仅注册路由不显示在左侧菜单)"      `
	MenuMeta  string      `json:"menuMeta"  dc:"路由元信息" `
	ApiPath   *gjson.Json `json:"apiPath"   dc:"接口权限 默认存数组"   `
	Children  []*MenuTree `json:"children,omitempty"`
}

type MenuItem struct {
	Id        int64  `json:"id"           `      // 主键ID
	ParentId  int64  `json:"parentId"          ` // 父级ID
	Title     string `json:"title"        `      // 菜单标题
	Icon      string `json:"icon"         `      // 图标
	Hide      uint   `json:"hide"         `      // 是否隐藏 0显示 1隐藏
	Path      string `json:"path"         `      // 菜单路径  page name
	Component string `json:"component"    `      // 菜单组件
	Sort      int    `json:"sort"         `      // 显示顺序
	Checked   bool   `json:"checked" orm:"-"`
}
