package controller

import (
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	//"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"strconv"
)

// FavoriteAction 对视频点赞和取消点赞的操作
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	payload, exist := c.Get(consts.IdentityKey)
	if !exist {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "操作失败"})
		return
	}
	uid := payload.(model.User).ID
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

	err = service.Like(uid, uint64(videoID), uint8(actionType))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "操作失败"})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: consts.STATUS_SUCCESS,
		StatusMsg:  "Success"})
}

// func FavoriteList (
// FavoriteList 从数据库中查询当前用户，并查询当前用户点赞过的视频（剩下逻辑注释本方法作者补写）
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	strID := c.Query("user_id")
	id, _ := strconv.Atoi(strID)
	claims, _ := c.Get(consts.IdentityKey)
	masterID := claims.(model.User).ID
	videoIDList, err := service.GetLikeVideoIDListByUserID(uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "操作失败"})
		return
	}
	//fmt.Println("len:", len(*videoIDList))
	if len(*videoIDList) == 0 {
		c.JSON(http.StatusOK, utils.H{
			"status_code": consts.STATUS_SUCCESS,
			"status_msg":  "Success",
			"video_list":  nil,
		})
		return
	}

	videoList, err := service.GetVideoListByIDs(videoIDList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "操作失败"})
		return
	}
	//fmt.Println("用户", masterID, "查看用户", id, "的点赞视频列表", "其长度为", len(*videoList))
	length := len(*videoList)
	response := make([]model.VideoInfo, length)
	for i := 0; i < length; i++ {
		response[i].ID = (*videoList)[i].ID
		response[i].PlayUrl = (*videoList)[i].PlayUrl
		response[i].Title = (*videoList)[i].Title
		response[i].CoverUrl = (*videoList)[i].CoverUrl
		response[i].Author = service.UserInfoGetByUserID((*videoList)[i].UserID, masterID)
		response[i].FavoriteCount = (*videoList)[i].FavoriteCount
		response[i].CommentCount = (*videoList)[i].CommentCount
		response[i].IsFavorite = service.IsFavorite(masterID, (*videoList)[i].ID)
	}

	c.JSON(http.StatusOK, model.VideoListResponse{
		Response: model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  consts.MsgFlags[consts.SUCCESS],
		},
		VideoList: response,
	})
}
