package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
)

func GetUserLikeListByVideoIDList(userId int, videoIDList []int, likeList *[]bool) error {
	var uniqueVideoList []model.VideoLikeCount
	result := global.DB.Model(&model.VideoLike{}).Select("video_id", "action_type").
		Where("user_id = ? AND video_id in ? ", userId, videoIDList).Group("video_id").Find(&uniqueVideoList)

	if result.Error != nil {
		return result.Error
	}
	numVideos := result.RowsAffected
	*likeList = make([]bool, 0, numVideos)

	mapVideoIDToLikeCount := make(map[int]int, numVideos)
	for _, each := range uniqueVideoList {
		mapVideoIDToLikeCount[each.VideoId] = each.LikeCount
	}

	for _, videoID := range videoIDList {
		*likeList = append(*likeList, func() bool {
			if uniqueVideoList[videoID].LikeCount == 0 {
				return false
			} else {
				return true
			}
		}())
	}
	return nil
}

// GetLikeCountListByVideoIDList 获得Like数通过视频ID(群)
func GetLikeCountListByVideoIDList(videoIDList []int, likeCountList *[]int) error {
	var uniqueVideoList []model.VideoLikeCount
	result := global.DB.Model(&model.VideoLike{}).Select("video_id", "COUNT(video_id) as like_count").
		Where("video_id in ?", videoIDList).Group("video_id").Find(&uniqueVideoList)
	if result.Error != nil {
		return result.Error
	}
	numVideos := result.RowsAffected
	*likeCountList = make([]int, 0, numVideos)
	mapVideoIDToLikeCount := make(map[int]int, numVideos)
	for _, each := range uniqueVideoList {
		mapVideoIDToLikeCount[each.VideoId] = each.LikeCount
	}
	for _, videoID := range videoIDList {
		*likeCountList = append(*likeCountList, mapVideoIDToLikeCount[videoID])
	}
	return nil
}

func Like(uid int64, videoId int64, actionType int32) error {

	// TODO 从token中获取用户ID
	// 查询数据库是否已经存在数据
	// 有： 更新 是否相同：是：不操作
	//                   不是：更新
	// 没有：是不是点赞，是：插入数据
	//                不是：不操作

	return nil
}
