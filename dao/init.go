package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/config"
)

// 我们需要加入这些东西

var DB *gorm.DB

func InitDB() {
	DB, _ = gorm.Open(mysql.Open(config.KM911LocalDSN), &gorm.Config{})
	if DB.Error != nil {
		println("database connect error")
		panic(DB.Error)
	}
	DB.AutoMigrate(&Users{}, &Videos{})

}
