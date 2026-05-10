package tools

// 泛型使用，封装了开发过程中常用的几个方法，Go 1.18 引入了泛型支持，注意查看golang版本

// InSlice 元素是否存在于切片中 这样 []int []int64 []string都可通用
func InSlice[K comparable](slice []K, key K) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}

// SliceToMapByField 对象切片转 map，key=user.id value = user
func SliceToMapByField[T any, K comparable](slice []T, keyFunc func(T) K) map[K]T {
	result := make(map[K]T)
	for _, v := range slice {
		key := keyFunc(v)
		result[key] = v
	}
	return result
}

// SliceToMapsByField 对象切片转 map 数组，key = user.id value = []user
func SliceToMapsByField[T any, K comparable](slice []T, fieldFunc func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := fieldFunc(v)
		result[key] = append(result[key], v)
	}
	return result
}

// SliceToMapsWithFieldFunc 对象切片转 map，key = user.name, value = 指定字段值的数组 (如 user.id)
func SliceToMapsWithFieldFunc[T any, K comparable, V any](slice []T, keyFunc func(T) K, valueFunc func(T) V) map[K][]V {
	result := make(map[K][]V)
	for _, v := range slice {
		key := keyFunc(v)
		value := valueFunc(v)
		result[key] = append(result[key], value)
	}
	return result
}

// DiffSlice 接收两个切片 a 和 b，a 是基数组，返回三个切片：
// 第一个切片包含 b 中但不在 a 中的元素，用于新增
// 第二个切片包含 a 中但不在 b 中的元素，用于删除
// 第三个切片包含 a 和 b 共同拥有的元素，用于编辑
func DiffSlice[T comparable](a, b []T) (added, deleted, edited []T) {
	bMap := make(map[T]bool)
	for _, item := range b {
		bMap[item] = true
	}

	for _, item := range a {
		if _, found := bMap[item]; found {
			edited = append(edited, item)
		} else {
			deleted = append(deleted, item)
		}
	}

	for item := range bMap {
		if !InSlice(a, item) {
			added = append(added, item)
		}
	}

	return added, deleted, edited
}

// DiffList 对比老、新两个列表，找出新增、修改、删除的数据
func DiffList[T comparable](oldList, newList []T, sameFunc func(T, T) bool) (createList, updateList, deleteList []T) {
	createList = append(createList, newList...) // 默认都认为是新增的，后续会进行移除

	// 通过以 oldList 为主遍历，找出 updateList 和 deleteList
	for _, oldObj := range oldList {
		isFound := false // 标记是否找到匹配的对象
		var foundObj T   // 用于存储找到的匹配对象

		for i := 0; i < len(createList); i++ {
			newObj := createList[i]
			if sameFunc(oldObj, newObj) {
				foundObj = newObj
				isFound = true
				// 移除找到的对象
				createList = append(createList[:i], createList[i+1:]...) // 移除元素
				break
			}
		}

		// 匹配添加到 updateList；不匹配则添加到 deleteList
		if isFound {
			updateList = append(updateList, foundObj)
		} else {
			deleteList = append(deleteList, oldObj)
		}
	}
	return
}
