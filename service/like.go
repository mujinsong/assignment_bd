package service

import (
	"assignment_bd/dao"
	"assignment_bd/global"
)

func GetUserLikeListByVideoIDList(userId uint, videoIDList []uint, likeList *[]bool) error {
	var uniqueVideoList []dao.VideoLikeCount
	result := global.DB.Model(&dao.VideoLike{}).Select("video_id", "action_type").
		Where("user_id = ? AND video_id in ? ", userId, videoIDList).Group("video_id").Find(&uniqueVideoList)
	if result.Error != nil {
		return result.Error
	}
	numVideos := result.RowsAffected
	*likeList = make([]bool, 0, numVideos)
	mapVideoIDToLikeCount := make(map[uint]uint, numVideos)
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
func GetLikeCountListByVideoIDList(videoIDList []uint, likeCountList *[]uint) error {
	var uniqueVideoList []dao.VideoLikeCount
	result := global.DB.Model(&dao.VideoLike{}).Select("video_id", "COUNT(video_id) as like_count").
		Where("video_id in ?", videoIDList).Group("video_id").Find(&uniqueVideoList)
	if result.Error != nil {
		return result.Error
	}
	numVideos := result.RowsAffected
	*likeCountList = make([]uint, 0, numVideos)
	mapVideoIDToLikeCount := make(map[uint]uint, numVideos)
	for _, each := range uniqueVideoList {
		mapVideoIDToLikeCount[each.VideoId] = each.LikeCount
	}
	for _, videoID := range videoIDList {
		*likeCountList = append(*likeCountList, mapVideoIDToLikeCount[videoID])
	}
	return nil
}

// 通过用户ID获取点赞列表
func GetLikeListByUserID(userID int64) ([]dao.VideoLike, error) {
	// TODO 获取用户详细信息和视频详细信息
	var likeList []dao.VideoLike

	if result := global.DB.Model(&dao.VideoLike{}).Where("user_id = ?", userID).Find(&likeList); result.Error != nil {
		return nil, result.Error
	}

	return likeList, nil

}
