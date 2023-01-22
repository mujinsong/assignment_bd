package dao

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id       int    `gorm:"column:id;type:int unsigned;not null;primaryKey"`
	Username string `gorm:"column:username;type:varchar(32);not null;primaryKey"`
	Password string `gorm:"column:password;type:varchar(32);not null;primaryKey"`
}
