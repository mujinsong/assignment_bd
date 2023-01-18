package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initRouter(r)
	// 监听8000端口 8080端口感觉有点常用 tomcat也是默认跑8080端口
	r.Run(":8000")  
}
