package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type TenantGetInfoByHostReq struct {
	g.Meta `path:"/tenant/info/by-host" method:"get" tags:"通用接口" summary:"通过host获取租户详情" noAuth:"true"`
}
type TenantGetInfoByHostRes struct {
	Id         int64       `json:"id"            orm:"id"            ` //
	TenantName string      `json:"tenant_name"   orm:"tenant_name"   ` // 租户名称
	Website    string      `json:"website"       orm:"website"       ` // 绑定域名
	ExpireAt   *gtime.Time `json:"expire_at"     orm:"expire_at"     ` // 过期时间
}
