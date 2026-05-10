package tree

import (
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/utility/tools"
)

const (
	idField       = "id"
	parentIdField = "parent_id"
	levelField    = "level"
	treeField     = "tree"
	nodePrefix    = "tr_"
	nodeSuffix    = " "
)

// Node 表示树节点的最小字段集合，对应表里的 id、parent_id、level、tree。
type Node struct {
	Id       int64
	ParentId int64
	Level    int
	Tree     string
}

// UpdateDO 是批量更新树字段时使用的 DO 对象，只更新 level 和 tree。
type UpdateDO struct {
	g.Meta `orm:"do:true"`
	Level  any
	Tree   any
}

// BuildChildPath 生成某个节点的下级 tree 路径前缀。
//
// 例如父级 tree 为 "tr_1 "、父级 id 为 2，则子级 tree 应为 "tr_1 tr_2 "。
func BuildChildPath(parentTree string, parentId int64) string {
	return fmt.Sprintf("%s%s%d%s", parentTree, nodePrefix, parentId, nodeSuffix)
}

// GenerateChild 根据 parentId 生成新节点应写入的 level 和 tree。
//
// parentId 小于等于 0 时表示顶级节点，返回 level=1、tree=""。
func GenerateChild(db *gdb.Model, parentId int64) (level int, treePath string, err error) {
	if parentId <= 0 {
		return 1, "", nil
	}

	parent, err := GetNode(db, parentId)
	if err != nil {
		return 0, "", err
	}
	if parent == nil {
		return 0, "", gerror.New("上级信息不存在")
	}

	return parent.Level + 1, BuildChildPath(parent.Tree, parentId), nil
}

// GetNode 查询指定节点的树字段信息。
//
// 这里只读取 id、parent_id、level、tree，避免业务表其它字段影响通用树操作。
func GetNode(db *gdb.Model, id int64) (*Node, error) {
	if id <= 0 {
		return nil, nil
	}

	record, err := db.WherePri(id).Fields(idField, parentIdField, levelField, treeField).One()
	if err != nil || record == nil {
		return nil, err
	}

	return recordToNode(record), nil
}

// DescendantIds 查询指定节点的所有下级节点 ID，不包含当前节点。
//
// 调用方已经有当前节点的 tree 时，使用这个方法可以少查一次当前节点。
func DescendantIds(db *gdb.Model, id int64, nodeTree string) ([]gdb.Value, error) {
	return db.
		WhereLike(treeField, tools.TrimLikeRight(BuildChildPath(nodeTree, id))).
		Fields(idField).
		Array()
}

// DescendantIdsById 先根据 id 查询节点信息，再查询这个节点的所有下级节点 ID。
func DescendantIdsById(db *gdb.Model, id int64) ([]gdb.Value, error) {
	node, err := GetNode(db, id)
	if err != nil || node == nil {
		return nil, err
	}
	return DescendantIds(db, node.Id, node.Tree)
}

// IsDescendant 判断 targetId 是否是 id 的下级节点。
//
// 常用于移动节点前校验，防止把节点移动到自己的子级下面。
func IsDescendant(db *gdb.Model, id int64, nodeTree string, targetId int64) (bool, error) {
	if targetId <= 0 {
		return false, nil
	}

	count, err := db.
		WherePri(targetId).
		WhereLike(treeField, tools.TrimLikeRight(BuildChildPath(nodeTree, id))).
		Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// RebuildDescendants 重算指定节点所有下级节点的 level 和 tree。
//
// 当前节点移动到新父级后，调用方先更新当前节点，再用这个方法递归修正所有下级链路。
func RebuildDescendants(db *gdb.Model, id int64, level int, nodeTree string) error {
	children, err := directChildren(db, id)
	if err != nil || len(children) == 0 {
		return err
	}

	childTree := BuildChildPath(nodeTree, id)
	childLevel := level + 1
	childIds := make([]int64, 0, len(children))
	for _, child := range children {
		childIds = append(childIds, child.Id)
	}

	if _, err = db.WhereIn(idField, childIds).Data(UpdateDO{
		Level: childLevel,
		Tree:  childTree,
	}).Update(); err != nil {
		return err
	}

	for _, child := range children {
		if err = RebuildDescendants(db, child.Id, childLevel, childTree); err != nil {
			return err
		}
	}
	return nil
}

// directChildren 查询指定节点的直属下级节点。
func directChildren(db *gdb.Model, parentId int64) ([]*Node, error) {
	records, err := db.
		Where(parentIdField, parentId).
		Fields(idField, parentIdField, levelField, treeField).
		All()
	if err != nil || len(records) == 0 {
		return nil, err
	}

	nodes := make([]*Node, 0, len(records))
	for _, record := range records {
		nodes = append(nodes, recordToNode(record))
	}
	return nodes, nil
}

// recordToNode 将数据库记录转换为通用树节点。
func recordToNode(record gdb.Record) *Node {
	return &Node{
		Id:       record[idField].Int64(),
		ParentId: record[parentIdField].Int64(),
		Level:    record[levelField].Int(),
		Tree:     record[treeField].String(),
	}
}
