package password

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/grand"
)

// Generate 生成密码
func Generate(pwd string) (string, string) {
	salt := grand.S(6, false)
	return gmd5.MustEncryptString(gmd5.MustEncryptString(pwd) + salt), salt
}

// Verify 验证密码
// @dbPwd 数据库存储密码
// @inputPwd 前端输入密码
// @slat 密码加密盐
func Verify(dbPwd, inputPwd, salt string) bool {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(inputPwd)+salt) == dbPwd
}
