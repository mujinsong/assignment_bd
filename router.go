package main

import (
	"assignment_bd/global"
	"github.com/cloudwego/hertz/pkg/app"
)

//在需要鉴权的路由上引入
func _taskMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{global.HzJwtMw.MiddlewareFunc()}
}
