package gvorm

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// CheckFieldExist 检测给定的字段是否唯一
func CheckFieldExist(db *gdb.Model, field string, value string, id ...int64) (err error) {
	if field == "" || value == "" {
		return gerror.New("字段不能为空")
	}

	db = db.Where(field, value)
	if len(id) > 0 && id[0] > 0 {
		// 不包含自己
		db = db.WhereNot("id", id[0])
	}
	var n int
	if n, err = db.Count(); err != nil {
		return
	}
	if n > 0 {
		return gerror.Newf("%s已被占用", value)
	}
	return
}
