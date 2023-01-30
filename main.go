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
	global.DB, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	middleware.JwtMwInit()

	r := server.Default()
	r.Use(middleware.CorsMw())

	{
		liekGroup := r.Group("/douyin/favorite", global.HzJwtMw.MiddlewareFunc())
		{
			// 点赞
			liekGroup.POST("/action/", controller.Like)
			// 获取用户的点赞列表
			liekGroup.GET("/list/", controller.LikeList)
		}

		commentGroup := r.Group("/douyin/comment", global.HzJwtMw.MiddlewareFunc())
		{
			// 评论
			commentGroup.POST("/action", controller.Comment)
			// 评论列表
			commentGroup.GET("/list", controller.CommentList)
		}
	}

	err = r.Run()
	if err != nil {
		return
	}
}
