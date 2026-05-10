package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goframe-vben/api/admin/common"
)

func (c *ControllerCommon) SaobeiPayNotify(ctx context.Context, req *common.SaobeiPayNotifyReq) (res *common.SaobeiPayNotifyRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
