package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
)

// GetCommentCountListByVideoIDList 获得评论数通过视频ID(群)

//func GetCommentCountListByVideoIDList(videoIDList []uint64, commentCountList *[]uint64) error {
//	var uniqueVideoList []model.VideoCommentCount
//	result := global.DB.Model(model.Comment{}).Select("video_id, COUNT(id) as count").
//		Where("action_type = 1 AND video_id in ?", videoIDList).Group("video_id").Find(&uniqueVideoList)
//	if result.Error != nil {
//		return result.Error
//	}
//	numVideos := result.RowsAffected
//	*commentCountList = make([]uint64, 0, numVideos)
//	mapVideoIDToCommentCount := make(map[uint64]uint64, numVideos)
//	for _, each := range uniqueVideoList {
//		mapVideoIDToCommentCount[uint64(each.VideoId)] = uint64(each.Count)
//	}
//	for _, videoID := range videoIDList {
//		*commentCountList = append(*commentCountList, mapVideoIDToCommentCount[uint64(videoID)])
//	}
//	return nil
//}

/*
	将评论存入数据库

用户可以对同一个视频进行多次评论
*/
//func SaveComment(comment *model.Comments) error {
//	// 首先通过video_id和user_id查询是否已经评论过
//
//}

func CreateComment(comment *model.Comments) error {
	comment.ActionType = 1
	result := global.DB.Table("comments").Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
获取某个视频的评论列表
*/
func GetCommentList(videoID uint64, uid uint64) (commentList []model.Comment, err error) {
	var comments []model.Comments
	result := global.DB.Table("comments").Where("video_id = ? AND action_type = 1", videoID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	// 获取用户信息
	for comment := range comments {
		commentList = append(commentList, model.Comment{
			ID:         comments[comment].ID,
			User:       UserInfoGetByUserID(comments[comment].UserID, uid),
			Content:    comments[comment].Content,
			CreateDate: comments[comment].CreateDate,
		})
	}
	return commentList, nil
}

/*
删除评论
直接修改数据库中的action_type 将其置为2 表示删除
*/
func DeleteComment(commentID uint64) error {
	result := global.DB.Table("comments").Where("id = ?", commentID).Update("action_type", uint(2))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
