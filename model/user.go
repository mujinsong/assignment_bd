package model

import "time"

// User 是存往数据库里的用户的基本信息（以登录信息为主）
type User struct {
	Id       int64     `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	Salt     string    `json:"salt,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}

// UserInfo 是这个用户所有的信息（视频作者信息）
type UserInfo struct {
	ID            int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注

}

func (u User) TableName() string {
	return "users"
}