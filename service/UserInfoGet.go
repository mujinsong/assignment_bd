package service

import (
	_ "assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"errors"
	"gorm.io/gorm"
)

func UserInfoGetByUserID(userID uint) (user *model.User, err error) {
	// 检查 userID 是否存在；若存在，获取用户信息
	err = global.DB.Where("user_id = ?", userID).Limit(1).Find(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

//todo
