package model

import "time"

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	CreatTime time.Time `json:"creat_time"`
}
