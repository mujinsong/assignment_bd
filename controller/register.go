package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/dao"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	UserTableList := []dao.Users{}
	dao.DB.Find(&UserTableList, "username = ? )", username)
	if len(UserTableList) != 0 {
		SameUserNameErrorResponse := UserNameErrorResponse
		SameUserNameErrorResponse.StatusMsg = "用户名已存在"
		c.JSON(200, SameUserNameErrorResponse)
		return
	}
	// 将其插入数据库 并获取返回的信息
	dao.DB.Create(&dao.Users{Username: username, Password: password})
	SuccessLoginResponse := SuccessResponse
	SuccessLoginResponse.UserId = 1 // TODO: 这里应该是插入数据库后返回的id
	SuccessLoginResponse.Token = "testToken"
	c.JSON(200, SuccessLoginResponse)
}
