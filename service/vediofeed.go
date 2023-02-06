package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

// GetFeedVideosAndAuthors 获取推送视频以及其作者并返回视频数
func GetFeedVideosAndAuthors(videoList *[]model.Video, users *[]model.User, LatestTime time.Time, MaxNumVideo int) (int, error) {

	result := global.DB.Model(&model.Video{}).Select("id,user_id").Where("created_at > ?", LatestTime).Order("created_at DESC").Limit(MaxNumVideo).Find(videoList)
	if result.RowsAffected == 0 {
		return 0, errors.New("没视频")
	}

	numVideos := len(*videoList)

	// 批量或者视频作者
	userIDList := make([]int64, numVideos)
	for i, video := range *videoList {
		userIDList[i] = video.UserId
	}
	result = global.DB.Model(&model.User{}).Where("user_id IN ?", userIDList).Find(users)
	if result.RowsAffected == 0 {
		return 0, errors.New("没作者")
	}
	return numVideos, nil
}

func GetVideoListByIDs(ctx *gin.Context, videoList *[]model.Video, videoIDs []int) error {
	//var uniqueVideoList []model.Video
	//result := global.DB.Where("video_id in ?", videoIDs).Find(&uniqueVideoList)
	//if result.Error != nil {
	//	return result.Error
	//}
	//numVideos := result.RowsAffected
	//*videoList = make([]model.Video, 0, numVideos)
	//mapVideoIDToVideo := make(map[int]model.Video, numVideos)
	//for _, video := range uniqueVideoList {
	//	mapVideoIDToVideo[video.Id] = video
	//}
	//// 查询like_count与comment_count
	//var commentCountList []int
	//var likeCountList []int
	//if err := GetCommentCountListByVideoIDList(ctx, videoIDs, &commentCountList); err != nil {
	//	return err
	//}
	////todo
	//err := GetLikeCountListByVideoIDList(videoIDs, &likeCountList)
	//if err != nil {
	//	return err
	//}
	//for i, videoID := range videoIDs {
	//	tmpVideo := mapVideoIDToVideo[videoID]
	//	tmpVideo.FavoriteCount = likeCountList[i]
	//	tmpVideo.FavoriteCount = commentCountList[i]
	//	*videoList = append(*videoList, tmpVideo)
	//}
	return nil
}

// GetVideoListByUserID 得到用户发表过的视频列表,返回视频数
func GetVideoListByUserID(userID int64, videoList *[]model.Video) (int, error) {
	result := global.DB.Where("user_id = ?", userID).Find(&videoList).Order("created_time DESC")
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, nil
	}
	numVideos := int(result.RowsAffected)

	return numVideos, nil

}

// PublishVideo 将用户上传的视频信息写入数据库(old)
//func PublishVideo(ctx *gin.Context,userID int64, videoID int64, videoName string, coverName string, title string) error {
//	video := model.Video{
//		Id:       videoID,
//		Title:    title,
//		PlayUrl:  videoName,
//		CoverUrl: coverName,
//		//FavoriteCount : 0,
//		//CommentCount : 0,
//		Author: model.User{
//			Id: int(userID),
//		},
//		CreatedAt: time.Now(),
//	}
//	if global.DB.WithContext(ctx).Create(&video).Error != nil {
//		return errors.New("video表插入失败")
//	}
//	return nil
//}
