package main

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func main() {
	var err error
	//数据库初始化，DB为全局变量
	global.DB, err = gorm.Open(mysql.Open(consts.DSN), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//todo 这块随便写的，后续改
	r := gin.Default()
	r.GET("/douyin/user/:UserId", func(c *gin.Context) {
		userIdStr := c.Param("UserId")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return
		}
		_, err = service.UserInfoGetByUserID(uint(userId))
		if err != nil {
			return
		}
	})

	err = r.Run()
	if err != nil {
		return
	}
}
