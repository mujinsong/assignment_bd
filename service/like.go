package service

import (
	"assignment_bd/dao"
	"assignment_bd/global"
	"github.com/gin-gonic/gin"
)

// GetLikeCountListByVideoIDList 获得评论数通过视频ID(群)
func GetLikeCountListByVideoIDList(ctx *gin.Context, videoIDList []uint, likeCountList *[]uint) error {
	var uniqueVideoList []dao.VideoLikeCount
	result := global.DB.Model(&dao.VideoLike{}).WithContext(ctx).Select("video_id", "COUNT(video_id) as like_count").
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
