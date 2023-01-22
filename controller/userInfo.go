package controller

import (
	"github.com/gin-gonic/gin"
)

/*
用户信息
*/
type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserInfoResponse struct {
	StatusCode int      `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	User       User `json:"user"`
}

func UserInfo(c *gin.Context) {
	// 解析url参数 ?user_id=xxx&token=xxx
	//user_id := c.Query("user_id")
	//token := c.Query("token")

	c.JSON(200, TestUserInfoResponse)
	// 从数据库中查询用户信息
	//var UserTableList []dao.Users
	//dao.DB.Find(&UserTableList, "id = ? AND token = ?", userId, token)
	//if len(UserTableList) == 0 {
	//	c.JSON(200, UserNameErrorResponse)
	//	return
	//}
	//if UserTableList[0].Token != token {
	//	c.JSON(200, PasswordErrorResponse)
	//	return
	//}
}
