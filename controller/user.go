package controller

import (
	"assignment_bd/api/backend"
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"assignment_bd/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register 用户注册账号
func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	// 注册用户到数据库
	userModel, err := service.Register(&model.Login{Username: username, Password: password})
	if err != nil {
		c.JSON(http.StatusOK, backend.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	// 生成对应 token
	tokenString := utils.RandStr(consts.TOKEN_LENGTH)
	// 返回成功并生成响应 json
	c.JSON(http.StatusOK, backend.UserLoginResponse{
		Response: backend.Response{StatusCode: 200, StatusMsg: "OK"},
		UserID:   uint64(userModel.Id),
		Token:    tokenString,
	})
}

// Login 用户登录
func Login(ctx context.Context, c *app.RequestContext) {
	//fmt.Println(c)
	username := c.Query("username")
	password := c.Query("password")
	//fmt.Println("mes:", username, password)
	// 从数据库查询用户信息
	userModel, err := service.Login(&model.Login{Username: username, Password: password})
	if err != nil {
		c.JSON(http.StatusOK, backend.Response{StatusCode: 1, StatusMsg: "用户名或密码错误"})
		return
	}
	// 生成对应 token
	tokenString := utils.RandStr(consts.TOKEN_LENGTH)
	// 返回成功并生成响应 json
	c.JSON(http.StatusOK, backend.UserLoginResponse{
		Response: backend.Response{StatusCode: 200, StatusMsg: "OK"},
		UserID:   uint64(userModel.Id),
		Token:    tokenString,
	})
}

// UserInfo 获取用户信息
func UserInfo(c *gin.Context) {
	// 获取指定用户的 ID
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, backend.Response{StatusCode: 1, StatusMsg: "request is invalid"})
		return
	}
	userModel, err := service.UserInfoGetByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, backend.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	// 获取当前用户的 ID
	//viewerID := c.GetUint64("UserID")
	// 查询当前用户是否关注指定用户
	//isFollow, err := service.GetFollowStatus(viewerID, userID)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, backend.Response{StatusCode: 1, StatusMsg: err.Error()})
	//	return
	//}
	// 返回成功并生成响应 json
	c.JSON(http.StatusOK, backend.UserResponse{
		Response: backend.Response{StatusCode: 0, StatusMsg: "OK"},
		User: model.User{
			Id:       uint(userID),
			Username: userModel.Username,
			//FollowCount:   userModel.FollowCount,
			//FollowerCount: userModel.FollowerCount,
			//TotalFavorited: userModel.TotalFavorited,
			//FavoriteCount:  userModel.FavoriteCount,
			//IsFollow: isFollow,
		},
	})
}
