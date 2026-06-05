package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/library/query"
	"strings"
)

// LoginRecordPageInput 用户分页
type LoginRecordPageInput struct {
	query.PageParam
	Username  string
	CreatedAt []string
}

func (q *LoginRecordPageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.Username != "" {
		m = m.WhereLike("username", "%"+strings.TrimSpace(q.Username)+"%")
	}
	if len(q.CreatedAt) == 2 && q.CreatedAt[0] != "" && q.CreatedAt[1] != "" {
		m = m.WhereBetween("created_at", q.CreatedAt[0], gtime.NewFromStr(q.CreatedAt[1]).EndOfDay().String())
	}
	return m
}

// LoginRecordCreateInput 登录日志新增
type LoginRecordCreateInput struct {
	Username  string // 用户账号
	LoginType uint   // 操作类型 1登录成功/2登录失败/3退出登录
	Remark    string // 备注
}
