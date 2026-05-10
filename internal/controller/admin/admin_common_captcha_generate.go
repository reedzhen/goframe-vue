package admin

import (
	"context"
	"goframe-vben/internal/library/captcha"

	"goframe-vben/api/admin/common"
)

func (c *ControllerCommon) CaptchaGenerate(ctx context.Context, req *common.CaptchaGenerateReq) (res *common.CaptchaGenerateRes, err error) {
	id, b64s, answer, err := captcha.Generate(ctx)
	return &common.CaptchaGenerateRes{
		Key:    id,
		Base64: b64s,
		Answer: answer,
	}, err
}
