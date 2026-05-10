package excel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"net/url"
	"time"
)

// Export 导出excel
func Export(ctx context.Context, header []string, rows [][]interface{}, fileName string, sheetName string) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			g.Log().Error(ctx, err)
		}
	}()

	// 创建一个工作表
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}

	// 第一行标题
	_ = f.SetSheetRow(sheetName, "A1", &header)

	// 设置单元格的值
	for k, row := range rows {
		// 从A2开始写入数据
		// row := []interface{}{v.SkuNo, v.SkuName, v.BizCnt, v.BizPrice, v.CreatedAt, v.OriginOrderNo, v.CustName, v.WarehouseName}
		if err := f.SetSheetRow(sheetName, "A"+gconv.String(k+2), &row); err != nil {
			return err
		}
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	w := g.RequestFromCtx(ctx).Response
	w.Header().Set("Content-Type", "application/octet-stream")
	downloadName := fmt.Sprintf("%s-%s", time.Now().Format("060102150405"), fileName)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(downloadName)))
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	if err := f.Write(w.Writer); err != nil {
		return err
	}
	return nil
}
