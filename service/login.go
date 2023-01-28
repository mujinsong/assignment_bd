package service

import (
	"assignment_bd/dao"
	"assignment_bd/global"
	"assignment_bd/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

// Login 执行登录
func Login(ctx context.Context, in dao.Login) (user *dao.User, err error) {
	//检测有没有这个用户名
	err = global.DB.WithContext(ctx).Where("username = ?", in.Username).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("未注册")
	}
	//有这个用户名，检测密码是否对
	if utils.EncryptPassword(in.Password, user.Salt) != user.Password {
		return nil, errors.New("帐号或密码不对")
	}
	//todo 后续根据情况处理

	return
}
