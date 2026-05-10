package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/library/gvorm"
	"goframe-vben/internal/library/gvorm/handler"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/library/tree"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"goframe-vben/utility/password"
	"strings"
)

type sUser struct {
}

func init() {
	service.RegisterUser(NewUser())
}

func NewUser() *sUser {
	return &sUser{}
}

// ValidateExists 获取个人信息
func (s *sUser) ValidateExists(ctx context.Context, id int64) (out *entity.SysUser, err error) {
	if err = dao.SysUser.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return nil, err
	}
	if out == nil {
		return nil, gerror.NewCode(codes.CodeUserNotFound)
	}
	return
}

// GetInfoByUsername 获取个人信息
func (s *sUser) GetInfoByUsername(ctx context.Context, username string) (out *entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where("username", strings.TrimSpace(username)).Scan(&out)
	return
}

// Page 用户分页
func (s *sUser) Page(ctx context.Context, in dto.UserPageInput) (out *query.Result, err error) {
	list := make([]*dto.UserPageOutput, 0)
	// .Handler(handler.FilterSubRoleUser(contexts.GetUserId(ctx)), handler.FilterUserScope())
	out, err = query.Page(dao.SysUser.Ctx(ctx).WithAll(), &in, &list)
	return
}

// Create 新建用户
func (s *sUser) Create(ctx context.Context, in dto.UserCreateInput) (out int64, err error) {
	var (
		db   = dao.SysUser.Ctx(ctx)
		curr = contexts.GetUser(ctx)
	)

	if err = gvorm.CheckFieldExist(dao.SysUser.Ctx(ctx), dao.SysUser.Columns().Username, in.Username); err != nil {
		return 0, gerror.NewCode(codes.CodeUserDuplicate)
	}

	// 验证传入角色Id在不在当前登录用户的下级角色里(超管不校验,父级id=0即新建贴牌时不校验)
	if in.ParentId > 0 && !curr.IsAdmin {
		if err = service.Role().VerifyRoleId(ctx, in.RoleId); err != nil {
			return 0, err
		}
	}

	var user entity.SysUser
	if err = gconv.Struct(in, &user); err != nil {
		return
	}
	user.Password, user.Salt = password.Generate(in.Password)
	user.Status = consts.UserStatusOk

	// 生成树
	if user.Level, user.Tree, err = tree.GenerateChild(db, in.CreatedBy); err != nil {
		return
	}

	// 新建用户
	out, err = db.Data(user).InsertAndGetId()

	return
}

// Update 修改用户
func (s *sUser) Update(ctx context.Context, in dto.UserUpdateInput) (err error) {
	user, err := s.ValidateExists(ctx, in.Id)
	if err != nil {
		return
	}

	if err = gvorm.CheckFieldExist(dao.SysUser.Ctx(ctx), "username", in.Username, in.Id); err != nil {
		return
	}

	// 编辑用户
	if _, err = dao.SysUser.Ctx(ctx).WherePri(in.Id).FieldsEx("id").Data(gconv.Map(in)).Update(); err != nil {
		return
	}

	// 如果传入的用户的角色发生变化，那么此用户需要重新登录
	// 因为token 里保存了 role_id，这样做的好处是我可以从token里直接拿到角色id从而判断当前用户是否是超级管理员，原本判断是否是超级管理员每次都要查库
	if user.RoleId != in.RoleId {

	}
	return
}

// Delete 删除用户
func (s *sUser) Delete(ctx context.Context, id int64) (err error) {
	user, err := s.ValidateExists(ctx, id)
	if err != nil {
		return
	}
	if user.IsAdmin == 1 {
		return gerror.NewCode(codes.CodeUserDeleteDeny)
	}

	if contexts.GetUserId(ctx) == id {
		return gerror.NewCode(codes.CodeUserDeleteSelf)
	}

	// 删除用户
	if _, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).Delete(); err != nil {
		return err
	}

	return
}

// DeleteByTenantId 通过贴牌Id删除用户
func (s *sUser) DeleteByTenantId(ctx context.Context, tenantId int64) (err error) {
	if _, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().TenantId, tenantId).Delete(); err != nil {
		return err
	}

	return
}

