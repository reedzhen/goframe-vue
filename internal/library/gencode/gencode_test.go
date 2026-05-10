package gencode

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"testing"
)

func Test_sGenCode_Run(t *testing.T) {
	// 项目路径 /Users/reedzhen/Code/go/src/goframe-vben
	workPath := gstr.TrimRightStr(gfile.Pwd(), "/internal/library/gencode")

	s := &GenCode{
		ModuleName:   "goframe-vben",                                                                      // 项目名称，请填写 go.mod 里的 module 名称
		WorkPath:     workPath,                                                                            // 工作目录
		TablePrefix:  []string{"sys_", "t_", "app_"},                                                      // 去除表前缀
		ColumnHidden: []string{"created_at", "updated_at", "updated_by", "created_by", "id", "tenant_id"}, // 过滤表字段
		Platform:     []string{"api"},                                                                     // 默认生成前后端的代码 "api", "ui"
		LazyMode:     "yes",                                                                               // yes搁项目里面生成/no搁项目外面生成（建议使用这种方式）
		AddonName:    "",                                                                                  // 插件名称，非插件代码填空
		ColumnSearch: []string{"dept_code", "dept_name"},                                                  // 生成查询条件的字段
	}

	tableName := "sys_dept"
	if err := s.Run(context.Background(), tableName); err != nil {
		g.Dump(err)
	}

	// gf gen service -f Lower
	// cmd 添加controller
	// 添加菜单
}
