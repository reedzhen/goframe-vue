package storage

import "github.com/gogf/gf/v2/net/ghttp"

// UploadParam 上传参数
type UploadParam struct {
	File       *ghttp.UploadFile // 上传文件对象
	Dir        string            // 存储文件夹
	RandomName bool              // 是否是随机名称 若false取下面Name
	Name       string            // 文件名
}
