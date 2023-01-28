package service

import (
	"assignment_bd/dao"
	"assignment_bd/global"
	"assignment_bd/utils"
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/grand"
)

//todo 注册，因为和登录一样需要传入的都是用户名的密码，所以我这里传的model.Login(可改)
func register(ctx context.Context, in *dao.Login) (out *dao.User, err error) {
	//密码用户名不能为空
	if in.Password == "" || in.Username == "" {
		return nil, errors.New("用户名和密码不能为空")
	}
	//查询是否已有该用户名
	result := global.DB.WithContext(ctx).Where("username = ?", in.Username).Take(&out)
	if result.RowsAffected != 0 {
		return nil, errors.New("已有该用户名，请登录或换一个用户名注册")
	}
	//加密
	userSalt := grand.S(10)
	out.Password = utils.EncryptPassword(in.Password, userSalt)
	out.Username = in.Username
	out.Salt = userSalt
	//插入数据库
	err = global.DB.WithContext(ctx).Create(out).Error
	return
}
