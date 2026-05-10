package handler

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/library/gftenant"
)

// FilterUserScope 过滤当前用户数据权限
func FilterUserScope(field ...string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		curr := contexts.GetUser(ctx)

		// 超管拥有全部权限
		if curr == nil || curr.IsAdmin {
			return m
		}

		//用户拥有多个角色时，取所有角色中数据范围最宽的那个（90% 的业务场景适用），具体逻辑：
		//遍历用户的所有角色，提取每个角色的 data_scope 枚举值；
		//取枚举值最大的那个作为该用户最终的 data_scope；
		//若有 “自定义数据范围”，则优先合并自定义规则（比如多个角色的自定义部门列表取并集）。
		// https://www.doubao.com/chat/38417253802885634

		//// 获取角色信息
		//role, err := service.Role().ValidateExists(ctx, curr.RoleId)
		//if err != nil {
		//	panic(err)
		//}
		//
		//fieldStr := "id"
		//if len(field) > 0 && field[0] != "" {
		//	fieldStr = field[0]
		//}
		//
		//switch role.DataScope {
		//case int(consts.RoleScopeAll): // 全部权限
		//	// ...
		//case int(consts.RoleScopeCurrDept): // 当前部门
		//	userIds, err := service.User().GetUserIdsByDeptIds(ctx, []int64{curr.DeptId})
		//	if err != nil {
		//		panic(err)
		//	}
		//	m = m.WhereIn(fieldStr, userIds)
		//case int(consts.RoleScopeDeptWithSub): // 当前部门及以下部门
		//	deptIds := service.Dept().GetDeptAndSubIds(ctx, curr.DeptId)
		//	userIds, err := service.User().GetUserIdsByDeptIds(ctx, deptIds)
		//	if err != nil {
		//		panic(err)
		//	}
		//	m = m.WhereIn(fieldStr, userIds)
		//}
		//return m

		return nil
	}
}

// FilterSubRoleUser 非超管用户只能操作自己的下级角色用户
func FilterSubRoleUser(userId int64) gdb.ModelHandler {
	//return func(m *gdb.Model) *gdb.Model {
	//	ctx := m.GetCtx()
	//	curr := contexts.GetUser(ctx)
	//
	//	// 超管拥有全部权限
	//	if curr == nil || curr.RoleCode == consts.SuperRoleCode {
	//		return m
	//	}
	//
	//	var roleId int64
	//	if curr.UserId == userId {
	//		// 当前登录用户直接从上下文中取角色ID
	//		roleId = contexts.GetRoleId(ctx)
	//	} else {
	//		rId, err := service.User().GetRoleIdByUserId(ctx, userId)
	//		if err != nil {
	//			panic(err)
	//		}
	//		roleId = rId
	//	}
	//
	//	m = m.Where("id <> ?", userId)
	//	if roleIds := service.Role().GetSubIds(ctx, roleId); len(roleIds) > 0 {
	//		m = m.WhereIn("role_id", roleIds)
	//	}
	//	return m
	//}

	return nil
}

// FilterTenant 便捷设置当前贴牌
func FilterTenant(table ...string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if v := gftenant.GetTenant(m.GetCtx()); v > 0 {
			if len(table) > 0 {
				m = m.Where(fmt.Sprintf("%s.%s=?", table[0], gftenant.GetColumnName()), v)
			} else {
				m = m.Where(fmt.Sprintf("%s=?", gftenant.GetColumnName()), v)
			}
		}
		return m
	}
}
