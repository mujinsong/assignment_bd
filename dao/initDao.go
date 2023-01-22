package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
// 尝试使用gorm连接mysql数据库

// DB 数据库连接池

var DB *gorm.DB

func InitDB() {
	
}