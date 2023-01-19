package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
	"assignment_bd/utility"
	"context"
	"errors"
	"gorm.io/gorm"
)

// Login 执行登录
func Login(ctx context.Context, in model.Login) error {
	//验证账号密码是否正确
	userInfo := model.User{}
	//检测有没有这个用户名
	err := global.DB.WithContext(ctx).Where("id = ?", in.Username).Take(&userInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("未注册")
	}
	//有这个用户名，检测密码是否对
	if utility.EncryptPassword(in.Password, userInfo.Salt) != userInfo.Password {
		return errors.New("帐号或密码不对")
	}
	//todo 后续根据情况处理

	return nil
}
