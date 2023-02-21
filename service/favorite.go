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
	if uid == 0 {
		return false
	}
	var likeLog []model.VideoLike
	global.DB.Table("likes").Where("user_id = ? AND video_id = ?", uid, videoId).Find(&likeLog)
	// 通过判断likeLog的长度来判断是否已经收藏
	if len(likeLog) == 0 {
		return false
	} else {
		// 校验actionType是否为1
		if likeLog[0].ActionType == 1 {
			return true
		}
		return false
	}
}
