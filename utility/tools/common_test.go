package tools

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

// 示例结构体
type User struct {
	ID   int
	Name string
}

func TestInSlice(t *testing.T) {
	// 用户切片
	names := []string{"Alice", "Bob", "Charlie"}
	g.Dump(InSlice(names, "Bob")) // true
	g.Dump(InSlice(names, "bob")) // false
}

func TestToMapByKey(t *testing.T) {
	// 用户切片
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	// 使用ID作为键转换为映射
	userMapById := SliceToMapByField(users, func(u User) int { return u.ID })
	g.Dump(userMapById)
	fmt.Println("By ID:", userMapById)

	// 使用Name作为键转换为映射
	userMapByName := SliceToMapByField(users, func(u User) string { return u.Name })
	fmt.Println("By Name:", userMapByName)
}

func TestGroupAndConvertToMapsByField(t *testing.T) {
	// 用户切片
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Alice"},
	}

	// 使用Name作为键分组并转换为映射数组
	groupedUsers := SliceToMapsByField(users, func(u User) string { return u.Name })
	g.Dump(groupedUsers)

	//{
	//	"Alice": [
	//		{
	//			ID:   1,
	//			Name: "Alice",
	//		},
	//		{
	//			ID:   3,
	//			Name: "Alice",
	//		},
	//	],
	//    "Bob":   [
	//		{
	//			ID:   2,
	//			Name: "Bob",
	//		},
	//	],
	//}

	// 打印结果
	for key, group := range groupedUsers {
		fmt.Printf("%s: %v\n", key, group)
	}
}

func TestSliceToMapsWithFieldFunc(t *testing.T) {
	// 用户切片
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Alice"},
	}

	// 使用Name作为键，ID数组作为值
	result := SliceToMapsWithFieldFunc(users, func(u User) string { return u.Name }, func(u User) int { return u.ID })
	g.Dump(result)

	// 预期结果:
	// {
	//     "Alice": [1, 3],
	//     "Bob":   [2],
	// }

	// 打印结果
	for key, ids := range result {
		fmt.Printf("%s: %v\n", key, ids)
	}
}

func TestDiffSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 4, 6}
	added, deleted, edited := DiffSlice(a, b)
	fmt.Println("Added:", added)
	fmt.Println("Deleted:", deleted)
	fmt.Println("Edited:", edited)

	//Added: [6]
	//Deleted: [1 3 5]
	//Edited: [2 4]

	// 测试字符串类型的示例
	aStr := []string{"apple", "banana", "cherry"}
	bStr := []string{"banana", "date", "fig"}
	addedStr, deletedStr, editedStr := DiffSlice(aStr, bStr)
	fmt.Println("Added (Strings):", addedStr)
	fmt.Println("Deleted (Strings):", deletedStr)
	fmt.Println("Edited (Strings):", editedStr)

	//Added (Strings): [date fig]
	//Deleted (Strings): [apple cherry]
	//Edited (Strings): [banana]
}

func TestDiffList(t *testing.T) {
	// 示例：使用 DiffList 函数
	type Person struct {
		ID   int
		Name string
	}

	// 比较函数
	sameFunc := func(p1, p2 Person) bool {
		return p1.ID == p2.ID
	}

	oldList := []Person{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	newList := []Person{
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	createList, updateList, deleteList := DiffList(oldList, newList, sameFunc)

	// 输出结果
	fmt.Println("新增列表:", createList)
	fmt.Println("修改列表:", updateList)
	fmt.Println("删除列表:", deleteList)

	//新增列表: [{3 Charlie}]
	//修改列表: [{2 Bob}]
	//删除列表: [{1 Alice}]
}
