package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// FavoriteAction 对视频点赞和取消点赞的操作（剩下逻辑注释本方法作者补写）
func FavoriteAction(ctx context.Context, c *app.RequestContext) {

}

// FavoriteList 从数据库中查询当前用户，并查询当前用户点赞过的视频（剩下逻辑注释本方法作者补写）
// TODO 这里的视频列表里的视频也都是写死的，以后可以考虑用 oss 来存储，数据库里存储 URL
func FavoriteList(ctx context.Context, c *app.RequestContext) {
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
	//	StatusCode: 1,
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
