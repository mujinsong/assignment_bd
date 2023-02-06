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
	global.DB, err = gorm.Open(mysql.Open(config.UseDSN), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 设置项目的 URL， 然后初始化路由
	r := server.Default(
		server.WithMaxRequestBodySize(1024*1024*1000), // 运行上传的最大文件大小 1000M
		server.WithHostPorts(config.Port),                 // 项目运行的端口
		server.WithStreamBody(true),                   // 开启流式 body
		server.WithIdleTimeout(60),                    // 连接空闲超时时间
	)
	middleware.MyJwt()
	initRouter(r)
	r.Spin()
}
