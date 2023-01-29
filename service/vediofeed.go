package service

import (
	"assignment_bd/dao"
	"assignment_bd/global"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

// GetFeedVideosAndAuthors 获取推送视频以及其作者并返回视频数
func GetFeedVideosAndAuthors(videoList *[]dao.Video, users *[]dao.User, LatestTime time.Time, MaxNumVideo int) (int, error) {

	result := global.DB.Model(&dao.Video{}).Select("id,user_id").Where("created_at > ?", LatestTime).Order("created_at DESC").Limit(MaxNumVideo).Find(videoList)
	if result.RowsAffected == 0 {
		return 0, errors.New("没视频")
	}

	numVideos := len(*videoList)

	// 批量或者视频作者
	userIDList := make([]uint, numVideos)
	for i, video := range *videoList {
		userIDList[i] = video.UserId
	}
	result = global.DB.Model(&dao.User{}).Where("user_id IN ?", userIDList).Find(users)
	if result.RowsAffected == 0 {
		return 0, errors.New("没作者")
	}
	return numVideos, nil
}
func GetVideoListByIDs(ctx *gin.Context, videoList *[]dao.Video, videoIDs []uint) error {
	var uniqueVideoList []dao.Video
	result := global.DB.Where("video_id in ?", videoIDs).Find(&uniqueVideoList)
	if result.Error != nil {
		return result.Error
	}
	numVideos := result.RowsAffected
	*videoList = make([]dao.Video, 0, numVideos)
	mapVideoIDToVideo := make(map[uint]dao.Video, numVideos)
	for _, video := range uniqueVideoList {
		mapVideoIDToVideo[video.Id] = video
	}
	// 查询like_count与comment_count
	var commentCountList []uint
	var likeCountList []uint
	if err := GetCommentCountListByVideoIDList(ctx, videoIDs, &commentCountList); err != nil {
		return err
	}
	//todo
	err := GetLikeCountListByVideoIDList(videoIDs, &likeCountList)
	if err != nil {
		return err
	}
	for i, videoID := range videoIDs {
		tmpVideo := mapVideoIDToVideo[videoID]
		tmpVideo.FavoriteCount = likeCountList[i]
		tmpVideo.CommentCount = commentCountList[i]
		*videoList = append(*videoList, tmpVideo)
	}
	return nil
}

// GetVideoIDListByUserID 得到用户发表过的视频id列表
func GetVideoIDListByUserID(ctx *gin.Context, userID uint64, videoIDList *[]uint) error {
	var videoList []dao.Video
	result := global.DB.WithContext(ctx).Where("author_id = ?", userID).Find(&videoList)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	numVideos := int(result.RowsAffected)
	*videoIDList = make([]uint, numVideos)
	for i, videoID := range videoList {
		// 最新的视频放在前面
		(*videoIDList)[numVideos-i-1] = videoID.Id
	}

	return nil

}

// PublishVideo 将用户上传的视频信息写入数据库(old)
//func PublishVideo(ctx *gin.Context,userID uint64, videoID uint64, videoName string, coverName string, title string) error {
//	video := dao.Video{
//		Id:       videoID,
//		Title:    title,
//		PlayUrl:  videoName,
//		CoverUrl: coverName,
//		//FavoriteCount : 0,
//		//CommentCount : 0,
//		Author: dao.User{
//			Id: uint(userID),
//		},
//		CreatedAt: time.Now(),
//	}
//	if global.DB.WithContext(ctx).Create(&video).Error != nil {
//		return errors.New("video表插入失败")
//	}
//	return nil
//}

// PublishVideo 将用户上传的视频信息写入数据库(new)
func PublishVideo(ctx *gin.Context, video dao.Video) error {
	if global.DB.WithContext(ctx).Create(&video).Error != nil {
		return errors.New("video表插入失败")
	}
	return nil
}
