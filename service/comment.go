package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
)

// GetCommentCountListByVideoIDList 获得评论数通过视频ID(群)
func GetCommentCountListByVideoIDList(videoIDList []uint64, commentCountList *[]uint64) error {
	var uniqueVideoList []model.VideoCommentCount
	result := global.DB.Model(model.Comment{}).Select("video_id, COUNT(id) as count").
		Where("action_type = 1 AND video_id in ?", videoIDList).Group("video_id").Find(&uniqueVideoList)
	if result.Error != nil {
		return result.Error
	}
	numVideos := result.RowsAffected
	*commentCountList = make([]uint64, 0, numVideos)
	mapVideoIDToCommentCount := make(map[uint64]uint64, numVideos)
	for _, each := range uniqueVideoList {
		mapVideoIDToCommentCount[uint64(each.VideoId)] = uint64(each.Count)
	}
	for _, videoID := range videoIDList {
		*commentCountList = append(*commentCountList, mapVideoIDToCommentCount[uint64(videoID)])
	}
	return nil
}
