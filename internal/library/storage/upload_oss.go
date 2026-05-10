package storage

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"goframe-vben/internal/service"
	"path"
	"strconv"
	"strings"
)

// OssDrive 阿里云oss驱动
type OssDrive struct {
}

// UploadFile 上传到阿里云oss
func (d *OssDrive) UploadFile(ctx context.Context, in UploadParam) (url string, err error) {
	// 获取上传配置
	conf, err := service.Config().GetUpload(ctx)
	if err != nil {
		return
	}

	// 打开文件
	file, err := in.File.Open()
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()

	// 获取文件后缀
	ext := path.Ext(in.File.Filename)

	// 生成文件名
	fileName := in.Name + ext
	if in.Name == "" {
		fileName = in.File.Filename
	}

	// 如果需要随机文件名，则生成随机文件名
	if in.RandomName {
		fileName = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(6)) + ext
	}

	// 清理目录路径，防止路径遍历攻击
	dir := path.Clean(in.Dir)
	// 生成文件路径
	filePath := fmt.Sprintf("%s/%s", dir, fileName)

	// 实例化oss
	client, err := oss.New(conf.UploadOssEndpoint, conf.UploadOssAccessKeyId, conf.UploadOssAccessKeySecret)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(conf.UploadOssBucket)
	if err != nil {
		return "", err
	}

	// 上传文件
	contentType := getContentType(ext)
	option := oss.ContentType(contentType)
	if err = bucket.PutObject(filePath, file, option); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", conf.UploadOssBucketUrl, filePath), nil // 直接访问地址
}
