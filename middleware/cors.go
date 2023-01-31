package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	"time"
)

func CorsMw() app.HandlerFunc {
	return cors.New(cors.Config{
		// 允许跨源访问的 origin 列表
		AllowOrigins: []string{"*"},
		// 允许客户端跨源访问所使用的 HTTP 方法列表
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		// 允许使用的头信息字段列表
		AllowHeaders: []string{"Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma"},
		// 允许暴露给客户端的响应头列表
		ExposeHeaders: []string{"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar"},
		// 允许客户端请求携带用户凭证
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
