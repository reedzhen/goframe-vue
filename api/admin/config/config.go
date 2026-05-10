package config

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta `path:"/config/update" method:"post" tags:"系统配置" summary:"配置编辑"`
	Group  string `json:"group" v:"required" dc:"配置分组"`
	Data   g.Map  `json:"data" v:"required" dc:"配置数据"`
}
type UpdateRes struct{}

type GetListByGroupReq struct {
	g.Meta `path:"/config/list" method:"get" tags:"系统配置" summary:"获取配置列表"`
	Group  string `json:"group" v:"required" in:"query" dc:"配置分组"`
}
type GetListByGroupRes struct {
	Group string      `json:"group" dc:"配置分组"`
	Data  g.MapStrAny `json:"data" dc:"配置数据"`
}
