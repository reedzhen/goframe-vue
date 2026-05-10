package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goframe-vben/api/admin/common"
)

func (c *ControllerCommon) WxPayNotify(ctx context.Context, req *common.WxPayNotifyReq) (res *common.WxPayNotifyRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
