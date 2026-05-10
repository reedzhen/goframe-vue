package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/model/entity"
)

type HelloReq struct {
	g.Meta `path:"/hello" method:"get" tags:"插件演示" summary:"测试接口获取数据"`
}
type HelloRes struct {
	*entity.SysUser `json:"user"`
}

type TaskReq struct {
	g.Meta `path:"/task" method:"get" tags:"插件演示" summary:"测试队列"`
}
type TaskRes struct {
}
