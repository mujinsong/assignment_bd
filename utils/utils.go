package utils

import (
	"assignment_bd/consts"
	"github.com/gogf/gf/v2/crypto/gmd5"
)

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := consts.MsgFlags[code]
	if ok {
		return msg
	}
	return consts.MsgFlags[consts.ERROR]
}
