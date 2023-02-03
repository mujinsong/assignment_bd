package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type VideoListResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户发布的视频列表
}

// Video 视频信息
type Video struct {
	//Author        UserInfo `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

// Publish 发布视频的操作 （剩下逻辑注释本方法作者补写）
func Publish(ctx context.Context, c *app.RequestContext) {

}

// PublishList 根据 user_id 查询用户 id，再查询这个用户发布的视频 （剩下逻辑注释本方法作者补写）
// TODO 这里的视频列表里的视频都是写死的，以后可以考虑用 oss 来存储，数据库里存储 URL
func PublishList(ctx context.Context, c *app.RequestContext) {
	//var user User
	//id := c.Query("user_id")
	//
	//result := config.DB.First(&user, id)
	//if result.Error != nil {
	//	fmt.Println(result.Error)        // 返回 error
	//	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	//}
	//
	//c.JSON(http.StatusOK, VideoListResponse{
	//	StatusCode: 0,
	//	StatusMsg:  nil,
	//	VideoList: []Video{
	//		{
	//			ID: 1,
	//			Author: UserInfo{
	//				FollowCount:   0,
	//				FollowerCount: 0,
	//				ID:            user.ID,
	//				IsFollow:      false,
	//				Name:          user.Username,
	//			},
	//			PlayURL:       "https://www.w3schools.com/html/movie.mp4",
	//			CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	//			FavoriteCount: 0,
	//			CommentCount:  0,
	//			IsFavorite:    false,
	//		},
	//	},
	//})
}
