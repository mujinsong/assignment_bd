package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/dao"
)

type LoginResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}

var UserNameErrorResponse = LoginResponse{
	StatusCode: 1,
	StatusMsg:  "用户名不存在 请注册",
	UserId:     0,
	Token:      "",
}

var PasswordErrorResponse = LoginResponse{
	StatusCode: 2,
	StatusMsg:  "密码错误",
	UserId:     0,
	Token:      "",
}

var SuccessResponse = LoginResponse{
	StatusCode: 0,
	StatusMsg:  "登录成功",
	UserId:     0,
	Token:      "testToken",
}

// TODO: 写得有点乱，你们可以自己优化一下
func Login(c *gin.Context) {
	// 解析url参数 ?username=xxx&password=xxx
	username := c.Query("username")
	password := c.Query("password")

	var UserTableList []dao.Users
	dao.DB.Find(&UserTableList, "username = ? AND password = ?", username, password)
	if len(UserTableList) == 0 {
		c.JSON(200, UserNameErrorResponse)
		return
	}
	if UserTableList[0].Password != password {
		c.JSON(200, PasswordErrorResponse)
		return
	}
	c.JSON(200, SuccessResponse)
}
