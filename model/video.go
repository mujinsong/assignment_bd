package model

import (
	"time"
)

// Video 对应的是数据库中的 video 结构，用来存储
type Video struct {
	ID            uint64    `json:"id,omitempty"`
	UserID        uint64    `json:"user_id,omitempty"`
	Title         string    `json:"title,omitempty"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	FavoriteCount uint64    `json:"favorite_count"`
	CommentCount  uint64    `json:"comment_count"`
}

// VideoInfo 视频信息，基本包含了视频的所有信息，不能直接通过数据库获取，需要各个字段拼装获取
type VideoInfo struct {
	ID            uint64   `json:"id"`             // 视频唯一标识
	Author        UserInfo `json:"author"`         // 视频作者信息
	PlayUrl       string   `json:"play_url"`       // 视频播放地址
	CoverUrl      string   `json:"cover_url"`      // 视频封面地址
	FavoriteCount uint64   `json:"favorite_count"` // 视频的点赞总数
	CommentCount  uint64   `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool     `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string   `json:"title"`          // 视频标题
}

// VideoCommentCount 评论数
type VideoCommentCount struct {
	VideoId uint64 `json:"video_id"`
	Count   uint64 `json:"count"`
}

func (receiver Video) TableName() string {
	return "videos"
}
