package tools

import (
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"time"
)

// GenerateCheckSum 生成CheckSum
func GenerateCheckSum(appKey, appSecret string) map[string]string {
	headers := make(map[string]string)
	curTime := gconv.String(time.Now().Unix())
	nonce := grand.S(16)

	checkSum := gsha1.Encrypt(appSecret + nonce + curTime)
	headers["Timestamp"] = curTime
	headers["NonceStr"] = nonce
	headers["CheckSum"] = checkSum
	headers["AppKey"] = appKey
	return headers
}

// ValidateCheckSum 验证checksum是否有效
func ValidateCheckSum(checksum, timestamp, nonceStr, appSecret string) bool {
	checkSum := gsha1.Encrypt(appSecret + nonceStr + timestamp)
	return checkSum == checksum
}
