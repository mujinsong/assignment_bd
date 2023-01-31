package model

import "time"

// Video 对应的是数据库中的 video 结构，用来存储
type Video struct {
	Id        int       `json:"id,omitempty"`
	UserId    int       `json:"user_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	PlayUrl   string    `json:"play_url,omitempty"`
	CoverUrl  string    `json:"cover_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// VideoInfo 视频信息，基本包含了视频的所有信息，不能直接通过数据库获取，需要各个字段拼装获取
type VideoInfo struct {
	Author        UserInfo `json:"author"`         // 视频作者信息
	CommentCount  int      `json:"comment_count"`  // 视频的评论总数
	CoverURL      string   `json:"cover_url"`      // 视频封面地址
	FavoriteCount int      `json:"favorite_count"` // 视频的点赞总数
	ID            int      `json:"id"`             // 视频唯一标识
	IsFavorite    bool     `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string   `json:"play_url"`       // 视频播放地址
	Title         string   `json:"title"`          // 视频标题
}

type VideoCommentCount struct {
	VideoId int
	Count   int
}

func (receiver Video) TableName() string {
	return "videos"
}
