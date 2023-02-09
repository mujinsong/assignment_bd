package model

type Comment struct {
	ID         uint64  `json:"id,omitempty"`
	UserID     uint64 `json:"user_id"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func (c Comment) TableName() string {
	return "comments"
}
