package controller

import (
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// Deprecated: Like 点赞操作，见favorite.go
func Like(ctx context.Context, c *app.RequestContext) {
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

	err = service.Like(uid, uint64(videoID), int32(actionType))
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

// Deprecated: LikeList 点赞列表，见favorite.go
func LikeList(ctx context.Context, c *app.RequestContext) {
}
