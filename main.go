package main

import (
	"tiktok/config"
	"tiktok/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initRouter(r)
	dao.InitDB()
	r.Run(config.Port)
}
