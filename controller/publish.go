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
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// Publish 发布视频的操作 （剩下逻辑注释本方法作者补写） 登录用户选择视频上传。
func Publish(ctx context.Context, c *app.RequestContext) {
	// 鉴权 待补充
	// 获取视频流 用户ID 视频标题
	//claims := jwt.ExtractClaims(ctx, c)
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
		PlayUrl: config.Server + config.Port + strings.Replace(localPath, "static", "", 1),
	}
	// 保存视频到本地
	c.SaveUploadedFile(data, localPath)

	// 异步执行 为了防止视频还没有上传完就生成封面 如果5秒后还没有上传完成就会报错
	time.AfterFunc(5*time.Second, func() {
		// 生成视频封面
		video.CoverUrl = config.Server + config.Port + strings.Replace(utils2.GetSnapshot(localPath), "static", "", 1)
		// 保存视频信息到数据库
		service.PublishVideo(&video)
	})

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	//claims := jwt.ExtractClaims(ctx, c)
	//当前用户ID
	//userID := uint64(claims[middleware.IdentityKey].(float64))
	//查询的ID
	uid, _ := utils2.GetUid(c)
	var videoList []model.Video
	_, err := service.GetVideoListByUserID(uid, &videoList)
	if err != nil {
		return
	}
	userID := utils2.StrToUint64(c.Query("user_id"))
	user := service.UserInfoGetByUserID(userID, uid)
	videoJOSN := make([]model.VideoInfo, len(videoList))
	//todo
	for i := 0; i < len(videoList); i++ {
		videoJOSN[i].ID = videoList[i].ID
		videoJOSN[i].CoverUrl = videoList[i].CoverUrl
		videoJOSN[i].Title = videoList[i].Title
		videoJOSN[i].PlayUrl = videoList[i].PlayUrl
		videoJOSN[i].Author = user
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "OK",
		"video_list":  videoJOSN,
	})
}
