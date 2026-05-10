package common

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaGenerateReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"通用接口" summary:"获取验证码" dc:"注意直接返回的是base64" noAuth:"true"`
}
type CaptchaGenerateRes struct {
	Key    string `json:"key"`
	Base64 string `json:"base64"`
	Answer string `json:"-"`
}
