package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"goframe-vben/internal/library/gencode"
)

// todo 一定要几个把数据库迁移给关了
var (
	Code = gcmd.Command{
		Name:  "code",
		Usage: "code auto generate",
		Brief: "go run main.go code",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			workPath := gfile.Pwd()

			s := &gencode.GenCode{
				ModuleName:   "goframe-vben",                                                                      // 项目名称，请填写 go.mod 里的 module 名称
				WorkPath:     workPath,                                                                            // 工作目录
				TablePrefix:  []string{"sys_", "t_", "app_", "erp_", "hotel_", "wyf_", "dm_"},                     // 去除表前缀
				ColumnHidden: []string{"created_at", "updated_at", "updated_by", "created_by", "id", "tenant_id"}, // 过滤表字段
				Platform:     []string{"api"},                                                                     // 默认生成前后端的代码 "api", "ui"
				LazyMode:     "no",                                                                                // yes搁项目里面生成/no搁项目外面生成（建议使用no）
				AddonName:    "fss",                                                                               // 插件名称，非插件代码填空
				ColumnSearch: []string{"place_id", "package_id", "partner_id"},                                    // 生成查询条件的字段
			}

			tableName := "fss_place_package"
			if err := s.Run(ctx, tableName); err != nil {
				g.Dump(err)
			} else {
				fmt.Println(fmt.Sprintf("🎉 🎉 🎉 3. The table [%s] has been generated successfully  🎈 🎈 🎈", tableName))
			}

			// todo 主键int64写死了
			return
		},
	}
)
