package captcha

import (
	"context"
	"github.com/mojocn/base64Captcha"
)

var (
	captchaStore = base64Captcha.DefaultMemStore
)

// Generate 创建验证码
func Generate(ctx context.Context) (id, b64s, answer string, err error) {
	captchaDriver := &base64Captcha.DriverString{
		Height:     80,
		Width:      240,
		NoiseCount: 0,
		//ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "1234567890",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	captcha := base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	id, b64s, answer, err = captcha.Generate()
	return
}

// VerifyAndClear 校验验证码，并清空缓存的验证码信息
func VerifyAndClear(ctx context.Context, id string, answer string) bool {
	return captchaStore.Verify(id, answer, true)
}
