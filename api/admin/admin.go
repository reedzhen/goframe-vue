// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"goframe-vben/api/admin/auth"
	"goframe-vben/api/admin/common"
	"goframe-vben/api/admin/config"
	"goframe-vben/api/admin/cron"
	"goframe-vben/api/admin/dict"
	"goframe-vben/api/admin/loginrecord"
	"goframe-vben/api/admin/menu"
	"goframe-vben/api/admin/operationrecord"
	"goframe-vben/api/admin/organization"
	"goframe-vben/api/admin/role"
	"goframe-vben/api/admin/user"
)

type IAdminAuth interface {
	Login(ctx context.Context, req *auth.LoginReq) (res *auth.LoginRes, err error)
	Logout(ctx context.Context, req *auth.LogoutReq) (res *auth.LogoutRes, err error)
	GetInfo(ctx context.Context, req *auth.GetInfoReq) (res *auth.GetInfoRes, err error)
}

type IAdminCommon interface {
	CaptchaGenerate(ctx context.Context, req *common.CaptchaGenerateReq) (res *common.CaptchaGenerateRes, err error)
	SaobeiPayNotify(ctx context.Context, req *common.SaobeiPayNotifyReq) (res *common.SaobeiPayNotifyRes, err error)
	WxPayNotify(ctx context.Context, req *common.WxPayNotifyReq) (res *common.WxPayNotifyRes, err error)
	GetRoutesWithMeta(ctx context.Context, req *common.GetRoutesWithMetaReq) (res *common.GetRoutesWithMetaRes, err error)
	TenantGetInfoByHost(ctx context.Context, req *common.TenantGetInfoByHostReq) (res *common.TenantGetInfoByHostRes, err error)
	FileUpload(ctx context.Context, req *common.FileUploadReq) (res *common.FileUploadRes, err error)
	CallbackDemo(ctx context.Context, req *common.CallbackDemoReq) (res *common.CallbackDemoRes, err error)
}

type IAdminConfig interface {
	ModuleList(ctx context.Context, req *config.ModuleListReq) (res *config.ModuleListRes, err error)
	ModuleCreate(ctx context.Context, req *config.ModuleCreateReq) (res *config.ModuleCreateRes, err error)
	ModuleUpdate(ctx context.Context, req *config.ModuleUpdateReq) (res *config.ModuleUpdateRes, err error)
	ModuleDelete(ctx context.Context, req *config.ModuleDeleteReq) (res *config.ModuleDeleteRes, err error)
	ItemPage(ctx context.Context, req *config.ItemPageReq) (res *config.ItemPageRes, err error)
	ItemCreate(ctx context.Context, req *config.ItemCreateReq) (res *config.ItemCreateRes, err error)
	ItemUpdate(ctx context.Context, req *config.ItemUpdateReq) (res *config.ItemUpdateRes, err error)
	ItemDelete(ctx context.Context, req *config.ItemDeleteReq) (res *config.ItemDeleteRes, err error)
	ValueGet(ctx context.Context, req *config.ValueGetReq) (res *config.ValueGetRes, err error)
	ValueSave(ctx context.Context, req *config.ValueSaveReq) (res *config.ValueSaveRes, err error)
}

type IAdminCron interface {
	Page(ctx context.Context, req *cron.PageReq) (res *cron.PageRes, err error)
	Create(ctx context.Context, req *cron.CreateReq) (res *cron.CreateRes, err error)
	Update(ctx context.Context, req *cron.UpdateReq) (res *cron.UpdateRes, err error)
	Delete(ctx context.Context, req *cron.DeleteReq) (res *cron.DeleteRes, err error)
	ChangeStatus(ctx context.Context, req *cron.ChangeStatusReq) (res *cron.ChangeStatusRes, err error)
	ExecuteOnce(ctx context.Context, req *cron.ExecuteOnceReq) (res *cron.ExecuteOnceRes, err error)
	RecordPage(ctx context.Context, req *cron.RecordPageReq) (res *cron.RecordPageRes, err error)
}

