package model

type VideoLike struct {
	Id         int  `json:"id"`
	UserId     int  `json:"user_id"`
	VideoId    int  `json:"video_id"`
	ActionType int8 `json:"action_type"`
}

func (v VideoLike) TableName() string {
	return "likes"
}

type VideoLikeCount struct {
	VideoId   int `json:"video_id"`
	LikeCount int `json:"like_count"`
}
