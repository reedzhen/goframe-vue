package dto

import (
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/query"
)

// ConfigModuleCreateInput 配置模块新增
type ConfigModuleCreateInput struct {
	Code        string
	Name        string
	Description string
	Sort        int
	Status      int
}

// ConfigModuleUpdateInput 配置模块编辑
type ConfigModuleUpdateInput struct {
	Id          int64
	Code        string
	Name        string
	Description string
	Sort        int
	Status      int
}

// ConfigModuleListInput 配置模块列表查询
type ConfigModuleListInput struct {
	Keywords string
	Status   *int
}

// ConfigItemPageInput 配置项分页查询
type ConfigItemPageInput struct {
	query.PageParam
	ModuleId int64
	Keywords string
	Status   *int
}

// Cond 组装配置项分页查询条件。
func (q *ConfigItemPageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.ModuleId > 0 {
		m = m.Where("module_id", q.ModuleId)
	}
	if q.Status != nil {
		m = m.Where("status", *q.Status)
	}
	if keywords := strings.TrimSpace(q.Keywords); keywords != "" {
		like := "%" + keywords + "%"
		m = m.Where("(name LIKE ? OR config_key LIKE ?)", like, like)
	}
	return m
}

// ConfigItemCreateInput 配置项新增
type ConfigItemCreateInput struct {
	ModuleId     int64
	Name         string
	ConfigKey    string
	ConfigValue  string
	DefaultValue string
	ValueType    string
	InputType    int
	InputParams  string
	Description  string
	Sort         int
	Status       int
	IsSystem     int
}

// ConfigItemUpdateInput 配置项编辑
type ConfigItemUpdateInput struct {
	Id           int64
	ModuleId     int64
	Name         string
	ConfigKey    string
	ConfigValue  string
	DefaultValue string
	ValueType    string
	InputType    int
	InputParams  string
	Description  string
	Sort         int
	Status       int
	IsSystem     int
}

// ConfigValueItemInput 配置值保存项
type ConfigValueItemInput struct {
	ConfigKey   string
	ConfigValue string
}

// ConfigSaveValuesInput 保存模块配置值
type ConfigSaveValuesInput struct {
	ModuleCode string
	Values     []ConfigValueItemInput
}

// ConfigGetValuesOutput 配置值返回
type ConfigGetValuesOutput struct {
	ModuleCode string         `json:"moduleCode"` // 模块编码
	Data       map[string]any `json:"data"`       // 配置键值
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

	SaobeiInstNo      string `json:"saobeiInstNo"`      //商户系统机构号inst_no
	SaobeiKey         string `json:"saobeiKey"`         // 商户系统令牌
	SaobeiTerminalId  string `json:"saobeiTerminalId"`  //支付系统：商户号终端号
	SaobeiAccessToken string `json:"saobeiAccessToken"` //支付系统： 令牌

	// 系统流程需要
	PayNo   string  `json:"payNo"`   //商户号
	PayId   uint64  `json:"payId"`   //支付ID
	FeeRate float64 `json:"feeRate"` //支付费率
}
