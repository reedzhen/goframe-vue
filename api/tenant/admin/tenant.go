package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/entity"
)

type GetPageReq struct {
	g.Meta `path:"/tenant/page" method:"get" tags:"租户管理" summary:"获取租户分页" deprecated:"true"`
	query.PageParam
	TenantName string `json:"tenant_name" in:"query" dc:"租户名称"`
	LinkMan    string `json:"link_man" in:"query" dc:"联系人姓名"`
	LinkPhone  string `json:"link_phone" in:"query" dc:"联系人手机"`
	Status     int    `json:"status" in:"query" dc:"租户状态 1正常/2停用"`
}
type GetPageRes struct {
	*query.Result
}

type GetListReq struct {
	g.Meta     `path:"/tenant/list" method:"get" tags:"租户管理" summary:"获取租户列表" deprecated:"true"`
	TenantName string `json:"tenant_name" in:"query" dc:"租户名称"`
	LinkMan    string `json:"link_man" in:"query" dc:"联系人姓名"`
	LinkPhone  string `json:"link_phone" in:"query" dc:"联系人手机"`
}
type GetListRes struct {
	List []*entity.SysTenant `json:"list"`
}

type CreateReq struct {
	g.Meta `path:"/tenant/create" method:"post" tags:"租户管理" summary:"新建租户" deprecated:"true"`
	CreateUpdateBase
	PackageId int64  `json:"package_id" v:"required" dc:"租户套餐编号"`
	LinkPhone string `json:"link_phone" v:"required" dc:"联系人手机"`
}
type CreateRes struct{}

type UpdateReq struct {
	g.Meta `path:"/tenant/update" method:"post" tags:"租户管理" summary:"修改租户" deprecated:"true"`
	Id     int64 `json:"id" v:"required" dc:"主键ID"`
	CreateUpdateBase
}
type UpdateRes struct{}

type CreateUpdateBase struct {
	TenantName   string              `json:"tenant_name" v:"required" dc:"租户名称"`
	LinkMan      string              `json:"link_man" v:"required" dc:"联系人姓名"`
	Status       consts.TenantStatus `json:"status" v:"enums" dc:"贴牌状态 1正常/2停用"`
	Website      string              `json:"website" v:"required" dc:"绑定域名"`
	ExpireAt     *gtime.Time         `json:"expire_at" v:"required" dc:"过期时间"`
	AccountCount int                 `json:"account_count" v:"required" dc:"账号额度"`
}

type DeleteReq struct {
	g.Meta `path:"/tenant/delete/{Id}" method:"post" tags:"租户管理" summary:"删除租户" deprecated:"true"`
	Id     int64 `in:"path" v:"required" dc:"主键ID"`
}
type DeleteRes struct{}

type GetInfoReq struct {
	g.Meta `path:"/tenant/info/{Id}" method:"get" tags:"租户管理" summary:"获取租户详情" deprecated:"true"`
	Id     int64 `in:"path" v:"required" dc:"主键ID"`
}
type GetInfoRes struct {
	Data *entity.SysTenant `json:"data"`
}
