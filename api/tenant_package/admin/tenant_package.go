package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/api"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/entity"
)

type GetPageReq struct {
	g.Meta `path:"/tenant/package/page" method:"get" tags:"租户套餐" summary:"获取租户套餐分页" deprecated:"true"`
	query.PageParam
	Name   string `json:"name" in:"query" dc:"套餐名"`
	Status int    `json:"status" in:"query" dc:"租户状态 1正常/2停用"`
}
type GetPageRes struct {
	*query.Result
}

type GetListReq struct {
	g.Meta `path:"/tenant/package/list" method:"get" tags:"租户套餐" summary:"获取租户套餐列表" deprecated:"true"`

	Name   string `json:"name" in:"query" dc:"套餐名"`
	Status int    `json:"status" in:"query" dc:"租户状态 1正常/2停用"`
}
type GetListRes struct {
	List []*entity.SysTenantPackage `json:"list"`
}

type CreateReq struct {
	g.Meta `path:"/tenant/package/create" method:"post" tags:"租户套餐" summary:"新建租户套餐" deprecated:"true"`
	CreateUpdateBase
}
type CreateRes struct{}

type UpdateReq struct {
	g.Meta `path:"/tenant/package/update" method:"post" tags:"租户套餐" summary:"修改租户套餐" deprecated:"true"`
	Id     int64 `json:"id" v:"required" dc:"主键ID"`
	CreateUpdateBase
}
type UpdateRes struct{}

type CreateUpdateBase struct {
	Name    string `json:"name" v:"required" dc:"套餐名"`
	Status  int    `json:"status" v:"required" dc:"租户状态 1正常/2停用"`
	MenuIds string `json:"menu_ids" v:"required" dc:"菜单ID 逗号分隔"`
	Remark  string `json:"remark" dc:"备注"`
}

type DeleteReq struct {
	g.Meta `path:"/tenant/package/delete/{Id}" method:"post" tags:"租户套餐" summary:"删除租户套餐" deprecated:"true"`
	Id     int64 `in:"path" v:"required" dc:"主键ID"`
}
type DeleteRes struct{}

type GetInfoReq struct {
	g.Meta `path:"/tenant/package/info/{Id}" method:"get" tags:"租户套餐" summary:"获取租户套餐详情" deprecated:"true"`
	Id     int64 `in:"path" v:"required" dc:"主键ID"`
}
type GetInfoRes struct {
	Data *entity.SysTenantPackage `json:"data"`
}

type GetMenuListReq struct {
	g.Meta `path:"/tenant/package/menu/list" method:"get" tags:"租户套餐" summary:"获取菜单列表" deprecated:"true"`
	Id     int64 `json:"id"  in:"query" dc:"套餐id"`
}
type GetMenuListRes []*api.MenuItem

type SyncTenantPermissionReq struct {
	g.Meta `path:"/tenant/package/permission/sync" method:"post" tags:"租户套餐" summary:"给使用此套餐的租户同步权限" deprecated:"true"`
	Id     int64 `json:"id" v:"required" dc:"主键ID"`
}
type SyncTenantPermissionRes struct{}
