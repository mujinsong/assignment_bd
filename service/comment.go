package service

import (
	"assignment_bd/dao"
	"assignment_bd/global"
)

// GetCommentCountListByVideoIDListSql 被调用当且仅当VideoID不在cache中，不得不通过sql查询
func GetCommentCountListByVideoIDListSql(videoIDList []uint64, commentCountList *[]int64) error {
	var uniqueVideoList []dao.VideoCommentCount
	result := global.DB.Model(&dao.Comment{}).Select("video_id", "COUNT(video_id) as comment_count").
		Where("video_id in ?", videoIDList).Group("video_id").Find(&uniqueVideoList)
	if result.Error != nil {
		return result.Error
	}
	numVideos := result.RowsAffected
	// 针对查询结果建立映射关系
	*commentCountList = make([]int64, 0, numVideos)
	mapVideoIDToCommentCount := make(map[uint64]int64, numVideos)
	for _, each := range uniqueVideoList {
		mapVideoIDToCommentCount[each.VideoId] = each.Count
	}
	for _, videoID := range videoIDList {
		*commentCountList = append(*commentCountList, mapVideoIDToCommentCount[videoID])
	}
	return nil
}
