package controller

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}
type FeedResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	NextTime   int64   `json:"next_time,omitempty"`
	VideoList  []Video `json:"video_list,omitempty"`
}

func Feed(c *gin.Context) {
	c.JSON(200, TestFeedResponse)

}
