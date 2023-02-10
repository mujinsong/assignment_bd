package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
)

/*
判断是否已经收藏视频
uid 是当前登录用户的id
*/
func IsFavorite(uid, videoId uint64) bool {

	var likeLog model.VideoLike
	global.DB.Table("likes").Where("user_id = ? AND video_id = ?", uid, videoId).First(&likeLog)
	if likeLog.ID == 0 {
		return false
	} else {
		return true
	}
}
