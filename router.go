package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	// 基础
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.GetUserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// ToDo 下面的接口先不写的，先把基础接口写好
	// 互动
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.GetFavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)
	//// 社交
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.GetFollowing)
	//apiRouter.GET("/relation/follower/list", controller.GetFollowers)

	r.Run()
}
