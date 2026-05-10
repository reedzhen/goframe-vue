package excel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"net/url"
	"time"
)

//// 构造导出数据
//header := strings.Join([]string{"订单号", "原始订单号"}, ",")
//bf := bytes.NewBufferString(header + "\n")
//for _, v := range list {
//	body := strings.Join([]string{v.OrderNo, v.OriginOrderNo}, ",")
//	bf.WriteString(body + "\n")
//}
//data := bf.Bytes()
//
//// 执行导出
//excel.Export(ctx, data, "订单资料")

// ExportCsv 导出csv
func ExportCsv(ctx context.Context, data []byte, fileName string) {
	ret := g.RequestFromCtx(ctx).Response
	ret.SetBuffer(data)
	ret.Header().Set("Content-Length", gconv.String(ret.BufferLength()))
	ret.Header().Set("Content-Type", "application/force-download")
	ret.Header().Set("Accept-Ranges", "bytes")
	ret.Header().Set("Content-Transfer-Encoding", "binary")
	downloadName := fmt.Sprintf("%s-%s.csv", time.Now().Format("060102150405"), fileName)
	ret.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(downloadName)))
	ret.Buffer()
}
