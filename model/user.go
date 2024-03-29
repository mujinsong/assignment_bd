package model

import "time"

// User 是存往数据库里的用户的基本信息（以登录信息为主）
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	Salt     string    `json:"salt,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
	Name     string    `json:"name,omitempty"`
	UserFollowCount
}
type UserFollowCount struct {
	FollowCount   uint64 `json:"follow_count,omitempty"`
	FollowerCount uint64 `json:"follower_count,omitempty"`
	//IsFollow uint64 `json:"is_follow"`
}

// UserInfo 是这个用户所有的信息（视频作者信息）
type UserInfo struct {
	//User
	ID            uint64 `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   uint64 `json:"follow_count"`   // 关注总数
	FollowerCount uint64 `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

type FriendUser struct {
	UserInfo
	Message string `json:"message"` // 和该好友的最新聊天消息
	MsgType int    `json:"msgType"` // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

func (u User) TableName() string {
	return "users"
}

//func (u UserInfo) TableName() string {
//	return "users"
//}
