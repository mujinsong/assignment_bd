package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/config"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {

}

// GetFavoriteList all users have same favorite video list
func GetFavoriteList(c *gin.Context) {
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