type IAdminDict interface {
	GetPage(ctx context.Context, req *dict.GetPageReq) (res *dict.GetPageRes, err error)
	Create(ctx context.Context, req *dict.CreateReq) (res *dict.CreateRes, err error)
	Update(ctx context.Context, req *dict.UpdateReq) (res *dict.UpdateRes, err error)
	Delete(ctx context.Context, req *dict.DeleteReq) (res *dict.DeleteRes, err error)
	GetList(ctx context.Context, req *dict.GetListReq) (res *dict.GetListRes, err error)
	DataGetPage(ctx context.Context, req *dict.DataGetPageReq) (res *dict.DataGetPageRes, err error)
	DataCreate(ctx context.Context, req *dict.DataCreateReq) (res *dict.DataCreateRes, err error)
	DataUpdate(ctx context.Context, req *dict.DataUpdateReq) (res *dict.DataUpdateRes, err error)
	DataDelete(ctx context.Context, req *dict.DataDeleteReq) (res *dict.DataDeleteRes, err error)
	DataGetList(ctx context.Context, req *dict.DataGetListReq) (res *dict.DataGetListRes, err error)
}

type IAdminLoginrecord interface {
	GetPage(ctx context.Context, req *loginrecord.GetPageReq) (res *loginrecord.GetPageRes, err error)
}

type IAdminMenu interface {
	GetList(ctx context.Context, req *menu.GetListReq) (res *menu.GetListRes, err error)
	Create(ctx context.Context, req *menu.CreateReq) (res *menu.CreateRes, err error)
	Update(ctx context.Context, req *menu.UpdateReq) (res *menu.UpdateRes, err error)
	Delete(ctx context.Context, req *menu.DeleteReq) (res *menu.DeleteRes, err error)
	GetAuthorizedList(ctx context.Context, req *menu.GetAuthorizedListReq) (res *menu.GetAuthorizedListRes, err error)
}

type IAdminOperationrecord interface {
	GetPage(ctx context.Context, req *operationrecord.GetPageReq) (res *operationrecord.GetPageRes, err error)
	Delete(ctx context.Context, req *operationrecord.DeleteReq) (res *operationrecord.DeleteRes, err error)
}

type IAdminOrganization interface {
	GetList(ctx context.Context, req *organization.GetListReq) (res *organization.GetListRes, err error)
	Create(ctx context.Context, req *organization.CreateReq) (res *organization.CreateRes, err error)
	Update(ctx context.Context, req *organization.UpdateReq) (res *organization.UpdateRes, err error)
	Delete(ctx context.Context, req *organization.DeleteReq) (res *organization.DeleteRes, err error)
	GetInfo(ctx context.Context, req *organization.GetInfoReq) (res *organization.GetInfoRes, err error)
}

type IAdminRole interface {
	GetPage(ctx context.Context, req *role.GetPageReq) (res *role.GetPageRes, err error)
	Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error)
	Update(ctx context.Context, req *role.UpdateReq) (res *role.UpdateRes, err error)
	Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error)
	GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error)
	GetMenuList(ctx context.Context, req *role.GetMenuListReq) (res *role.GetMenuListRes, err error)
	SaveMenuList(ctx context.Context, req *role.SaveMenuListReq) (res *role.SaveMenuListRes, err error)
	ChangeDataScope(ctx context.Context, req *role.ChangeDataScopeReq) (res *role.ChangeDataScopeRes, err error)
}

type IAdminUser interface {
	GetPage(ctx context.Context, req *user.GetPageReq) (res *user.GetPageRes, err error)
	GetInfo(ctx context.Context, req *user.GetInfoReq) (res *user.GetInfoRes, err error)
	Create(ctx context.Context, req *user.CreateReq) (res *user.CreateRes, err error)
	Update(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error)
	Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error)
	CheckFieldExist(ctx context.Context, req *user.CheckFieldExistReq) (res *user.CheckFieldExistRes, err error)
	UpdatePwd(ctx context.Context, req *user.UpdatePwdReq) (res *user.UpdatePwdRes, err error)
	ResetPwd(ctx context.Context, req *user.ResetPwdReq) (res *user.ResetPwdRes, err error)
	ChangeStatus(ctx context.Context, req *user.ChangeStatusReq) (res *user.ChangeStatusRes, err error)
	GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error)
}
