package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/config"
)

type VideoListResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户发布的视频列表
}

// Video 视频信息
type Video struct {
	Author        UserInfo `json:"author"`         // 视频作者信息
	CommentCount  int64    `json:"comment_count"`  // 视频的评论总数
	CoverURL      string   `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64    `json:"favorite_count"` // 视频的点赞总数
	ID            int64    `json:"id"`             // 视频唯一标识
	IsFavorite    bool     `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string   `json:"play_url"`       // 视频播放地址
	Title         string   `json:"title"`          // 视频标题
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	//token := c.PostForm("token")

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	var user User
	id := c.Query("user_id")

	result := config.DB.First(&user, id)
	if result.Error != nil {
		fmt.Println(result.Error)        // 返回 error
		fmt.Println(result.RowsAffected) // 返回插入记录的条数
	}

	c.JSON(http.StatusOK, VideoListResponse{
		StatusCode: 1,
		StatusMsg:  nil,
		VideoList: []Video{
			{
				ID: 1,
				Author: UserInfo{
					FollowCount:   0,
					FollowerCount: 0,
					ID:            user.ID,
					IsFollow:      false,
					Name:          user.Username,
				},
				PlayURL:       "https://www.w3schools.com/html/movie.mp4",
				CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
				FavoriteCount: 0,
				CommentCount:  0,
				IsFavorite:    false,
			},
		},
	})

}
