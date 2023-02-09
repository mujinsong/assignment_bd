package controller

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"assignment_bd/service"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

// Register 用户注册账号（剩下逻辑注释本方法作者补写）
func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	// 注册用户到数据库（service 层的操作）
	println("注册操作")
	_, err := service.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
		return
	}
	global.HzJwtMw.LoginHandler(ctx, c)
}

// Login 用户登录
func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	_, err := service.Login(&model.Login{Username: username, Password: password})
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
		return
	}

	global.HzJwtMw.LoginHandler(ctx, c)
}

// UserInfo 获取用户信息（剩下逻辑注释本方法作者补写）
func UserInfo(ctx context.Context, c *app.RequestContext) {
	// 获取指定用户的 ID，并请求用户详细信息 UserInfo
	strID := c.Query("user_id")
	userInfo, err := service.UserInfoGetByUserID(strID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "用户信息获取失败"})
		return
	}
	c.JSON(http.StatusOK, model.UserInfoResponse{
		Response: model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  "进入个人信息页面",
		},
		UserInfo: userInfo,
	})
}
