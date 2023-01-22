package dao

import "time"

type Video struct {
	Id            uint64    `json:"id,omitempty"`
	Author        User      `json:"author,omitempty"`
	Title         string    `json:"title,omitempty"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	CreatedAt     time.Time `json:"create_at,omitempty"`
}
type VideoCommentCount struct {
	VideoId uint64
	Count   int64
}

func (receiver Video) TableName() string {
	return "videos"
}
