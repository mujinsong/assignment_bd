package controller

import (
	"assignment_bd/api/backend"
	"assignment_bd/dao"
	"assignment_bd/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var videolist []dao.Video
	var videoId []uint64
	err := service.GetVideoListByIDs(&videolist, videoId)
	if err != nil {
		fmt.Println(errors.New("获取视频列表失败"))
		return
	}
	c.JSON(http.StatusOK, backend.FeedResponse{
		Response:  backend.Response{StatusCode: 0},
		VideoList: videolist,
		NextTime:  time.Now().Unix(),
	})
}
