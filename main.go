package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/config"
)

func main() {
	config.DB, _ = gorm.Open(mysql.Open(config.PankerDSN), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("连接数据库异常")
	//	panic(err) //抛出异常
	//}
	initRouter(gin.Default())
}
