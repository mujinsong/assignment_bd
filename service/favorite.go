package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
)

/*
	判断是否已经收藏视频
*/
func IsFavorite(userId, videoId uint64) bool {
	// 通过userId和videoId查询是否存在记录
	var likeLog model.VideoLike
	global.DB.Table("likes").Where("user_id = ? AND video_id = ?", userId, videoId).First(&likeLog)
	if likeLog.ID == 0 {
		return false
	} else {
		return true
	}
}
