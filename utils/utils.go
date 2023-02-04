package utils

import (
	"assignment_bd/consts"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"math/rand"
	"regexp"
	"time"
	"unicode/utf8"
)

func VerifyUsernamePassword(username, password string) (bool, error) {
	//密码用户名不能为空
	if password == "" || username == "" {
		return false, errors.New("用户名和密码不能为空")
	}
	// 验证用户名合法性
	if utf8.RuneCountInString(username) > consts.MAX_USERNAME_LENGTH ||
		utf8.RuneCountInString(username) <= 0 {
		return false, errors.New("非法用户名")
	}
	// 验证密码合法性
	if ok, _ := regexp.MatchString(consts.MIN_PASSWORD_PATTERN, password); !ok {
		return false, errors.New("密码长度6-32，由字母大小写下划线组成")
	}
	return true, nil
}

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
