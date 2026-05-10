package operationrecord

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/query"
)

type GetPageReq struct {
	g.Meta `path:"/operation-record/page" method:"get" tags:"操作日志" summary:"获取操作日志分页"`
	query.PageParam
	Username  string   `json:"username" in:"query" dc:"用户账号"`
	Module    string   `json:"module" in:"query" dc:"操作模块"`
	CreatedAt []string `json:"createdAt" in:"query" dc:"操作时间"`
}
type GetPageRes query.Result

type DeleteReq struct {
	g.Meta `path:"/operation-record/delete" method:"post" tags:"操作日志" summary:"删除超过n天的日志"`
	Day    int `json:"day" d:"7" dc:"超过n天的日志"`
}
type DeleteRes struct{}
