package storage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// 定义 MIME 类型常量
const (
	MimeJpeg        = "image/jpeg"
	MimePng         = "image/png"
	MimeGif         = "image/gif"
	MimeBmp         = "image/bmp"
	MimeTiff        = "image/tiff"
	MimeSvg         = "image/svg+xml"
	MimeWebp        = "image/webp"
	MimePdf         = "application/pdf"
	MimeDoc         = "application/msword"
	MimeDocx        = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MimePpt         = "application/vnd.ms-powerpoint"
	MimePptx        = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	MimeXls         = "application/vnd.ms-excel"
	MimeXlsx        = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	MimeTxt         = "text/plain"
	MimeRtf         = "application/rtf"
	MimeOdt         = "application/vnd.oasis.opendocument.text"
	MimeOds         = "application/vnd.oasis.opendocument.spreadsheet"
	MimeOdp         = "application/vnd.oasis.opendocument.presentation"
	MimeWps         = "application/vnd.ms-works"
	MimeCsv         = "text/csv"
	MimeMp4         = "video/mp4"
	MimeAvi         = "video/x-msvideo"
	MimeMov         = "video/quicktime"
	MimeMkv         = "video/x-matroska"
	MimeWebm        = "video/webm"
	MimeFlv         = "video/x-flv"
	MimeMpeg        = "video/mpeg"
	MimeMp3         = "audio/mpeg"
	MimeWav         = "audio/wav"
	MimeOgg         = "audio/ogg"
	MimeAac         = "audio/aac"
	MimeFlac        = "audio/flac"
	MimeZip         = "application/zip"
	MimeTar         = "application/x-tar"
	MimeGzip        = "application/gzip"
	MimeBz2         = "application/x-bzip2"
	Mime7z          = "application/x-7z-compressed"
	MimeRar         = "application/x-rar-compressed"
	MimeOctetStream = "application/octet-stream"
)

// 文件扩展名和 MIME 类型的映射表
var contentTypeMap = map[string]string{
	".jpg":  MimeJpeg,
	".jpeg": MimeJpeg,
	".png":  MimePng,
	".gif":  MimeGif,
	".bmp":  MimeBmp,
	".tiff": MimeTiff,
	".tif":  MimeTiff,
	".svg":  MimeSvg,
	".webp": MimeWebp,
	".pdf":  MimePdf,
	".doc":  MimeDoc,
	".docx": MimeDocx,
	".ppt":  MimePpt,
	".pptx": MimePptx,
	".xls":  MimeXls,
	".xlsx": MimeXlsx,
	".txt":  MimeTxt,
	".rtf":  MimeRtf,
	".odt":  MimeOdt,
	".ods":  MimeOds,
	".odp":  MimeOdp,
	".wps":  MimeWps,
	".csv":  MimeCsv,
	".mp4":  MimeMp4,
	".avi":  MimeAvi,
	".mov":  MimeMov,
	".mkv":  MimeMkv,
	".webm": MimeWebm,
	".flv":  MimeFlv,
	".mpeg": MimeMpeg,
	".mpg":  MimeMpeg,
	".mp3":  MimeMp3,
	".wav":  MimeWav,
	".ogg":  MimeOgg,
	".aac":  MimeAac,
	".flac": MimeFlac,
	".zip":  MimeZip,
	".tar":  MimeTar,
	".gz":   MimeGzip,
	".gzip": MimeGzip,
	".bz2":  MimeBz2,
	".7z":   Mime7z,
	".rar":  MimeRar,
}

// getContentType 根据文件扩展名获取内容类型
func getContentType(ext string) string {
	// 输入验证
	if len(ext) == 0 || ext[0] != '.' {
		g.Log().Errorf(context.Background(), "Invalid file extension: %s", ext)
		return MimeOctetStream
	}

	// 查找 MIME 类型
	if contentType, ok := contentTypeMap[ext]; ok {
		return contentType
	}

	// 默认返回值
	g.Log().Warningf(context.Background(), "Unknown file extension: %s, returning default MIME type: %s", ext, MimeOctetStream)
	return MimeOctetStream
}
