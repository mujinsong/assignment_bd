package controller

import (
	"assignment_bd/config"
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	utils2 "assignment_bd/utils"
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// Publish 发布视频的操作 （剩下逻辑注释本方法作者补写） 登录用户选择视频上传。
func Publish(ctx context.Context, c *app.RequestContext) {
	// 解析参数
	data, err := c.FormFile("data")
	if err != nil {
		log.Printf("获取视频流失败:%v", err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	user, _ := c.Get(consts.IdentityKey)
	userID := user.(model.User).ID
	log.Printf("获取到用户id:%v\n", userID)
	title := c.PostForm("title")
	log.Printf("获取到视频title:%v\n", title)

	localPath := "static/video/" + utils2.RandVideoName(data.Filename)
	video := model.Video{
		Title:   title,
		UserID:  userID,
		PlayUrl: config.UseServer + config.Port + strings.Replace(localPath, "static", "", 1),
	}
	// 保存视频到本地
	c.SaveUploadedFile(data, localPath)

	// 异步执行 为了防止视频还没有上传完就生成封面 如果5秒后还没有上传完成就会报错
	time.AfterFunc(5*time.Second, func() {
		// 生成视频封面
		video.CoverUrl = config.UseServer + config.Port + strings.Replace(utils2.GetSnapshot(localPath), "static", "", 1)
		// 保存视频信息到数据库
		service.PublishVideo(&video)
	})
	c.JSON(http.StatusOK, model.Response{
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
	//println("用户", user.Name, "该用户的视频列表数目为：", len(videoJOSN))
	if uid == userID {
		//println("查询的是自己的视频列表")
	}

	c.JSON(http.StatusOK, model.VideoListResponse{
		Response: model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  consts.MsgFlags[consts.SUCCESS],
		},
		VideoList: videoJOSN,
	})
}
