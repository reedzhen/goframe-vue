package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/upload" method:"post" mime:"multipart/form-data" tags:"通用接口" summary:"上传图片"`
	Dir    string            `json:"dir" dc:"目录"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type FileUploadRes struct {
	Url string `json:"url" dc:"文件直接访问地址"`
}
