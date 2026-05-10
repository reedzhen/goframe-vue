package dto

import "github.com/gogf/gf/v2/encoding/gjson"

type MenuItem struct {
	Id        int64  `json:"id"           `      // 主键ID
	ParentId  int64  `json:"parentId"          ` // 父级ID
	Title     string `json:"title"        `      // 菜单标题
	Icon      string `json:"icon"         `      // 图标
	Hide      uint   `json:"hide"         `      // 是否隐藏 0显示 1隐藏
	Path      string `json:"path"         `      // 菜单路径  page name
	Component string `json:"component"    `      // 菜单组件
	Sort      int    `json:"sort"         `      // 显示顺序

	//Children []*MenuItem `json:"children" orm:"-"`
	Checked bool `json:"checked" orm:"-"`
}

// MenuGetListInput 菜单获取列表
type MenuGetListInput struct {
	Title     string // 菜单名称
	Path      string // 菜单路由地址
	Authority string // 权限标识
	MenuType  int    // 类型
	ParentId  int64  // 上级id, 0是顶级
}

//// MenuGetAuthorizedListOutput 获取有权限的菜单列表
//type MenuGetAuthorizedListOutput struct {
//	Id        int64  `json:"id"         ` // 菜单id
//	ParentId  int64  `json:"parent_id"  ` // 上级id, 0是顶级
//	Title     string `json:"title"      ` // 菜单名称
//	Path      string `json:"path"       ` // 菜单路由地址
//	Component string `json:"component"  ` // 菜单组件地址, 目录可为空
//	MenuType  uint   `json:"menu_type"  ` // 类型, 0菜单, 1按钮
//	Sort      int    `json:"sort"       ` // 排序号
//	Authority string `json:"authority"  ` // 权限标识
//	Target    string `json:"target"     ` // 打开位置
//	Icon      string `json:"icon"       ` // 菜单图标
//	Color     string `json:"color"      ` // 图标颜色
//	Hide      uint   `json:"hide"       ` // 是否隐藏, 0否, 1是(仅注册路由不显示在左侧菜单)
//	Active    string `json:"active"     ` // 菜单侧栏选中的path
//	MenuMeta  string `json:"meta"`        // 其它路由元信息
//}

type MenuCreateUpdateBase struct {
	ParentId  int64       // 上级id, 0是顶级
	Title     string      // 菜单名称
	Path      string      // 菜单路由地址
	Component string      // 菜单组件地址
	MenuType  int         // 类型 menu菜单/button按钮
	Sort      int         // 排序号
	Authority string      // 权限标识
	Icon      string      // 菜单图标
	Hide      int         // 状态 是否隐藏  2否/1是
	MenuMeta  *gjson.Json // 路由元信息
	ApiPath   *gjson.Json // 接口权限 默认存数组
}

// MenuCreateInput 菜单新增
type MenuCreateInput struct {
	CreatedBy int64 // 添加人
	MenuCreateUpdateBase
}

// MenuUpdateInput 菜单编辑
type MenuUpdateInput struct {
	Id        int64 // 菜单id
	UpdatedBy int64
	MenuCreateUpdateBase
}

// MenuApiPathOutput 菜单接口权限返回
type MenuApiPathOutput struct {
	ApiPath []string `json:"api_path"` // 接口权限
}
