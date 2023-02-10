package model

/*
/douyin/comment/action/ 评论返回的结构体
*/
type Comment struct {
	ID         uint64   `json:"id,omitempty"`
	User       UserInfo `json:"user,omitempty"`
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`
}

/*
评论的表结构 comments
*/
type Comments struct {
	ID         uint64 `json:"id,omitempty"`
	UserID     uint64 `json:"user_id,omitempty"`
	VideoID    uint64 `json:"video_id,omitempty"`
	ActionType uint8  `json:"action_type,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func (c Comment) TableName() string {
	return "comments"
}
