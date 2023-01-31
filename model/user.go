package model

import "time"

type User struct {
	Id       uint      `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	Salt     string    `json:"salt,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}

func (u User) TableName() string {
	return "users"
}