// UpdatePwd 修改密码
func (s *sUser) UpdatePwd(ctx context.Context, oldPwd, newPwd string) (err error) {
	uid := contexts.GetUserId(ctx)

	user, err := s.ValidateExists(ctx, uid)
	if err != nil {
		return
	}

	if !password.Verify(user.Password, oldPwd, user.Salt) {
		return gerror.NewCode(codes.CodeUserOldPwdError)
	}

	data := g.Map{}
	data["password"], data["salt"] = password.Generate(newPwd)
	data["updated_by"] = uid
	_, err = dao.SysUser.Ctx(ctx).WherePri(user.Id).Data(data).Update()
	return
}

// ResetPwd 重置密码 123456
func (s *sUser) ResetPwd(ctx context.Context, userId int64, newPwd string) (err error) {
	user, err := s.ValidateExists(ctx, userId)
	if err != nil {
		return
	}
	if user.IsAdmin == consts.UserIsAdminYes {
		return gerror.NewCode(codes.CodeUserResetPwdDeny)
	}

	data := g.Map{}
	data["password"], data["salt"] = password.Generate(newPwd)
	data["updated_by"] = contexts.GetUserId(ctx)
	_, err = dao.SysUser.Ctx(ctx).WherePri(userId).Data(data).Update()
	return
}

// ChangeStatus 修改用户状态
func (s *sUser) ChangeStatus(ctx context.Context, userId int64, status int) (err error) {
	user, err := s.ValidateExists(ctx, userId)
	if err != nil {
		return
	}
	if user.IsAdmin == consts.UserIsAdminYes {
		return gerror.NewCode(codes.CodeUserStatusDeny)
	}

	if contexts.GetUserId(ctx) == userId {
		return gerror.NewCode(codes.CodeUserStatusSelf)
	}

	data := g.Map{}
	data["status"] = status
	data["updated_by"] = contexts.GetUserId(ctx)
	_, err = dao.SysUser.Ctx(ctx).WherePri(userId).Data(data).Update()
	return
}

// ChangeLastLoginAt 登录成功记录登录时间
func (s *sUser) ChangeLastLoginAt(ctx context.Context, userId int64) (err error) {
	_, err = dao.SysUser.Ctx(ctx).Data(do.SysUser{LastLoginAt: gtime.Now()}).WherePri(userId).Update()
	return
}

// GetCountByRoleId 获取数量
func (s *sUser) GetCountByRoleId(ctx context.Context, roleId int64) int {
	cnt, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().RoleId, roleId).Count()
	if err != nil {
		return 0
	}
	return cnt
}

// GetUserIdsByDeptIds 获取用户Ids
func (s *sUser) GetUserIdsByDeptIds(ctx context.Context, deptIds []int64) ([]int64, error) {
	v, err := dao.SysUser.Ctx(ctx).WhereIn("dept_id", deptIds).Fields("id").Array()
	if err != nil {
		return nil, err
	}
	return gconv.Int64s(v), nil
}

// GetList 获取用户列表
func (s *sUser) GetList(ctx context.Context, in dto.UserGetListInput) ([]*entity.SysUser, error) {
	var (
		db  = dao.SysUser.Ctx(ctx)
		col = dao.SysUser.Columns()
		out = make([]*entity.SysUser, 0)
	)
	if in.OrganizationId > 0 {
		db = db.Where(col.OrganizationId, in.OrganizationId)
	}
	if in.Nickname != "" {
		db = db.Where(col.Nickname, in.Nickname)
	}
	if in.Username != "" {
		db = db.Where(col.Username, in.Username)
	}
	if err := db.Handler(handler.FilterUserScope()).Scan(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetRoleIdByUserId 获取角色Id
func (s *sUser) GetRoleIdByUserId(ctx context.Context, userId int64) (int64, error) {
	v, err := dao.SysUser.Ctx(ctx).WherePri(userId).Fields("role_id").Value()
	if err != nil {
		return 0, err
	}
	return v.Int64(), nil
}

// CheckFieldExist 检测给定的字段是否唯一
func (s *sUser) CheckFieldExist(ctx context.Context, field string, value string, id ...int64) (err error) {
	return gvorm.CheckFieldExist(dao.SysUser.Ctx(ctx), field, value, id...)
}
