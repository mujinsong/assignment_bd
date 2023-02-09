package model

type VideoLike struct {
	ID         uint64 `json:"id"`
	UserId     uint64 `json:"user_id"`
	VideoId    uint64 `json:"video_id"`
	ActionType uint8  `json:"action_type"`
}

func (v VideoLike) TableName() string {
	return "likes"
}

type VideoLikeCount struct {
	VideoId   uint64 `json:"video_id"`
	LikeCount uint64 `json:"like_count"`
}
