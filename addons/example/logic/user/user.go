package system

import (
	"context"
	"goframe-vben/addons/example/service"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/model/entity"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// GetProfileById 获取个人信息
func (s *sUser) GetProfileById(ctx context.Context, userId int64) (out *entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).WherePri(userId).Scan(&out)
	return
}
