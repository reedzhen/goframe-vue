package gencode

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/utility/tools"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	PlatformApi = "api"
	PlatformUi  = "ui"
)

type GenCode struct {
	ModuleName   string // 项目名称，请填写 go.mod 里的 module 名称
	TablePrefix, // 去除数组中定义好的表前缀
	ColumnHidden, // 数组中的字段将被隐藏
	ColumnSearch []string // 需要查询的字段
	Platform  []string // 默认生成前后端的代码
	WorkPath  string   // 工作目录
	LazyMode  string   // yes搁项目里面生成/no搁项目外面生成（建议no）
	AddonName string   // 插件名称，非插件代码填空
	JsonCase  string   // JSON tag casing: Camel, CamelLower, Kebab, Snake 默认小骆驼
}

// Run 执行代码自动生成
func (s *GenCode) Run(ctx context.Context, tableName string) error {

	// 获取表备注
	tableComment, err := s.getTableComment(ctx, tableName)
	if err != nil || strings.TrimSpace(tableComment) == "" {
		return gerror.New("请在数据库中给表加上备注")
	}

	tableShortName := s.getTableShortName(tableName)

	// 获取字段信息
	fields, err := s.getFields(ctx, tableName)
	if err != nil {
		return err
	}

	// 判断是否有时间类型，引入golang时间包
	hasTime := false
	for _, v := range fields {
		if gstr.Contains(v["type"], "*gtime.Time") {
			hasTime = true
			break
		}
	}

	// 模板变量
	fillData := g.Map{
		"module_name": s.ModuleName,
		"short_name":  tableShortName, // dict
		"name":        tableName,      // sys_dict
		"fields":      fields,
		"comment":     tableComment, // 表注释
		"addon_name":  s.AddonName,  // 插件名称
		"has_time":    hasTime,      // 是否引入时间包
	}

	// 插件名称
	addonName := strings.TrimSpace(s.AddonName)

	tableShortName2 := gstr.Replace(gstr.ToLower(tableShortName), "_", "")

	// 生成后台文件
	if tools.InSlice(s.Platform, PlatformApi) {
		for _, d := range []string{"model", "api", "logic", "controller"} {
			outputPath := s.getOutputPath(PlatformApi)

			if d == "controller" {
				controllerMethods := []string{"create", "update", "delete", "get_list", "get_page", "get_info"}
				for _, method := range controllerMethods {
					content, err := s.fillAndGetContent(fmt.Sprintf("controller_%s.tpl", method), PlatformApi, fillData)
					if err != nil {
						return err
					}

					path := fmt.Sprintf("%s/internal/%s/%s/admin_%s_%s.go", outputPath, d, "admin", tableShortName2, method)
					// 如果是插件,路径删除 internal
					if addonName != "" {
						path = gstr.Replace(path, "/internal", "", 1)
					}

					g.Dump(path)
					if err = gfile.PutContents(path, content); err != nil {
						return err
					}
				}
				continue
			}

			// 获取并填充模板
			content, err := s.fillAndGetContent(fmt.Sprintf("%s.tpl", d), PlatformApi, fillData)
			if err != nil {
				return err
			}

			path := ""
			switch d {
			case "api":
				path = fmt.Sprintf("%s/api/%s/%s/%s.go", outputPath, "admin", tableShortName2, tableShortName)
			case "model":
				path = fmt.Sprintf("%s/internal/%s/dto/%s.go", outputPath, d, tableShortName)
			case "logic":
				path = fmt.Sprintf("%s/internal/%s/%s/%s.go", outputPath, d, tableShortName2, tableShortName)
			}

			// 如果是插件,路径删除 internal
			if addonName != "" {
				path = gstr.Replace(path, "/internal", "", 1)
			}

			g.Dump(path)

			// 模板保存到文件
			if err = gfile.PutContents(path, content); err != nil {
				return err
			}

			// 执行 make ctrl 生成控制器
			if s.LazyMode == "yes" && d == "api" {
				cmd := exec.Command("make", "ctrl")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Printf("Error executing [make ctrl]: %v\n", err)
				} else {
					fmt.Println(fmt.Sprintf("🎉 🎉 🎉 1. The command [make ctrl] has been generated successfully  🎈 🎈 🎈"))
				}
			}

			// 执行 make service 生成service
			if s.LazyMode == "yes" && d == "logic" {
				cmd := exec.Command("make", "service")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Printf("Error executing [make service]: %v\n", err)
				} else {
					fmt.Println(fmt.Sprintf("🎉 🎉 🎉 2. The command [make service] has been generated successfully  🎈 🎈 🎈"))
				}
			}
		}
	}

	// 生成前端文件
	if tools.InSlice(s.Platform, PlatformUi) {
		for _, d := range []string{"api", "component_search", "component_edit", "index"} {
			// 填充模板
			content, err := s.fillAndGetContent(fmt.Sprintf("%s.tpl", d), PlatformUi, fillData)
			if err != nil {
				return err
			}

			outputPath := s.getOutputPath(PlatformUi)

			path := ""
			pathName := gstr.CaseKebab(tableShortName)
			switch d {
			case "api": // 接口列表
				path = fmt.Sprintf("%s/api/%s/index.js", outputPath, pathName)
			case "component_search": // 查询组件
				path = fmt.Sprintf("%s/views/%s/components/%s", outputPath, pathName, pathName+"-search.vue")
			case "component_edit": // 编辑组件
				path = fmt.Sprintf("%s/views/%s/components/%s", outputPath, pathName, pathName+"-edit.vue")
			case "index": // 页面
				path = fmt.Sprintf("%s/views/%s/index.vue", outputPath, pathName)
			}

			// 模板保存到文件
			if err = gfile.PutContents(path, content); err != nil {
				return err
			}
		}
	}

	return nil
}

