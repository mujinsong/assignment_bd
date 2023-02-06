package backend

type User struct {
	Id            uint   `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type LoginORRegisterResponse struct {
	Response
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}
