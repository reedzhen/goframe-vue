// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"goframe-vben/addons/example/api/user/admin"
)

type IUserAdmin interface {
	Hello(ctx context.Context, req *admin.HelloReq) (res *admin.HelloRes, err error)
	Task(ctx context.Context, req *admin.TaskReq) (res *admin.TaskRes, err error)
}
