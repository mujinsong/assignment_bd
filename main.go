package main

import (
	"assignment_bd/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init mysql
	dao.Init()
	r := gin.Default()
	// 加载路由
	initRouter(r)

	//pprof
	//pprof.Register(r)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run()
	if err != nil {
		return
	}
}
