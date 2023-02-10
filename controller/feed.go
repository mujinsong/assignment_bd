package controller

import (
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"assignment_bd/utils"
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

/*
	项目本该部署到服务器上利用nginx提供视频流，这里为了方便测试，
*/

// Feed 视频流接口（给客户端推送短视频）（剩下逻辑注释本方法作者补写）
// TODO 完善视频推送的逻辑 和 查询优化
func Feed(ctx context.Context, c *app.RequestContext) {
	// 查询数据库中的视频信息
	lasttime := c.Query("latest_time")
	println(lasttime)
	uid, _ := utils.GetUid(c)
	var feedResponse model.FeedResponse
	feedResponse.Response.StatusCode = consts.STATUS_SUCCESS
	feedResponse.Response.StatusMsg = "获取视频列表成功"
	feedResponse.NextTime = uint64(time.Now().Unix())

	videolist := service.FindVideos()
	// 通过视频列表中的userId获取用户信息
	for _, video := range videolist {
		feedResponse.VideoList = append(feedResponse.VideoList,
			model.VideoInfo{
				ID:            video.ID,
				Author:        service.UserInfoGetByUserID(video.UserID, uid),
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    service.IsFavorite(uid, video.ID),
				Title:         video.Title,
			},
		)
	}

	// 首先获取视频列表 然后通过视频列表中的userId获取用户信息
	c.JSON(http.StatusOK, feedResponse)

}
