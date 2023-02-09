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
func GetLikeCountListByVideoIDList(videoIDList []uint64, likeCountList *[]int) error {
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
		*likeCountList = append(*likeCountList, mapVideoIDToLikeCount[int(videoID)])
	}
	return nil
}

// Like 点赞视频操作
func Like(uid int64, videoID int64, actionType int32) error {

	// TODO 从token中获取用户ID
	// 查询数据库是否已经存在数据
	// 有： 更新 是否相同：是：不操作
	//                   不是：更新
	// 没有：是不是点赞，是：插入数据
	//                不是：不操作
	likeID := 0
	result := global.DB.Model(model.VideoLike{}).Select("id").Where("user_id = ? AND video_id = ?", uid, videoID).Take(likeID)
	if result.RowsAffected == 0 {
		likeModel := model.VideoLike{
			VideoId:    int(videoID),
			UserId:     int(uid),
			ActionType: int8(actionType % 2),
		}
		res := global.DB.Create(&likeModel)
		if res.RowsAffected == 0 {
			return errors.New("创建第一次模型失败，点赞失败")
		}
		return nil
	}
	likeModel := model.VideoLike{}
	res := global.DB.Take(&likeModel)
	if res.RowsAffected == 0 {
		return errors.New("获取点赞状态失败")
	}
	likeModel.ActionType = int8(actionType % 2)
	res = global.DB.Save(&likeModel)
	if res.RowsAffected == 0 {
		return errors.New("更改状态失败")
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
