package dao

type Follow struct {
	ID         uint `json:"id"`
	UserID     uint `json:"user_id"`
	FollowerID uint `json:"follower_id"`
	ActionType int  `json:"action_type"`
}

func (f Follow) TableName() string {
	return "follows"
}
