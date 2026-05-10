package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/api"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/library/tree"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"goframe-vben/utility/tools"
	"strings"
)

type sRole struct{}

func init() {
	service.RegisterRole(NewRole())
}

func NewRole() *sRole {
	return &sRole{}
}

// Page 角色分页
func (s *sRole) Page(ctx context.Context, in dto.RolePageInput) (out *query.Result, err error) {
	var list []*dto.RolePageOutput
	out, err = query.Page(dao.SysRole.Ctx(ctx), &in, &list)
	if err != nil || out == nil {
		return
	}

	// 外挂菜单Id
	roleIds := gconv.Int64s(gdb.ListItemValues(list, "Id"))
	rmList, err := s.GetRoleMenuListByCond(ctx, roleIds)
	if err != nil {
		return
	}
	// 使用role作为键，menu数组作为值
	menuMap := tools.SliceToMapsWithFieldFunc(rmList, func(u *entity.SysRoleMenu) int64 { return u.RoleId }, func(u *entity.SysRoleMenu) int64 { return u.MenuId })
	for _, v := range list {
		vv, ok := menuMap[v.Id]
		if !ok {
			continue
		}
		v.Permissions = vv
	}

	return
}

// Create 新建角色
func (s *sRole) Create(ctx context.Context, in dto.RoleCreateInput) (out int64, err error) {

	//// 校验
	//if in.Code != "" {
	//	if err = clnorm.CheckFieldExist(db, "code", in.Code); err != nil {
	//		return
	//	}
	//}

	// 生成树
	//if in.Level, in.Tree, err = tree.GenerateChild(db, in.ParentId); err != nil {
	//	return
	//}

	//if err = dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	//	// 新增角色
	//	lastInsertId, err := dao.SysRole.Ctx(ctx).Data(in).InsertAndGetId()
	//	if err != nil {
	//		return err
	//	}
	//
	//	// 保存角色和menu的关系
	//	l := make([]*entity.SysRoleMenu, 0)
	//	for _, v := range in.Permissions {
	//		l = append(l, &entity.SysRoleMenu{
	//			RoleId: lastInsertId,
	//			MenuId: v,
	//			//TenantId: in.TenantId,
	//		})
	//	}
	//	if len(l) > 0 {
	//		if _, err := dao.SysRoleMenu.Ctx(ctx).Data(l).Insert(); err != nil {
	//			return err
	//		}
	//	}
	//
	//	out = lastInsertId
	//
	//	return nil
	//}); err != nil {
	//	return
	//}

	return dao.SysRole.Ctx(ctx).Data(in).InsertAndGetId()
}

// Update 编辑角色
func (s *sRole) Update(ctx context.Context, in dto.RoleUpdateInput) (err error) {
	//if in.ParentId == in.Id {
	//	return gerror.New("上级角色不能是自己")
	//}

	//// 获取角色明细
	//role, err := s.ValidateExists(ctx, in.Id)
	//if err != nil {
	//	return err
	//}

	//// 获取当前部门的的所有子级
	//childIds, _ := tree.DescendantIds(dao.SysRole.Ctx(ctx), role.Id, role.Tree)
	//if tools.InSlice(gconv.Int64s(childIds), in.ParentId) {
	//	return gerror.New("不能将子级角色设置为上级角色")
	//}

	param := gconv.Map(in)

	//return dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	//	//if role.ParentId != in.ParentId {
	//	//	// 更新当前角色的tree和level
	//	//	level, treePath, err := tree.GenerateChild(dao.SysRole.Ctx(ctx), in.ParentId)
	//	//	if err != nil {
	//	//		return err
	//	//	}
	//	//	// 更新子级角色的tree和level
	//	//	if err := tree.RebuildDescendants(dao.SysRole.Ctx(ctx), role.Id, level, treePath); err != nil {
	//	//		return err
	//	//	}
	//	//
	//	//	param["level"] = level
	//	//	param["tree"] = tree
	//	//}
	//
	//	if _, err := dao.SysRole.Ctx(ctx).WherePri(in.Id).FieldsEx("id").Data(param).Update(); err != nil {
	//		return err
	//	}
	//
	//	// 保存角色和menu的关系
	//	if err := s.SaveRoleMenuList(ctx, dto.RoleSaveMenuInput{
	//		RoleId:  in.Id,
	//		MenuIds: in.Permissions,
	//	}); err != nil {
	//		return err
	//	}
	//
	//	return nil
	//})

	if _, err = dao.SysRole.Ctx(ctx).WherePri(in.Id).FieldsEx("id").Data(param).Update(); err != nil {
		return err
	}

	return
}

// Delete 删除角色
func (s *sRole) Delete(ctx context.Context, id int64) error {
	role, err := s.ValidateExists(ctx, id)
	if err != nil {
		return err
	}

	if cnt := s.GetChildCount(ctx, role.Id); cnt > 0 {
		return gerror.NewCode(codes.CodeRoleHasChild)
	}

	if cnt := service.User().GetCountByRoleId(ctx, id); cnt > 0 {
		return gerror.NewCode(codes.CodeRoleInUse)
	}

	err = dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除角色
		if _, err = dao.SysRole.Ctx(ctx).WherePri(id).Delete(); err != nil {
			return err
		}
		// 删除角色和menu的关系
		if _, err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, id).Delete(); err != nil {
			return err
		}
		return nil
	})

	return nil
}

// DeleteByTenantId 删除角色
func (s *sRole) DeleteByTenantId(ctx context.Context, tenantId int64) error {
	// 删除角色
	if _, err := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().TenantId, tenantId).Delete(); err != nil {
		return err
	}
	// 删除角色菜单绑定
	if _, err := dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().TenantId, tenantId).Delete(); err != nil {
		return err
	}

	return nil
}

