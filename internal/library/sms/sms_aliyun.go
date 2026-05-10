// Package sms
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sms

import (
	"context"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type AliYunDrive struct {
	AppKey    string
	AppSecret string
	SignName  string
}

func NewAliYunSms(appKey string, appSecret string, signName string, templateCode string) *AliYunDrive {
	return &AliYunDrive{AppKey: appKey, AppSecret: appSecret, SignName: signName}
}

type SendSmsResponse struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}

// Send 调用阿里云短信发送接口
func (s *AliYunDrive) Send(ctx context.Context, in SendCode) (err error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", s.AppKey, s.AppSecret)
	if err != nil {
		return
	}

	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = in.Phone
	req.SignName = s.SignName
	req.TemplateCode = in.TemplateCode
	req.TemplateParam = "{\"code\":" + in.Code + "}"

	res, err := client.SendSms(req)
	if err != nil {
		return err
	}

	var msg SendSmsResponse
	_ = gjson.Unmarshal(res.GetHttpContentBytes(), &msg)
	if msg.Code != "OK" {
		return errors.New(msg.Message)
	}
	return
}
