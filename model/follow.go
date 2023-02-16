package model

type Follow struct {
	ID         uint64 `json:"id"`
	UserID     uint64 `json:"user_id"`
	FollowerID uint64 `json:"follower_id"`
	ActionType uint8  `json:"action_type"`
}

func (f Follow) TableName() string {
	return "follows"
}
