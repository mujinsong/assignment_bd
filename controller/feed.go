package controller

import (
	"assignment_bd/api/backend"
	"assignment_bd/consts"
	"assignment_bd/dao"
	"assignment_bd/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"net/http"
	"time"
)

//// Feed same demo video list for every request
//func Feed(c *gin.Context) {
//	var videolist []dao.Video
//	var ctx gin.Context
//	var videoId []uint
//	err := service.GetVideoListByIDs(&ctx, &videolist, videoId)
//	if err != nil {
//		fmt.Println(errors.New("获取视频列表失败"))
//		return
//	}
//	c.JSON(http.StatusOK, backend.FeedResponse{
//		Response:  backend.Response{StatusCode: 0},
//		VideoList: videolist,
//		NextTime:  time.Now().Unix(),
//	})
//}

// Feed 视频流接口（给客户端推送短视频）
func Feed(c *gin.Context) {
	// 不传latest_time默认为当前时间
	var CurrentTime = time.Now()
	var CurrentTimeStr = utils.ToString(CurrentTime)
	var LatestTimeStr = c.DefaultQuery("latest_time", CurrentTimeStr)
	LatestTime, err := time.Parse(time.RFC3339, LatestTimeStr)
	if err != nil {
		// 无法解析latest_time
		c.JSON(http.StatusBadRequest, backend.Response{StatusCode: 1, StatusMsg: "parameter latest_time is wrong"})
		return
	}
	// 得到本次要返回的视频以及其作者
	var videoList []dao.Video
	var userList []dao.User
	numVideos, err := service.GetFeedVideosAndAuthors(&videoList, &userList, LatestTime, consts.GET_VIDEO_NUM)

	if err != nil {
		// 访问数据库出错
		c.JSON(http.StatusInternalServerError, backend.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if numVideos == 0 {
		// 没有满足条件的视频 使用当前时间再获取一遍
		numVideos, _ = service.GetFeedVideosAndAuthors(&videoList, &userList, CurrentTime, consts.GET_VIDEO_NUM)
		if numVideos == 0 {
			// 后端没有视频了
			c.JSON(http.StatusOK, backend.FeedResponse{
				Response:  backend.Response{StatusCode: 0},
				VideoList: nil,
				NextTime:  CurrentTime, // 没有视频可刷时返回当前时间
			})
			return
		}
	}

	var (
		videoJsonList = make([]dao.Video, 0, numVideos)
		videoJson     dao.Video
		user          dao.User
		userJson      dao.User
		//isFavoriteList []bool
		//isFollowList   []bool
		//isLogged       = false // 用户是否传入了合法有效的token（是否登录）
	)

	//var userID uint64
	//// 判断传入的token是否合法，用户是否存在
	//if token := c.Query("token"); token != "" {
	//	claims, err := util.ParseToken(token)
	//	if err == nil {
	//		// token合法
	//		userID = claims.UserID
	//		isLogged = true
	//	}
	//}
	//
	//if isLogged {
	//	// 当用户登录时 批量获取用户是否点赞了列表中的视频以及是否关注了视频的作者
	//	videoIDList := make([]uint64, numVideos)
	//	userIDList := make([]uint64, numVideos)
	//	for i, video := range videoList {
	//		videoIDList[i] = video.VideoID
	//		userIDList[i] = video.AuthorID
	//	}
	//	// 批量获取用户是否用视频点赞
	//	isFavoriteList, err = service.GetFavoriteStatusList(userID, videoIDList)
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
	//		return
	//	}
	//	// 批量获取用户是否关注作者
	//	isFollowList, err = service.GetFollowStatusList(userID, userIDList)
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
	//		return
	//	}
	//}

	// 未登录时默认为未关注未点赞
	var isFavorite = false
	var isFollow = false

	for i, video := range videoList {
		//if isLogged {
		//	// 当用户登录时，判断是否关注当前作者
		//	isFollow = isFollowList[i]
		//	isFavorite = isFavoriteList[i]
		//}

		// 二次确认返回的视频与封面是服务器存在的
		//VideoLocation := filepath.Join(global.VIDEO_ADDR, video.PlayName)
		//if _, err = os.Stat(VideoLocation); err != nil {
		//	continue
		//}
		//CoverLocation := filepath.Join(global.COVER_ADDR, video.CoverName)
		//if _, err = os.Stat(CoverLocation); err != nil {
		//	continue
		//}
		// 填充JSON返回值
		user = userList[i]
		userJson.Id = user.Id
		userJson.Username = user.Username
		userJson.FollowCount = user.FollowCount
		userJson.FollowerCount = user.FollowerCount
		//userJson.TotalFavorited = user.TotalFavorited
		//userJson.FavoriteCount = user.FavoriteCount
		userJson.IsFollow = isFollow

		videoJson.Id = video.Id
		videoJson.UserId = userJson.Id
		videoJson.PlayUrl = "http://" + c.Request.Host + "/static/video/" + video.PlayUrl
		videoJson.CoverUrl = "http://" + c.Request.Host + "/static/cover/" + video.CoverUrl
		videoJson.FavoriteCount = video.FavoriteCount
		videoJson.CommentCount = video.CommentCount
		videoJson.Title = video.Title
		videoJson.IsFavorite = isFavorite

		videoJsonList = append(videoJsonList, videoJson)
	}

	//本次返回的视频中发布最早的时间
	nextTime := videoList[numVideos-1].CreatedAt

	c.JSON(http.StatusOK, backend.FeedResponse{
		Response:  backend.Response{StatusCode: 0},
		VideoList: videoJsonList,
		NextTime:  nextTime,
	})
}
