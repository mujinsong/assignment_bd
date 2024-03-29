package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
	"errors"
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

	mapVideoIDToLikeCount := make(map[uint64]uint64, numVideos)
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
func GetLikeCountListByVideoIDList(videoIDList []uint64, likeCountList *[]uint64) error {
	var uniqueVideoList []model.VideoLikeCount
	result := global.DB.Model(&model.VideoLike{}).Select("video_id", "COUNT(video_id) as like_count").
		Where("video_id in ?", videoIDList).Group("video_id").Find(&uniqueVideoList)
	if result.Error != nil {
		return result.Error
	}
	numVideos := result.RowsAffected
	*likeCountList = make([]uint64, 0, numVideos)
	mapVideoIDToLikeCount := make(map[uint64]uint64, numVideos)
	for _, each := range uniqueVideoList {
		mapVideoIDToLikeCount[each.VideoId] = each.LikeCount
	}
	for _, videoID := range videoIDList {
		*likeCountList = append(*likeCountList, mapVideoIDToLikeCount[videoID])
	}
	return nil
}

// Like 点赞视频操作
func Like(uid uint64, videoID uint64, actionType uint8) error {
	// 查询数据库是否已经存在数据
	// 有： 更新 是否相同：是：不操作
	//                   不是：更新
	// 没有：是不是点赞，是：插入数据
	//                不是：不操作
	var likeLog []model.VideoLike
	global.DB.Table("likes").Where("user_id = ? AND video_id = ?", uid, videoID).Find(&likeLog)
	if len(likeLog) == 0 {
		// 不存在记录
		if actionType == 1 {
			// 第一次点赞

			global.DB.Create(&model.VideoLike{
				UserID:     uid,
				VideoId:    videoID,
				ActionType: actionType,
			})
		}
	} else {
		// 存在记录
		// 此时actionType==2 说明用户是点赞然后取消点赞
		if actionType == 2 {
			// 取消点赞 将记录中的actionType改为2即可
			likeLog[0].ActionType = 2
			global.DB.Save(&likeLog)
		} else {
			// 这个时候是用户取消点赞然后又点赞
			likeLog[0].ActionType = 1
			global.DB.Save(&likeLog[0])
		}
	}
	// 更新视频的点赞数
	err := UpdateVideoLikes(videoID, actionType)
	if err != nil {
		return err
	}
	return nil
}

/*
更新videos表中的点赞数
*/
func UpdateVideoLikes(videoID uint64, actionType uint8) error {
	var favoriteCount uint64
	res := global.DB.Table("videos").Select("favorite_count").Where("id = ?", videoID).Take(&favoriteCount)
	if res.RowsAffected == 0 {
		return errors.New("获取视频点赞数失败")
	}
	if actionType == 1 {
		favoriteCount++
	} else {
		favoriteCount--
	}
	res = global.DB.Table("videos").Where("id = ?", videoID).Update("favorite_count", favoriteCount)
	if res.RowsAffected == 0 {
		return errors.New("更新视频点赞数失败")
	}
	return nil
}

// GetLikeVideoIDListByUserID 获取用户喜欢的列表
func GetLikeVideoIDListByUserID(uid uint64) (*[]uint64, error) {
	var videoList []uint64
	res := global.DB.Select("video_id").Where("user_id = ? AND action_type = 1", uid).Model(model.VideoLike{}).Find(&videoList)
	if res.Error != nil {
		return nil, res.Error
	}
	return &videoList, nil
}
