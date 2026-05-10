package common

import "github.com/gogf/gf/v2/frame/g"

type GetRoutesWithMetaReq struct {
	g.Meta `path:"/route/list" method:"get" tags:"通用接口" summary:"获取接口路由列表" dc:"配置接口权限时使用" noAuth:"true"`
}

type GetRoutesWithMetaRes []*RoutesWithMetaItem

type RoutesWithMetaItem struct {
	Tags    string `json:"tags" dc:"标签/分组"`
	Summary string `json:"summary" dc:"简要描述"`
	Path    string `json:"path" dc:"接口路径"`
	Method  string `json:"method" dc:"方法 get/post"`
}
