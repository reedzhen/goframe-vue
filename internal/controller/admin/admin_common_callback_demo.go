package admin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-vben/api/admin/common"
)

type NotifyRsp struct {
	Code    int    `json:"code" dc:"状态码 0成功"`
	Message string `json:"message" dc:"提示消息"`
}

func (c *ControllerCommon) CallbackDemo(ctx context.Context, req *common.CallbackDemoReq) (res *common.CallbackDemoRes, err error) {
	r := g.RequestFromCtx(ctx)
	g.Dump(req.ClientId)
	g.Dump(req.ModuleCode)
	g.Dump(req.Type)
	r.Response.WriteJsonExit(&NotifyRsp{Code: 0, Message: "success"})
	return
}
