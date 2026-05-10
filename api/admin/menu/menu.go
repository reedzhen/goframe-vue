package menu

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/api"
)

type GetListReq struct {
	g.Meta    `path:"/menu/list" method:"get" tags:"菜单管理" summary:"获取菜单列表" noAuth:"true"`
	Title     string `json:"title"      ` // 菜单名称
	Path      string `json:"path"       ` // 菜单路由地址
	Authority string `json:"authority"  ` // 权限标识
	MenuType  int    `json:"menuType"   ` // 类型, 0菜单, 1按钮
	ParentId  int64  `json:"parentId"   ` // 上级id, 0是顶级
}
type GetListRes []*api.MenuTree

type CreateReq struct {
	g.Meta    `path:"/menu/create" method:"post" tags:"菜单管理" summary:"新增菜单"`
	ParentId  int64       `json:"parentId" dc:"上级id, 0是顶级"`
	Title     string      `json:"title" v:"required" dc:"菜单名称"`
	Path      string      `json:"path" v:"required" dc:"菜单路由地址"`
	Component string      `json:"component" dc:"菜单组件地址"`
	MenuType  int         `json:"menuType" dc:"类型 1菜单/2按钮"`
	Sort      int         `json:"sort" v:"required" dc:"排序号"`
	Authority string      `json:"authority" dc:"权限标识"`
	Icon      string      `json:"icon" dc:"菜单图标"`
	Hide      int         `json:"hide" dc:"是否隐藏, 2否, 1是"`
	MenuMeta  *gjson.Json `json:"menuMeta" dc:"路由元信息"`
	ApiPath   *gjson.Json `json:"apiPath"  dc:"接口权限 默认存数组"`
}
type CreateRes struct{}

type UpdateReq struct {
	g.Meta    `path:"/menu/update" method:"post" tags:"菜单管理" summary:"修改菜单"`
	Id        int64       `json:"id" dc:"菜单id" `
	ParentId  int64       `json:"parentId" dc:"上级id, 0是顶级"`
	Title     string      `json:"title" v:"required" dc:"菜单名称"`
	Path      string      `json:"path"  dc:"菜单路由地址"`
	Component string      `json:"component" dc:"菜单组件地址"`
	MenuType  int         `json:"menuType" dc:"类型 1菜单/2按钮"`
	Sort      int         `json:"sort" v:"required" dc:"排序号"`
	Authority string      `json:"authority" dc:"权限标识"`
	Icon      string      `json:"icon" dc:"菜单图标"`
	Hide      int         `json:"hide" dc:"是否隐藏, 2否, 1是"`
	MenuMeta  *gjson.Json `json:"menuMeta" dc:"路由元信息"`
	ApiPath   *gjson.Json `json:"apiPath"  dc:"接口权限 默认存数组"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/menu/delete/{Id}" method:"post" tags:"菜单管理" summary:"删除菜单"`
	Id     int64 `in:"path" v:"required" in:"path" dc:"菜单id"`
}
type DeleteRes struct{}

type GetAuthorizedListReq struct {
	g.Meta `path:"/menu/authorized-list" method:"get" tags:"菜单管理" summary:"获取当前用户拥有的菜单列表" deprecated:"true"`
}
type GetAuthorizedListRes []*api.MenuTree
