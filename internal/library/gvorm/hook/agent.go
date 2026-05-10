package hook

//// AgentCreate 新增代理
//var AgentCreate = gdb.HookHandler{
//	Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
//		for _, record := range in.Data {
//			if gconv.Int64(record["pid"]) == 0 {
//				record["level"] = 0
//				record["path"] = "-"
//			} else {
//				var parent gdb.Record
//				parent, err = dao.SysAgent.Ctx(ctx).WherePri(record["pid"]).One()
//				if err != nil {
//					return
//				}
//				// 将层级设为父类目的层级 + 1
//				record["level"] = parent["level"].Int() + 1
//				// 将 path 值设为父类目的 path 追加父级 ID 以及最后跟上一个 - 分隔符
//				record["path"] = fmt.Sprintf("%s%d-", parent["path"].String(), parent["id"].Int64())
//			}
//		}
//
//		result, err = in.Next(ctx)
//		if err != nil {
//			return
//		}
//
//		//// 定义一个访问器，获取所有祖先类目的 ID 值
//		//public function getPathIdsAttribute()
//		//{
//		//	// trim($str, '-') 将字符串两端的 - 符号去除
//		//	// explode() 将字符串以 - 为分隔切割为数组
//		//	// 最后 array_filter 将数组中的空值移除
//		//	return array_filter(explode('-', trim($this->path, '-')));
//		//}
//		//
//		//// 定义一个访问器，获取所有祖先类目并按层级排序
//		//public function getAncestorsAttribute()
//		//{
//		//	return Category::query()
//		//	// 使用上面的访问器获取所有祖先类目 ID
//		//	->whereIn('id', $this->path_ids)
//		//	// 按层级排序
//		//	->orderBy('level')
//		//	->get();
//		//}
//
//		return
//	},
//}