// getOutputPath 获取代码输出目录
func (s *GenCode) getOutputPath(platform string) (path string) {
	if platform == PlatformApi {
		// 判断是否为插件代码
		if strings.TrimSpace(s.AddonName) != "" {
			if s.LazyMode == "yes" {
				path = filepath.Join(s.WorkPath, "addons", strings.TrimSpace(s.AddonName))
			} else {
				path = filepath.Join(s.WorkPath, "code", platform, "addons", strings.TrimSpace(s.AddonName))
			}
		} else {
			if s.LazyMode == "yes" {
				path = s.WorkPath
			} else {
				path = filepath.Join(s.WorkPath, "code", platform)
			}
		}
	}
	if platform == PlatformUi {
		path = filepath.Join(s.WorkPath, "code", platform)
	}
	return
}

// getTableShortName 获取没有前缀的表名称 t_user => user
func (s *GenCode) getTableShortName(tableName string) string {
	// 表名(不含表前缀) user
	shortName := tableName
	for _, prefix := range s.TablePrefix {
		if gstr.HasPrefix(tableName, prefix) {
			shortName = strings.TrimPrefix(tableName, prefix)
			break
		}
	}
	return shortName
}

// getFields 获取表格字段
func (s *GenCode) getFields(ctx context.Context, tableName string) ([]map[string]string, error) {
	// 获取字段信息（顺序不对）
	mapFields, err := g.DB().TableFields(ctx, tableName)
	if err != nil {
		return nil, err
	}

	// 获取正确顺序数据库字段
	sortedFields, err := s.getSortedTableFields(ctx, tableName)
	if err != nil {
		return nil, err
	}

	// 获取要展示的字段
	var fields []map[string]string
	for _, v := range sortedFields {
		field, ok := mapFields[v]
		if !ok {
			return nil, gerror.New("字段获取失败")
		}

		if !tools.InSlice(s.ColumnHidden, v) {
			t, err := g.DB().CheckLocalTypeForField(ctx, field.Type, true)
			if err != nil {
				return nil, err
			}

			typeGo := string(t)
			typeJs := "string"
			if typeGo == "int" || typeGo == "uint" || typeGo == "float64" {
				typeJs = "number"
			} else if typeGo == "datetime" {
				typeGo = "*gtime.Time"
			} else if typeGo == "time" {
				typeGo = "string"
			}

			needSearch := tools.InSlice(s.ColumnSearch, v)

			// 简化备注，默认按照空格分隔，前端作为label
			comment := field.Comment
			if field.Comment != "" {
				comment = gstr.SplitAndTrim(comment, " ")[0]
			}

			var jsonName string
			switch s.JsonCase {
			case "Camel":
				jsonName = gstr.CaseCamel(field.Name)
			case "CamelLower":
				jsonName = gstr.CaseCamelLower(field.Name)
			case "Kebab":
				jsonName = gstr.CaseKebab(field.Name)
			case "Snake":
				jsonName = field.Name // Default matches DB column name usually or explicit snake_case
			default:
				jsonName = gstr.CaseCamelLower(field.Name)
			}

			fields = append(fields, map[string]string{
				"name":          field.Name,               // name
				"json_name":     jsonName,                 // new: json tag name
				"type":          typeGo,                   // int string float...
				"comment_short": comment,                  // 姓名
				"comment":       field.Comment,            // 姓名 其他备注
				"type_js":       typeJs,                   // number string
				"need_search":   gconv.String(needSearch), // 查询字段
			})
		}
	}
	return fields, nil
}

