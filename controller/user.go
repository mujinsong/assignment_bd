package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tiktok/config"
	"time"
)

// User 是定义的用户的数据结构
type User struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Salt     string    `json:"salt"`
	CreateAt time.Time `json:"create_at"`
}

// UserResponse 是用户注册或者登录之后返回的数据结构（两者是同一种数据结构）
type UserResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

// UserInfoResponse 是
type UserInfoResponse struct {
	StatusCode int64     `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string   `json:"status_msg"`  // 返回状态描述
	UserInfo   *UserInfo `json:"user"`        // 用户信息
}

// UserInfo 是这个用户所有的信息（视频作者信息）
type UserInfo struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

func Feed(c *gin.Context) {
	c.JSON(200, "TestFeedResponse")
}

func Register(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")

	username := c.Query("username")
	password := c.Query("password")

	user := User{
		Username: username,
		Password: password,
		Salt:     "2333",
		CreateAt: time.Now(),
	}

	result := config.DB.Create(&user) // 通过数据的指针来创建

	code := int64(0)
	msg := "用户注册成功"
	if result.Error != nil {
		code = 1
		msg = "用户注册失败"
		fmt.Println(result.Error)        // 返回 error
		fmt.Println(result.RowsAffected) // 返回插入记录的条数
	}

	c.JSON(200, UserResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Token:      "ok",
		UserID:     user.ID,
	})
}

func Login(c *gin.Context) {
	var user User
	username := c.PostForm("username")
	password := c.PostForm("password")

	username = c.Query("username")
	password = c.Query("password")

	result := config.DB.
		Where("username = ?", username).
		Where("password = ?", password).First(&user)

	fmt.Println(user)

	code := int64(0)
	msg := "用户登录成功"
	if result.Error != nil {
		code = 1
		msg = "用户登录失败"
		fmt.Println(result.Error)        // 返回 error
		fmt.Println(result.RowsAffected) // 返回插入记录的条数
	}

	c.JSON(200, UserResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Token:      "ok",
		UserID:     user.ID,
	})
}

// GetUserInfo 用来获取用户的详细信息
// TODO 用户的的关注信息等等都是写死的，因为还没有实现关注功能
func GetUserInfo(c *gin.Context) {
	var user User
	id := c.Query("user_id")

	result := config.DB.First(&user, id)

	code := int64(0)
	msg := "用户信息获取成功"
	if result.Error != nil {
		msg = "用户信息获取失败"
		fmt.Println(result.Error)        // 返回 error
		fmt.Println(result.RowsAffected) // 返回插入记录的条数
	}

	c.JSON(200, UserInfoResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		UserInfo: &UserInfo{
			FollowCount:   0,
			FollowerCount: 0,
			ID:            user.ID,
			IsFollow:      false,
			Name:          user.Username,
		},
	})
}
