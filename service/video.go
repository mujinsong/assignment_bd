package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
	"errors"
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
	userIDList := make([]uint64, numVideos)
	for i, video := range *videoList {
		userIDList[i] = video.UserId
	}
	result = global.DB.Model(&model.User{}).Where("user_id IN ?", userIDList).Find(users)
	if result.RowsAffected == 0 {
		return 0, errors.New("没作者")
	}
	return numVideos, nil
}

func GetVideoListByIDs(videoIDs *[]uint64) (*[]model.Video, error) {
	var uniqueVideoList []model.Video
	result := global.DB.Where("id in ?", *videoIDs).Find(&uniqueVideoList).Order("id DESC")
	if result.Error != nil {
		return nil, result.Error
	}
	return &uniqueVideoList, nil
}

// GetVideoListByUserID 得到用户发表过的视频列表,返回视频数
func GetVideoListByUserID(userID uint64, videoList *[]model.Video) (int, error) {
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
//func PublishVideo(ctx *gin.Context,userID uint64, videoID uint64, videoName string, coverName string, title string) error {
//	video := model.Video{
//		ID:       videoID,
//		Title:    title,
//		PlayUrl:  videoName,
//		CoverUrl: coverName,
//		//FavoriteCount : 0,
//		//CommentCount : 0,
//		Author: model.User{
//			ID: int(userID),
//		},
//		CreatedAt: time.Now(),
//	}
//	if global.DB.WithContext(ctx).Create(&video).Error != nil {
//		return errors.New("video表插入失败")
//	}
//	return nil
//}

func FindVideos() []model.Video {
	var videos []model.Video
	global.DB.Limit(30).Find(&videos)
	return videos
}

func FindVideoAuthor(authorid uint64) (userinfo model.UserInfo) {
	global.DB.Table("users").Where("id = ?", authorid).Find(&userinfo)
	return userinfo
}
