package dao

type VideoLike struct {
	ID         uint `json:"id"`
	UserID     uint `json:"user_id"`
	VideoId    uint `json:"video_id"`
	ActionType int8 `json:"action_type"`
}

func (v VideoLike) TableName() string {
	return "likes"
}

type VideoLikeCount struct {
	VideoId   uint `json:"video_id"`
	LikeCount uint `json:"like_count"`
}
