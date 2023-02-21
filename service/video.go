package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
	"errors"
	"strconv"
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
		userIDList[i] = video.UserID
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

/*
返回按投稿时间倒序的视频列表 参数为最后一次获取的视频的时间
*/
func FindVideos(lasttime string) []model.Video {
	var videos []model.Video
	last_time, err := strconv.Atoi(lasttime)
	if err != nil {
		return videos
	}
	global.DB.Table("videos").Order("created_at DESC").Limit(30).Where("created_at <= ?", time.Unix(int64(last_time), 0)).Find(&videos)
	return videos
}

//func FindUser(authorid uint64) (userinfo model.UserInfo) {
//	global.DB.Table("users").Where("id = ?", authorid).Find(&userinfo)
//	return userinfo
//}

/*
根据actionType的值来判断是增加还是减少视频的评论数
*/
func UpdateVideoCommentCount(videoID uint64, actionType uint8) error {
	var CommentCount uint64
	result := global.DB.Table("videos").Select("comment_count").Where("id = ?", videoID).Find(&CommentCount)
	if result.Error != nil {
		return result.Error
	}
	if actionType == 1 {
		CommentCount++
	} else {
		CommentCount--
	}
	result = global.DB.Table("videos").Where("id = ?", videoID).Update("comment_count", CommentCount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
