package dao

type Comment struct {
	ID         uint64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func (c Comment) TableName() string {
	return "comments"
}
