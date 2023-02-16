package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
)

func PublishVideo(video *model.Video) error {
	err := global.DB.Create(video).Error
	if err != nil {
		return err
	}
	return nil
}

/*
通过用户ID获取其发布的视频列表
*/
func GetPublishList(userID uint64) ([]model.Video, error) {
	var videos []model.Video
	err := global.DB.Table("videos").Where("user_id = ?", userID).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
