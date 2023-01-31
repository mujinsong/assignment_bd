package backend

import "assignment_bd/dao"

// UserLoginResponse 用户登录响应结构体
type UserLoginResponse struct {
	Response
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

// UserResponse 用户信息响应结构体
type UserResponse struct {
	Response
	User dao.User `json:"user"`
}
