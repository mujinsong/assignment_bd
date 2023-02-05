package dao

import "time"

type Video struct {
	Id            uint      `json:"id,omitempty"`
	UserId        uint      `json:"user_id,omitempty"`
	Title         string    `json:"title,omitempty"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount uint      `json:"favorite_count,omitempty"`
	CommentCount  uint      `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type VideoCommentCount struct {
	VideoId uint
	Count   uint
}

func (receiver Video) TableName() string {
	return "videos"
}
