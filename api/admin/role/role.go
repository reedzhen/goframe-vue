package role

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/api"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/entity"
)

type GetPageReq struct {
	g.Meta `path:"/role/page" method:"get" tags:"角色管理" summary:"获取角色分页"`
	query.PageParam
	Name string `json:"name" in:"query"`
	Code string `json:"code" in:"query"`
}
type GetPageRes struct {
	*query.Result
}

type CreateReq struct {
	g.Meta `path:"/role/create" method:"post" tags:"角色管理" summary:"新增角色"`
	Name   string            `json:"name" v:"required" dc:"角色名称"`
	Status consts.RoleStatus `json:"status" v:"enums" dc:"角色状态 1正常/2停用"`
	Remark string            `json:"remark" dc:"备注"`
	Code   string            `json:"code" dc:"角色标识"`
}
type CreateRes struct{}

type UpdateReq struct {
	g.Meta `path:"/role/update" method:"post" tags:"角色管理" summary:"编辑角色"`
	Id     int64             `json:"id" v:"required" dc:"角色id"`
	Name   string            `json:"name" v:"required" dc:"角色名称"`
	Code   string            `json:"code" dc:"角色标识"`
	Remark string            `json:"remark" dc:"备注"`
	Status consts.RoleStatus `json:"status" v:"enums" dc:"角色状态 1正常/2停用"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/role/delete/{Id}" method:"post" tags:"角色管理" summary:"删除角色"`
	Id     int64 ` json:"id" in:"path" v:"required" dc:"角色id"`
}
type DeleteRes struct{}

type GetListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色管理" summary:"获取角色列表" noAuth:"true"`
	Name   string `json:"name" in:"query" dc:"角色名称"`
}
type GetListRes struct {
	List []*entity.SysRole `json:"list"`
}

type GetMenuListReq struct {
	g.Meta `path:"/role/menu-list" method:"get" tags:"角色管理" summary:"获取当前角色的菜单列表"`
	Id     int64 `json:"id" v:"required" in:"query" dc:"角色id"`
}
type GetMenuListRes []*api.MenuItem

type SaveMenuListReq struct {
	g.Meta  `path:"/role/menu-save" method:"post" tags:"角色管理" summary:"保存当前角色的菜单列表" `
	Id      int64   `json:"id" v:"required"  dc:"角色id"`
	MenuIds []int64 `json:"menuIds" v:"required" dc:"菜单id数组"`
}
type SaveMenuListRes struct{}

type ChangeDataScopeReq struct {
	g.Meta    `path:"/role/change-scope" method:"post" tags:"角色管理" summary:"修改数据权限"`
	Id        int64                `json:"id" v:"required"  dc:"角色id"`
	DataScope consts.RoleDataScope `json:"dataScope" v:"enums" dc:"角色状态 数据范围 1全部/2当前部门/3当前及以下部门/4自定义部门"`
}
type ChangeDataScopeRes struct{}
