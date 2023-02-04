package main

import (
	"assignment_bd/config"
	"assignment_bd/global"
	"assignment_bd/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var err error
	//数据库初始化，DB为全局变量
	global.DB, err = gorm.Open(mysql.Open(config.DSN_local), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 设置项目的 URL， 然后初始化路由
	r := server.Default()
	middleware.MyJwt()
	initRouter(r)
	r.Spin()
}
