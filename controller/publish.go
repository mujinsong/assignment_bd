package controller

import (
	"assignment_bd/model"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"strconv"
)

// Publish 发布视频的操作 （剩下逻辑注释本方法作者补写） 登录用户选择视频上传。
func Publish(ctx context.Context, c *app.RequestContext) {
	// 鉴权 待补充
	// 获取视频流 用户ID 视频标题
	data, err := c.FormFile("data")
	userId, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	log.Printf("获取到用户id:%v\n", userId)
	title := c.PostForm("title")
	log.Printf("获取到视频title:%v\n", title)
	if err != nil {
		log.Printf("获取视频流失败:%v", err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videoInfo := GetVideo(title)
	// 上传视频至数据库 并补充其余信息
	err = videoInfo.Publish(data, userId, title)
	if err != nil {
		log.Printf("Publish(data, userId) 失败：%v", err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	log.Printf("Publish(data, userId) 成功")

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// PublishList 根据 user_id 查询用户 id，再查询这个用户发布的视频 （剩下逻辑注释本方法作者补写） 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
// TODO 这里的视频列表里的视频都是写死的，以后可以考虑用 oss 来存储，数据库里存储 URL
func PublishList(ctx context.Context, c *app.RequestContext) {

}

//type VideoInfo struct {
//	ID            int64    `json:"id"`             // 视频唯一标识
//	Author        UserInfo `json:"author"`         // 视频作者信息
//	PlayURL       string   `json:"play_url"`       // 视频播放地址
//	CoverURL      string   `json:"cover_url"`      // 视频封面地址
//	FavoriteCount int64    `json:"favorite_count"` // 视频的点赞总数
//	CommentCount  int64    `json:"comment_count"`  // 视频的评论总数
//	IsFavorite    bool     `json:"is_favorite"`    // true-已点赞，false-未点赞
//	Title         string   `json:"title"`          // 视频标题
//}

// GetVideo 拼装一部分 另一部分在上传时进行
func GetVideo(title string) model.VideoInfo {
	var videoInfo model.VideoInfo
	// 视频作者信息
	//var userService service.UserInfo
	//videoInfo.Author =
	// 视频点赞总数 ： 功能未实现 默认0
	videoInfo.FavoriteCount = 0
	// 视频评论总数 ： 功能未实现 默认0
	videoInfo.CommentCount = 0
	// true-已点赞，false-未点赞
	videoInfo.IsFavorite = false
	// 视频标题
	videoInfo.Title = title
	return videoInfo
}
