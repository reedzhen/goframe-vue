package model

// DefaultTree 默认树字段
type DefaultTree struct {
	Id       int64  // 主键Id
	ParentId int64  // 父级Id
	Level    int    // 树层级
	Tree     string // 树
}
