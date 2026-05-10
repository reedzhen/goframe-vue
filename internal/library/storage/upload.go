package storage

import (
	"bytes"
	"context"
	"github.com/disintegration/imaging"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-vben/internal/consts"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math"
)

type UploadDrive interface {
	UploadFile(ctx context.Context, in UploadParam) (url string, err error)
}

func New(name ...string) UploadDrive {
	var (
		driveType = consts.UploadDriveLocal
		drive     UploadDrive
	)

	if len(name) > 0 && name[0] != "" {
		driveType = name[0]
	}

	switch driveType {
	case consts.UploadDriveLocal:
		drive = &LocalDrive{}
	case consts.UploadDriveOss:
		drive = &OssDrive{}
	}
	return drive
}

// CompressImage 调整图片尺寸 主要目的是确保图像的最长边不超过 maxEdge，同时保持图像的宽高比不变
// 可以传入 bytes.NewReader([]byte) 或者 file 对象
func CompressImage(file io.Reader, maxEdge int) ([]byte, error) {
	// 尝试解码不同格式的图像
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, gerror.Newf("failed to decode image: %v", err)
	}

	// 获取图像尺寸
	x := float64(img.Bounds().Dx())
	y := float64(img.Bounds().Dy())

	// 调整图像尺寸
	width := int(x)
	height := int(y)

	if math.Max(x, y) > float64(maxEdge) {
		if x > y {
			width = maxEdge
			height = int(math.Round(float64(maxEdge) * (y / x)))
		} else {
			width = int(math.Round(float64(maxEdge) * (x / y)))
			height = maxEdge
		}
	}

	// 调整图像
	resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)

	// 创建缓冲区
	b := bytes.NewBuffer(nil)

	// 根据图像格式保存
	switch format {
	case "png":
		err = png.Encode(b, resizedImg)
	case "jpeg", "jpg":
		err = jpeg.Encode(b, resizedImg, nil)
	default:
		return nil, gerror.Newf("unsupported image format: %s", format)
	}
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
