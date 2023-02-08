package controller

import (
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

// FavoriteAction 对视频点赞和取消点赞的操作
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	//videoIdq := c.Query("video_id")
	//actionTypeq := c.Query("action_type")
	//
	//uid, err := utils.GetUid(c)
	//if err != nil {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
	//	return
	//}
	//
	//if videoIdq == "" || actionTypeq == "" {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: "参数错误"})
	//	return
	//}
	//
	//videoId, err := strconv.ParseInt(videoIdq, 10, 64)
	//if err != nil {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
	//	return
	//}
	//
	//actionType, err := strconv.ParseInt(actionTypeq, 10, 32)
	//if err != nil {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
	//	return
	//}
	//
	//if err := service.Like(uid, videoId, int32(actionType)); err != nil {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, model.Response{StatusCode: consts.STATUS_SUCCESS, StatusMsg: "成功"})
	payload, exist := c.Get(consts.IdentityKey)
	if !exist {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "操作失败"})
		return
	}
	uid := payload.(model.User).Id
	//fmt.Println(uid)
	videoIDStr := c.Query("video_id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		return
	}
	actionType := 0
	temp := c.Query("action_type")
	if temp == "1" {
		actionType = 1
	} else {
		actionType = 2
	}
	err = service.Like(uid, int64(videoID), int32(actionType))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "操作失败"})
		return
	}
	c.JSON(http.StatusInternalServerError, model.Response{
		StatusCode: consts.STATUS_SUCCESS,
		StatusMsg:  "Success"})
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
