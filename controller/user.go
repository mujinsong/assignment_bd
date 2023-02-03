package controller

import (
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"assignment_bd/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// Register 用户注册账号（剩下逻辑注释本方法作者补写）
func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	// 注册用户到数据库（service 层的操作）
	userModel, err := service.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
		return
	}

	// 生成对应 token
	tokenString := utils.RandStr(consts.TOKEN_LENGTH)

	// 返回成功并生成响应 json
	c.JSON(http.StatusOK, model.UserLoginResponse{
		Response: model.Response{StatusCode: consts.STATUS_SUCCESS, StatusMsg: "已经注册成功"},
		UserID:   userModel.Id,
		Token:    tokenString,
	})
}

// Login 用户登录（剩下逻辑注释本方法作者补写）
func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	// 从数据库查询用户信息
	userModel, err := service.Login(&model.Login{Username: username, Password: password})
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: "用户名或密码错误"})
		return
	}

	// 生成对应 token
	tokenString := utils.RandStr(consts.TOKEN_LENGTH)

	// 返回成功并生成响应 json
	c.JSON(http.StatusOK, model.UserLoginResponse{
		Response: model.Response{StatusCode: consts.STATUS_SUCCESS, StatusMsg: "登录成功"},
		UserID:   userModel.Id,
		Token:    tokenString,
	})
}

// UserInfo 获取用户信息（剩下逻辑注释本方法作者补写）
// todo 这里返回的 UserInfo 先写死了，因为我们其它例如获取用户关注数的功能还没写
func UserInfo(ctx context.Context, c *app.RequestContext) {
	// 获取指定用户的 ID，并请求用户详细信息 UserInfo
	userModel, err := service.UserInfoGetByUserID(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "用户信息获取失败"})
		return
	}

	//var users []model.UserInfo
	//global.DB.Where("id = ?", 6).Find(&users)
	//fmt.Println(users)

	// 返回成功并生成响应 json
	// 这里的数据写死了
	c.JSON(http.StatusOK, model.UserInfoResponse{
		Response: model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  "进入个人信息页面",
		},
		UserInfo: &model.UserInfo{
			//User: *userModel,
			ID:            userModel.Id,
			Name:          userModel.Username,
			FollowCount:   3,
			FollowerCount: 3,
			IsFollow:      false,
			Avatar:        "http://img.panker916.space/32bf3ac8c3e1420bafee3ab84cb5f17e",
		},
	})
}
