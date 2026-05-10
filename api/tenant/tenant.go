// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tenant

import (
	"context"

	"goframe-vben/api/tenant/admin"
)

type ITenantAdmin interface {
	GetPage(ctx context.Context, req *admin.GetPageReq) (res *admin.GetPageRes, err error)
	GetList(ctx context.Context, req *admin.GetListReq) (res *admin.GetListRes, err error)
	Create(ctx context.Context, req *admin.CreateReq) (res *admin.CreateRes, err error)
	Update(ctx context.Context, req *admin.UpdateReq) (res *admin.UpdateRes, err error)
	Delete(ctx context.Context, req *admin.DeleteReq) (res *admin.DeleteRes, err error)
	GetInfo(ctx context.Context, req *admin.GetInfoReq) (res *admin.GetInfoRes, err error)
}
