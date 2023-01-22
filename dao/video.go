package dao

import "gorm.io/gorm"

type Videos struct {
	gorm.Model
	Id       int    `gorm:"column:id;type:int unsigned;not null;primaryKey"`
	UserId   int    `gorm:"column:user_id;type:int unsigned;not null;primaryKey"`
	VideoUrl string `gorm:"column:video_url;type:varchar(32);not null;primaryKey"`
	CoverUrl string `gorm:"column:cover_url;type:varchar(32);not null;primaryKey"`
	Title    string `gorm:"column:title;type:varchar(32);not null;primaryKey"`
}
