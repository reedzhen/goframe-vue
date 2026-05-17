package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/entity"
)

type ModuleListReq struct {
	g.Meta   `path:"/config/module/list" method:"get" tags:"配置中心" summary:"配置模块列表"`
	Keywords string `json:"keywords" in:"query" dc:"关键字"`
	Status   *int   `json:"status" in:"query" dc:"状态 1启用/2禁用"`
}
type ModuleListRes []*entity.SysConfigModule

type ModuleCreateReq struct {
	g.Meta      `path:"/config/module/create" method:"post" tags:"配置中心" summary:"新增配置模块"`
	Code        string `json:"code" v:"required" dc:"模块编码"`
	Name        string `json:"name" v:"required" dc:"模块名称"`
	Description string `json:"description" dc:"模块说明"`
	Sort        int    `json:"sort" dc:"排序值"`
	Status      int    `json:"status" dc:"状态 1启用/2禁用，默认启用"`
}
type ModuleCreateRes struct{}

type ModuleUpdateReq struct {
	g.Meta      `path:"/config/module/update" method:"post" tags:"配置中心" summary:"编辑配置模块"`
	Id          int64  `json:"id" v:"required" dc:"主键ID"`
	Code        string `json:"code" v:"required" dc:"模块编码"`
	Name        string `json:"name" v:"required" dc:"模块名称"`
	Description string `json:"description" dc:"模块说明"`
	Sort        int    `json:"sort" dc:"排序值"`
	Status      int    `json:"status" dc:"状态 1启用/2禁用，默认启用"`
}
type ModuleUpdateRes struct{}

type ModuleDeleteReq struct {
	g.Meta `path:"/config/module/delete/{Id}" method:"post" tags:"配置中心" summary:"删除配置模块"`
	Id     int64 `json:"id" v:"required" dc:"主键ID"`
}
type ModuleDeleteRes struct{}

type ItemPageReq struct {
	g.Meta   `path:"/config/item/page" method:"get" tags:"配置中心" summary:"配置项分页"`
	ModuleId int64  `json:"moduleId" in:"query" dc:"模块ID"`
	Keywords string `json:"keywords" in:"query" dc:"关键字"`
	Status   *int   `json:"status" in:"query" dc:"状态 1启用/2禁用"`
	query.PageParam
}
type ItemPageRes struct {
	*query.Result
}

type ItemCreateReq struct {
	g.Meta       `path:"/config/item/create" method:"post" tags:"配置中心" summary:"新增配置项"`
	ModuleId     int64  `json:"moduleId" v:"required" dc:"模块ID"`
	Name         string `json:"name" v:"required" dc:"配置项名称"`
	ConfigKey    string `json:"configKey" v:"required" dc:"配置键名"`
	ConfigValue  string `json:"configValue" dc:"配置值"`
	DefaultValue string `json:"defaultValue" dc:"默认值"`
	ValueType    string `json:"valueType" v:"required" dc:"值类型"`
	InputType    int    `json:"inputType" v:"required" dc:"输入类型"`
	InputParams  string `json:"inputParams" dc:"控件参数"`
	Description  string `json:"description" dc:"配置说明"`
	Sort         int    `json:"sort" dc:"排序值"`
	Status       int    `json:"status" dc:"状态 1启用/2禁用，默认启用"`
	IsSystem     int    `json:"isSystem" dc:"是否系统项 1是/2否，默认否"`
}
type ItemCreateRes struct{}

type ItemUpdateReq struct {
	g.Meta       `path:"/config/item/update" method:"post" tags:"配置中心" summary:"编辑配置项"`
	Id           int64  `json:"id" v:"required" dc:"主键ID"`
	ModuleId     int64  `json:"moduleId" v:"required" dc:"模块ID"`
	Name         string `json:"name" v:"required" dc:"配置项名称"`
	ConfigKey    string `json:"configKey" v:"required" dc:"配置键名"`
	ConfigValue  string `json:"configValue" dc:"配置值"`
	DefaultValue string `json:"defaultValue" dc:"默认值"`
	ValueType    string `json:"valueType" v:"required" dc:"值类型"`
	InputType    int    `json:"inputType" v:"required" dc:"输入类型"`
	InputParams  string `json:"inputParams" dc:"控件参数"`
	Description  string `json:"description" dc:"配置说明"`
	Sort         int    `json:"sort" dc:"排序值"`
	Status       int    `json:"status" dc:"状态 1启用/2禁用，默认启用"`
	IsSystem     int    `json:"isSystem" dc:"是否系统项 1是/2否，默认否"`
}
type ItemUpdateRes struct{}

type ItemDeleteReq struct {
	g.Meta `path:"/config/item/delete/{Id}" method:"post" tags:"配置中心" summary:"删除配置项"`
	Id     int64 `json:"id" v:"required" dc:"主键ID"`
}
type ItemDeleteRes struct{}

type ValueGetReq struct {
	g.Meta     `path:"/config/value/get" method:"get" tags:"配置中心" summary:"获取模块配置值"`
	ModuleCode string `json:"moduleCode" v:"required" in:"query" dc:"模块编码"`
}
type ValueGetRes struct {
	ModuleCode string         `json:"moduleCode" dc:"模块编码"`
	Data       map[string]any `json:"data" dc:"配置键值"`
}

type ValueSaveItem struct {
	ConfigKey   string `json:"configKey" v:"required" dc:"配置键名"`
	ConfigValue string `json:"configValue" dc:"配置值"`
}

type ValueSaveReq struct {
	g.Meta     `path:"/config/value/save" method:"post" tags:"配置中心" summary:"保存模块配置值"`
	ModuleCode string          `json:"moduleCode" v:"required" dc:"模块编码"`
	Values     []ValueSaveItem `json:"values" v:"required" dc:"配置值列表"`
}
type ValueSaveRes struct{}
