package sms

import (
	"context"
	"fmt"
)

const (
	SmsDriveAliYun  = "aliyun"  // 阿里云
	SmsDriveTencent = "tencent" // 腾讯云
)

type SendCode struct {
	Phone        string // 手机号
	Code         string // 验证码或短信内容
	TemplateCode string // 短信模板
}

// Drive 短信驱动
type Drive interface {
	Send(ctx context.Context, in SendCode) (err error)
}

func New(name ...string) Drive {
	var (
		instanceName = SmsDriveAliYun
		drive        Drive
	)

	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}

	switch instanceName {
	case SmsDriveAliYun:
		drive = &AliYunDrive{
			AppKey:    "",
			AppSecret: "",
			SignName:  "",
		}
	default:
		panic(fmt.Sprintf("暂不支持短信驱动:%v", instanceName))
	}
	return drive
}
