package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ConfigUpdateInput 系统配置编辑
type ConfigUpdateInput struct {
	Group string
	Data  g.MapStrAny
}

// ConfigGetListOutput 系统配置列表返回
type ConfigGetListOutput struct {
	Group string      `json:"group"` // 配置分组
	Data  g.MapStrAny `json:"data"`  // 配置数据
}

type ConfigUploadOutput struct {
	UploadDrive              string `json:"uploadDrive"` // local oss
	UploadOssEndpoint        string `json:"uploadOssEndpoint"`
	UploadOssAccessKeyId     string `json:"uploadOssAccessKeyId"`
	UploadOssAccessKeySecret string `json:"uploadOssAccessKeySecret"`
	UploadOssBucket          string `json:"uploadOssBucket"`
	UploadOssBucketUrl       string `json:"uploadOssBucketUrl"`
}

type ConfigPayOutput struct {
	// 支付方式 wechat微信支付/saobei扫呗
	PayMethod string `json:"payMethod"`
	// 微信支付
	WxPayAppId      string `json:"payWxPayAppId"`
	WxPayMchId      string `json:"payWxPayMchId"`
	WxPaySerialNo   string `json:"payWxPaySerialNo"`
	WxPayAPIv3Key   string `json:"payWxPayAPIv3Key"`
	WxPayPrivateKey string `json:"payWxPayPrivateKey"`
	WxPayJsApiUrl   string `json:"payWxPayJsApiUrl"`
	// 扫呗支付
	SaobeiMerchantNo  string `json:"paySaobeiMerchantNo"`  // 扫呗支付商户号
	SaobeiTerminalId  string `json:"paySaobeiTerminalId"`  // 扫呗支付终端号
	SaobeiAccessToken string `json:"paySaobeiAccessToken"` // 扫呗支付秘钥
}

type ConfigBasicOutput struct {
	BasicTenantName string `json:"tenant_name"`
	BasicTenantLogo string `json:"tenant_logo"`
	BasicIcpCode    string `json:"icp_code"`
	BasicCopyright  string `json:"copyright"`
	BasicAppid      string `json:"appid"`
	BasicHomeCover1 string `json:"home_cover1"`
	BasicHomeCover2 string `json:"home_cover2"`
	BasicHomeCover3 string `json:"home_cover3"`
}

// ConfigUploadLogoInput 上传logo
type ConfigUploadLogoInput struct {
	File *ghttp.UploadFile // 上传文件对象
	Dir  string            `json:"dir"` // 上传文件保存的目录
}

// PayConfig 支付相关配置
type PayConfig struct {
	Debug           bool   `json:"debug"`
	PaySource       string `json:"paySource"`          // 微信 扫呗
	WXSpMchId       string `json:"spMchId"`            //微信服务商ID
	WXSpAppId       string `json:"spAppId"`            //微信服务商APPID
	WXPaySerial     string `json:"paySerial"`          //微信服务商证书序列号
	WXPayApiKey     string `json:"payApiKey"`          //微信支付APIkey（v3）
	WxPayPrivateKey string `json:"payWxPayPrivateKey"` //微信支付私钥KEY
	WXSubMchId      string `json:"subMchId"`           //子商户ID
	WXSubAppId      string `json:"subAppId"`           //子商户APPID
	// 微信支付
	//SaobeiMerchantNo  string `json:"saobeiMerchantNo"`  //扫呗商户号
	SaobeiInstNo      string `json:"saobeiInstNo"`      //商户系统机构号inst_no
	SaobeiKey         string `json:"saobeiKey"`         // 商户系统令牌
	SaobeiTerminalId  string `json:"saobeiTerminalId"`  //支付系统：商户号终端号
	SaobeiAccessToken string `json:"saobeiAccessToken"` //支付系统： 令牌

	// 系统流程需要
	PayNo   string  `json:"payNo"`   //商户号
	PayId   uint64  `json:"payId"`   //支付ID
	FeeRate float64 `json:"feeRate"` //支付费率
}
