package controller

import (
	"assignment_bd/config"
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	utils2 "assignment_bd/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strings"
)

// Publish 发布视频的操作 （剩下逻辑注释本方法作者补写） 登录用户选择视频上传。
func Publish(ctx context.Context, c *app.RequestContext) {
	// 解析参数
	data, err := c.FormFile("data")
	if err != nil {
		log.Printf("获取视频流失败:%v", err)
		c.JSON(consts.SUCCESS, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	user, _ := c.Get(consts.IdentityKey)
	userID := user.(model.User).ID
	title := c.PostForm("title")
	videoPath := "static/video/" + utils2.RandVideoName(data.Filename)
	video := model.Video{
		Title:   title,
		UserID:  userID,
		PlayUrl: config.UseServer + config.Port + strings.Replace(videoPath, "static", "", 1),
	}
	// 保存视频到本地
	c.SaveUploadedFile(data, videoPath)
	video.CoverUrl = config.UseServer + config.Port + strings.Replace(utils2.GetSnapshot(videoPath), "static", "", 1)
	service.PublishVideo(&video)
	c.JSON(consts.SUCCESS, model.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

/* PublishList 查询用户发布的视频列表
 */
func PublishList(ctx context.Context, c *app.RequestContext) {
	// 解析参数
	//uid 是当前登录用户的id
	//userID 是要查询的用户的id
	uid, _ := utils2.GetUid(c)
	userID := utils2.StrToUint64(c.Query("user_id"))
	user := service.UserInfoGetByUserID(userID, uid)
	videoList, err := service.GetPublishList(userID)
	if err != nil {
		return
	}
	lens := len(videoList)
	videoJOSN := make([]model.VideoInfo, lens)
	for i := 0; i < lens; i++ {
		videoJOSN[i].ID = videoList[i].ID
		videoJOSN[i].Author = user

		videoJOSN[i].CoverUrl = videoList[i].CoverUrl
		videoJOSN[i].Title = videoList[i].Title
		videoJOSN[i].PlayUrl = videoList[i].PlayUrl

		videoJOSN[i].IsFavorite = service.IsFavorite(uid, videoList[i].ID)
		videoJOSN[i].FavoriteCount = videoList[i].FavoriteCount
		videoJOSN[i].CommentCount = videoList[i].CommentCount
	}
	c.JSON(consts.SUCCESS, model.VideoListResponse{
		Response: model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  consts.MsgFlags[consts.SUCCESS],
		},
		VideoList: videoJOSN,
	})
}
