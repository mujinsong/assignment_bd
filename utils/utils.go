package utils

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"errors"
	"github.com/hertz-contrib/jwt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gogf/gf/v2/crypto/gmd5"
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

/*
从请求上下文中获取uid
*/
func GetUid(c *app.RequestContext) (uid uint64, err error) {
	value, exists := c.Get(consts.IdentityKey)
	if !exists {
		return 0, errors.New("获取value失败")
	}
	user, ok := value.(model.User)
	if !ok {
		return 0, errors.New("断言错误")
	}
	uid = user.ID
	return
}

func GetUidFromToken(token string) (uid uint64, err error) {
	tokenValue, err := global.HzJwtMw.ParseTokenString(token)
	if err != nil {
		return 0, err
	}
	claims := jwt.ExtractClaimsFromToken(tokenValue)
	uid = uint64(int(claims[consts.IdentityKey].(float64)))
	return uid, nil
}

func StrToUint64(str string) uint64 {
	num, _ := strconv.Atoi(str)
	return uint64(num)
}
func StrToUint8(str string) uint8 {
	num, _ := strconv.Atoi(str)
	return uint8(num)
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetUnionID(user1, user2 uint64) string {
	return strconv.Itoa(int(user1)) + "_" + strconv.Itoa(int(user2))
}
