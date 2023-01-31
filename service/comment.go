package service

// GetCommentCountListByVideoIDList 获得评论数通过视频ID(群)
//func GetCommentCountListByVideoIDList(ctx *gin.Context, videoIDList []uint, commentCountList *[]uint) error {
//	var uniqueVideoList []model.VideoCommentCount
//	result := global.DB.Model(&model.Comment{}).WithContext(ctx).Select("video_id", "COUNT(video_id) as comment_count").
//		Where("video_id in ?", videoIDList).Group("video_id").Find(&uniqueVideoList)
//	if result.Error != nil {
//		return result.Error
//	}
//	numVideos := result.RowsAffected
//	*commentCountList = make([]uint, 0, numVideos)
//	mapVideoIDToCommentCount := make(map[uint]uint, numVideos)
//	for _, each := range uniqueVideoList {
//		mapVideoIDToCommentCount[each.VideoId] = each.Count
//	}
//	for _, videoID := range videoIDList {
//		*commentCountList = append(*commentCountList, mapVideoIDToCommentCount[videoID])
//	}
//	return nil
//}
