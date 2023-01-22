package dao

import "time"

type User struct {
	Id            uint      `json:"id,omitempty"`
	Username      string    `json:"username,omitempty"`
	Password      string    `json:"password,omitempty"`
	Salt          string    `json:"salt,omitempty"`
	FollowCount   int64     `json:"follow_count,omitempty"`
	FollowerCount int64     `json:"follower_count,omitempty"`
	IsFollow      bool      `json:"is_follow,omitempty"`
	CreatTime     time.Time `json:"creat_time,omitempty"`
}

func (u User) TableName() string {
	return "users"
}
