package config

import "gorm.io/gorm"

var (
	PankerDSN = "root:123456@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB        *gorm.DB
)
