package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goframe-vben/api/admin/cron"
)

func (c *ControllerCron) Page(ctx context.Context, req *cron.PageReq) (res *cron.PageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
