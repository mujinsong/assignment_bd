package main

import (
	"assignment_bd/config"
	"assignment_bd/controller"
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
	//todo 这块随便写的，后续改
	r := server.Default()
	r.Use(middleware.CorsMw())
	middleware.JwtMwInit()
	r.POST("/douyin/user/register/", controller.Register)
	r.POST("/douyin/user/login/", controller.Login)
	err = r.Run()
	if err != nil {
		return
	}
}