// GetChildCount 获取数量
func (s *sRole) GetChildCount(ctx context.Context, parentId int64) int {
	cnt, err := dao.SysRole.Ctx(ctx).Where("parent_id", parentId).Count()
	if err != nil {
		return 0
	}
	return cnt
}

// GetList 获取角色列表(默认获取当前用户角色的下属角色列表)
func (s *sRole) GetList(ctx context.Context, in dto.RoleGetListInput) ([]*entity.SysRole, error) {
	var (
		out = make([]*entity.SysRole, 0) // 返回值
		//curr = contexts.GetUser(ctx)
		db = dao.SysRole.Ctx(ctx)
	)

	// 非超管只能看到当前用户角色的下级角色
	//if curr.IsAdmin { // 超管
	//	role, err := s.ValidateExists(ctx, curr.RoleId)
	//	if err != nil {
	//		return nil, err
	//	}
	//	db = db.WhereLike("tree", tools.TrimLikeRight(tree.BuildChildPath(role.Tree, curr.RoleId)))
	//}

	if in.Name != "" {
		db = db.WhereLike("name", "%"+strings.TrimSpace(in.Name)+"%")
	}
	if err := db.Scan(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetInfo 获取详情
func (s *sRole) GetInfo(ctx context.Context, id int64) (out *entity.SysRole, err error) {
	if err = dao.SysRole.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	return
}

// ValidateExists 验证并获取详情
func (s *sRole) ValidateExists(ctx context.Context, id int64) (out *entity.SysRole, err error) {
	if err = dao.SysRole.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.NewCode(codes.CodeRoleNotFound)
	}
	return
}

// GetRoleMenuList 获取角色下所有菜单
func (s *sRole) GetRoleMenuList(ctx context.Context, roleId int64) ([]*api.MenuItem, error) {
	var (
		menus []*api.MenuItem // 可以展示的所有菜单
		//curr  = contexts.GetUser(ctx)
	)

	if err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Hide, consts.MenuIsHideNo).Order("sort asc").Scan(&menus); err != nil {
		return nil, err
	}

	//if curr.IsAdmin { // 超管
	//	if err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Hide,consts.MenuIsHideNo).Order("sort asc").Scan(&menus); err != nil {
	//		return nil, err
	//	}
	//} else {
	//	// 获取当前角色所有菜单id（我只能把我所拥有的菜单权限显示出来）
	//	menuIdsValue, err := dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, curr.RoleId).Fields("DISTINCT menu_id").Array()
	//	if err != nil {
	//		return nil, err
	//	}
	//	if err := dao.SysMenu.Ctx(ctx).WherePri(menuIdsValue).Where(dao.SysMenu.Columns().Hide,consts.MenuIsHideNo).Order("sort asc").Scan(&menus); err != nil {
	//		return nil, err
	//	}
	//}
	if len(menus) == 0 {
		return nil, nil
	}

	// 选中角色拥有的菜单即勾选中的菜单
	var roleMenus []*entity.SysRoleMenu
	if err := dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, roleId).Scan(&roleMenus); err != nil {
		return nil, err
	}
	for _, v := range menus {
		for _, v1 := range roleMenus {
			if v.Id == v1.MenuId {
				v.Checked = true
				break
			}
		}
	}

	return menus, nil
}

// SaveRoleMenuList 保存角色下所有菜单
func (s *sRole) SaveRoleMenuList(ctx context.Context, in dto.RoleSaveMenuInput) (err error) {

	l := make([]*entity.SysRoleMenu, 0)
	for _, v := range in.MenuIds {
		l = append(l, &entity.SysRoleMenu{
			RoleId: in.RoleId,
			MenuId: v,
		})
	}

	err = dao.SysRoleMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, in.RoleId).Delete(); err != nil {
			return err
		}

		if _, err = dao.SysRoleMenu.Ctx(ctx).Data(l).Insert(); err != nil {
			return err
		}
		return nil
	})

	return
}

// VerifyRoleId 验证传入角色Id在不在当前登录用户的下级角色里
func (s *sRole) VerifyRoleId(ctx context.Context, roleId int64) error {
	//subIds, _ := tree.DescendantIdsById(dao.SysRole.Ctx(ctx), contexts.GetRoleId(ctx))
	//
	//if !tools.InSlice(gconv.Int64s(subIds), roleId) {
	//	return gerror.New("角色Id是无效的")
	//}

	return nil
}

// GetRoleMenuListByCond 获取角色下所有菜单
func (s *sRole) GetRoleMenuListByCond(ctx context.Context, roleIds []int64) (out []*entity.SysRoleMenu, err error) {
	err = dao.SysRoleMenu.Ctx(ctx).WhereIn("role_id", roleIds).Scan(&out)
	return
}

// ChangeDataScope 修改角色数据权限
func (s *sRole) ChangeDataScope(ctx context.Context, roleId int64, dataScope int) (err error) {
	//if roleId == consts.SuperRoleId {
	//	return gerror.New("此角色不允许修改数据权限")
	//}

	_, err = dao.SysRole.Ctx(ctx).WherePri(roleId).Data(do.SysRole{
		DataScope: dataScope,
	}).Update()
	return
}

// GetSubIds 获取下级角色的Id
func (s *sRole) GetSubIds(ctx context.Context, roleId int64) []int64 {
	if v, err := tree.DescendantIdsById(dao.SysRole.Ctx(ctx), roleId); err == nil && len(v) > 0 {
		return gconv.Int64s(v)
	}
	return nil
}
