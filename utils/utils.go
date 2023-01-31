package utils

import (
	"assignment_bd/consts"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"math/rand"
	"time"
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

func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
