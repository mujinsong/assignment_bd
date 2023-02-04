package main

import (
	"assignment_bd/controller"
	"assignment_bd/global"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func initRouter(r *server.Hertz) {
	//todo 这块随便写的，后续改
	//r.Use(middleware.CorsMw())
	//err := global.HzJwtMw.MiddlewareInit()
	//if err != nil {
	//	return
	//}
	r.POST("/douyin/user/login/", global.HzJwtMw.LoginHandler) // 用户登录接口
	apiRouter := r.Group("/douyin", global.HzJwtMw.MiddlewareFunc())
	{
		// 基础接口
		apiRouter.GET("/feed/", controller.Feed)               // 视频流接口
		apiRouter.GET("/user/", controller.UserInfo)           // 用户信息
		apiRouter.POST("/user/register/", controller.Register) // 用户注册接口

		apiRouter.POST("/publish/action/", controller.Publish)  // 视频投稿
		apiRouter.GET("/publish/list/", controller.PublishList) // 视频发布列表

		// 互动接口
		apiRouter.POST("/favorite/action/", controller.FavoriteAction) // 赞操作
		apiRouter.GET("/favorite/list/", controller.FavoriteList)      // 喜欢列表
		apiRouter.POST("/comment/action/", controller.CommentAction)   // 评论操作
		apiRouter.GET("/comment/list/", controller.CommentList)        // 视频评论列表

		// 社交接口
		apiRouter.POST("/relation/action/", controller.RelationAction)     // 关系操作
		apiRouter.GET("/relation/follow/list/", controller.FollowList)     // 用户关注列表
		apiRouter.GET("/relation/follower/list/", controller.FollowerList) // 用户粉丝列表
		apiRouter.GET("/relation/friend/list/", controller.FriendList)     // 用户好友列表
		apiRouter.GET("/message/chat/", controller.MessageChat)            // 聊天记录
		apiRouter.POST("/message/action/", controller.MessageAction)       // 消息操作
	}

	//err := r.Run()
	//if err != nil {
	//	return
	//}
}

// 在需要鉴权的路由上引入
func _taskMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{global.HzJwtMw.MiddlewareFunc()}
}