// fillAndGetContent 填充模板并且获取填充后的模板
func (s *GenCode) fillAndGetContent(tplName string, tplPath string, fillData g.Map) (string, error) {
	// 创建模板引擎
	view := gview.New()
	view.BindFuncMap(g.Map{
		"UcFirst": func(str string) string {
			return gstr.UcFirst(str)
		},
		"CaseCamel": func(str string) string {
			return gstr.CaseCamel(str)
		},
		"CaseCamelLower": func(str string) string {
			return gstr.CaseCamelLower(str) // 小骆驼 reedTest
		},
		"CaseKebab": func(str string) string {
			return gstr.CaseKebab(str) // 中划线 reed-test
		},
		"CaseLower": func(str string) string {
			return gstr.Replace(gstr.ToLower(str), "_", "") // 小写 reedtest
		},
		"UcFirstCaseLower": func(str string) string {
			return gstr.UcFirst(gstr.Replace(gstr.ToLower(str), "_", "")) // 小写 reedtest,然后首字母大写，解决多单词命名问题
		},
	})
	// 配置
	_ = view.SetConfigWithMap(g.Map{
		//"Paths":      []string{s.WorkPath + "/resource/template/gen_code/" + tplPath}, // 模板文件搜索目录路径
		"Paths":      []string{s.WorkPath + "/internal/library/gencode/template/" + tplPath}, // 模板文件搜索目录路径
		"Delimiters": []string{"${", "}"},                                                    // 模板引擎变量分隔符号。默认为 ["{{", "}}"]
	})
	// 填充模板
	result, err := view.Parse(context.Background(), tplName, g.Map{"table": fillData})
	if err != nil {
		return "", err
	}

	// 剔除多余的换行
	content, err := s.trimBreak(result)
	if err != nil {
		return "", err
	}

	return content, nil
}

// getTableComment 获取表名
func (s *GenCode) getTableComment(ctx context.Context, tableName string) (tableComment string, err error) {
	dbName := g.DB().GetConfig().Name // 数据库名称
	data, err := g.DB().GetArray(ctx, `SELECT table_comment FROM information_schema.TABLES WHERE table_schema = ? and table_name = ?`, dbName, tableName)
	if err != nil {
		return
	}
	if len(data) > 0 {
		tableComment = strings.TrimRight(data[0].String(), "表")
	}
	return
}

// getSortedTableFields 获取所有字段(正确顺序)
func (s *GenCode) getSortedTableFields(ctx context.Context, tableName string) (columns []string, err error) {
	dbName := g.DB().GetConfig().Name // 数据库名称
	data, err := g.DB().GetArray(ctx, `select COLUMN_NAME from information_schema.COLUMNS where table_schema = ? and table_name = ?`, dbName, tableName)
	if err != nil {
		return
	}
	return gconv.Strings(data), nil
}

// trimBreak 剔除多余的换行
func (s *GenCode) trimBreak(str string) (string, error) {
	b, err := gregex.Replace("(([\\s\t]*)\r?\n){3,}", []byte("$3\n"), []byte(str))
	if err != nil {
		return "", err
	}
	return gconv.String(b), nil
}
