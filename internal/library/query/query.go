package query

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"strings"
	"unicode"
)

// Query 分页查询条件
type Query interface {
	GetPageIndex() int
	GetPageSize() int
	Cond(*gdb.Model) *gdb.Model
	GetOrder() string
}

// PageParam 分页参数
type PageParam struct {
	Page  int    `json:"page" in:"query" dc:"当前页 默认：1"`
	Limit int    `json:"limit" in:"query" dc:"每页数量 默认：20 最大：100"`
	Sort  string `json:"sort" in:"query" dc:"排序字段"`
	Order string `json:"order" in:"query" dc:"排序方式 asc desc"`
}

// GetPageIndex 获取当前页码
func (p *PageParam) GetPageIndex() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

// GetPageSize 获取分页大小
func (p *PageParam) GetPageSize() int {
	if p.Limit <= 0 {
		return 20
	}
	if p.Limit > 100 {
		return 100
	}
	return p.Limit
}

// Asc 正序排
func (p *PageParam) Asc(column string) {
	p.Sort = column
	p.Order = "asc"
}

// Desc 倒序排
func (p *PageParam) Desc(column string) {
	p.Sort = column
	p.Order = "desc"
}

// GetOrder 获取排序
func (p *PageParam) GetOrder() string {
	sort, order := strings.TrimSpace(p.Sort), strings.ToLower(strings.TrimSpace(p.Order))
	if sort == "" || order == "" {
		return ""
	}
	if order != "asc" && order != "desc" {
		return ""
	}
	if !isSafeOrderField(sort) {
		return ""
	}
	return fmt.Sprintf("%s %s", sort, order)
}

func isSafeOrderField(field string) bool {
	field = strings.TrimSpace(field)
	if field == "" {
		return false
	}
	for _, r := range field {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '.' {
			continue
		}
		return false
	}
	return true
}

// Result 分页结果
type Result struct {
	List interface{} `json:"list"`
	//Current int         `json:"current"`
	//Pages   int         `json:"pages"`
	//Size    int         `json:"size"`
	Count int `json:"count"`
}

func (r *Result) WithRecords(data interface{}) *Result {
	r.List = data
	return r
}

// Page 分页查询
func Page(m *gdb.Model, query Query, bean interface{}) (out *Result, err error) {
	pageIndex := query.GetPageIndex()
	pageSize := query.GetPageSize()

	out = &Result{
		List: []interface{}{},
		//Current: pageIndex,
		//Pages:   0,
		//Size:    pageSize,
	}

	// 条件查询
	m = query.Cond(m)

	// 分页
	listM := m.Page(pageIndex, pageSize)

	// 排序
	order := query.GetOrder()
	if len(order) > 0 {
		listM = listM.Order(order)
	}

	// 查询并获取总行数
	if err = listM.ScanAndCount(bean, &out.Count, false); err != nil {
		return
	}

	// 总共多少页
	//out.Pages = int(math.Ceil(float64(out.Total) / float64(out.Size)))

	return out.WithRecords(bean), nil
}
