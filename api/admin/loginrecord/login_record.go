package loginrecord

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/query"
)

type GetPageReq struct {
	g.Meta `path:"/login-record/page" method:"get" tags:"登录日志" summary:"获取登录日志分页" `
	query.PageParam
	Username  string   `json:"username" in:"query" dc:"账号"`
	CreatedAt []string `json:"created_at" in:"query" dc:"登录时间"`
}
type GetPageRes struct {
	*query.Result
}
